package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "microservice/model"
	. "microservice/util"
)

func main() {
	handler := MyHandler{Route: make(RouteMap)}

	handler.Route["/"] = protectedTodo

	server := http.Server{
		Addr:    ":8002",
		Handler: &handler}

	log.Printf("%s %s", APP_SERVER, "service is started.\n")
	server.ListenAndServe()
}

/*
	- protectedTodo
	* Parameters : http response writer, http request

	the function to get user todo list
	1. Get authenticated user data from authentication server
	2. Get all db data from db server
	3. Response user todo list
*/
func protectedTodo(res http.ResponseWriter, req *http.Request) {
	var claims Claims

	body, err := MakeRequest(req, AUTH_LINK)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		PrintLog(LoggerWarning, res, err, APP_SERVER+" Failed to authenticate the user")
		return
	}
	if err = json.Unmarshal(body, &claims); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerWarning, res, err, APP_SERVER+" Failed to authenticate the user")
		return
	}

	todoList, err := loadTodo(req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerWarning, res, err, APP_SERVER+" Failed to load todo list")
		return
	}

	json.NewEncoder(res).Encode(*todoList)
}

// access db server to get all todo data
func loadTodo(req *http.Request) (*TodoList, error) {
	var todoList TodoList

	body, err := MakeRequest(req, DB_LINK)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &todoList); err != nil {
		return nil, err
	}

	return &todoList, nil
}
