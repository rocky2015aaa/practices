package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	. "microservice/model"
	. "microservice/util"
)

func main() {
	handler := MyHandler{Route: make(RouteMap)}

	handler.Route["/"] = protectedData

	server := http.Server{
		Addr:    ":8003",
		Handler: &handler,
	}

	log.Printf("%s %s", DB_SERVER, "service is started.\n")
	server.ListenAndServe()
}

/*
	- protectedData
	* Parameters : http response writer, http request

	the function to get todo list
	1. Check if the user is authenticated
	2. Get all db data
*/
func protectedData(res http.ResponseWriter, req *http.Request) {
	var claims Claims
	// check authentication from auth service
	body, err := MakeRequest(req, AUTH_LINK)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		PrintLog(LoggerWarning, res, err, DB_SERVER+" Failed to authenticate the user")
		return
	}
	if err = json.Unmarshal(body, &claims); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerWarning, res, err, DB_SERVER+" Failed to authenticate the user")
		return
	}

	todo, err := generateData("./data.json")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerError, res, err, DB_SERVER+" Failed to load data")
		return
	}

	todoList := TodoList{
		UserName: claims.UserName,
		Todos:    todo,
	}

	json.NewEncoder(res).Encode(todoList)
}

func generateData(jsonFilePath string) ([]Todo, error) {
	var todo []Todo

	raw, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(raw, &todo); err != nil {
		return nil, err
	}

	return todo, nil
}
