package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	"goMicroservice/auth/proto"
	"goMicroservice/db/proto"
	. "goMicroservice/util"
)

func main() {
	service := web.NewService(
		web.Name("Router"),
		web.Version("1.0"),
		web.Address(":8000"),
	)

	service.Handle(ROUTE_PATH, http.HandlerFunc(handler))

	log.Println(ROUTER_SERVER + " Server is started")
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
	- handler
	* Parameters : http response writer, http request

	The handler function to get todo data
	1. Check if request is correct
	2. Get the request to load todo data
	3. Fetch the todo data
*/
func handler(res http.ResponseWriter, req *http.Request) {
	if !isRequestValid(res, req) {
		return
	}

	appReq, err := getAppRequest(res, req)
	if err != nil {
		return
	}

	data, err := getData(res, appReq)
	if err != nil {
		return
	}

	res.Write(data)
}

func isRequestValid(res http.ResponseWriter, req *http.Request) bool {
	// check valid path
	if req.URL.Path != ROUTE_PATH {
		res.WriteHeader(http.StatusNotFound)
		PrintLog(LoggerWarning, res, nil, ROUTER_SERVER+" Page Not Found")
		return false
	}

	// check request method
	if req.Method != "GET" {
		res.WriteHeader(http.StatusMethodNotAllowed)
		PrintLog(LoggerWarning, res, nil, ROUTER_SERVER+" "+http.ErrNotSupported.Error())
		return false
	}

	return true
}

// set the request to load data
func getAppRequest(res http.ResponseWriter, req *http.Request) (client.Request, error) {
	token := req.Header.Get("authentication_token")
	if token == "" {
		err := errors.New(ROUTER_SERVER + " Missed Authentication token")
		res.WriteHeader(http.StatusBadGateway)
		PrintLog(LoggerWarning, res, err, err.Error())
		return nil, errors.New(err.Error())
	}

	return client.NewRequest("App", "App.GetTodoList", &auth.AuthRequest{
		Token: token,
	}), nil
}

func getData(res http.ResponseWriter, req client.Request) ([]byte, error) {
	// load data
	resp := &db.DbResponse{}
	err := client.Call(context.Background(), req, resp)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		PrintLog(LoggerWarning, res, err, ROUTER_SERVER+" Failed to access data")
		return nil, err
	}

	// put data into json format
	data, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerError, res, err, ROUTER_SERVER+" Failed to return data")
		return nil, err
	}

	return data, nil
}
