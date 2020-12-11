package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	. "microservice/model"
	. "microservice/util"
)

/*
	* Router Service
	Provider routing path
	Except logging out, all module just access other service to handle proper behavior.
	And then show proper result from response.
*/

func main() {
	handler := MyHandler{Route: make(RouteMap)}

	handler.Route["/"] = app
	handler.Route["/auth"] = auth
	handler.Route["/db"] = accessDB
	handler.Route["/logout"] = logout

	server := http.Server{
		Addr:    ":8000",
		Handler: &handler,
	}

	log.Print("[router] service is started.\n")
	server.ListenAndServe()
}

func app(res http.ResponseWriter, req *http.Request) {
	var todoList TodoList

	body, err := MakeRequest(req, APP_LINK)
	if err != nil {
		// the reason is authentication failure or server error
		// so just show accepted code
		res.WriteHeader(http.StatusAccepted)
		PrintLog(LoggerWarning, res, err, getErrorMessage(body))
		return
	}

	if err = json.Unmarshal(body, &todoList); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerWarning, res, err, getErrorMessage(body))
		return
	}

	json.NewEncoder(res).Encode(todoList)

}

func auth(res http.ResponseWriter, req *http.Request) {
	var claims Claims

	body, err := MakeRequest(req, AUTH_LINK)
	if err != nil {
		// the reason is authentication failure or server error
		// so just show accepted code
		res.WriteHeader(http.StatusAccepted)
		PrintLog(LoggerWarning, res, err, getErrorMessage(body))
		return
	}
	if err = json.Unmarshal(body, &claims); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerWarning, res, err, getErrorMessage(body))
		return
	}

	json.NewEncoder(res).Encode(claims)
}

func accessDB(res http.ResponseWriter, req *http.Request) {
	var todoList TodoList

	body, err := MakeRequest(req, DB_LINK)
	if err != nil {
		// the reason is authentication failure or server error
		// so just show accepted code
		res.WriteHeader(http.StatusAccepted)
		PrintLog(LoggerWarning, res, err, getErrorMessage(body))
		return
	}
	if err = json.Unmarshal(body, &todoList); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		PrintLog(LoggerWarning, res, err, getErrorMessage(body))
		return
	}
	json.NewEncoder(res).Encode(todoList)
}

// Redirect authentication Service to handle logout
func logout(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, AUTH_LINK+"/logout", 307)
}

func getErrorMessage(msg []byte) string {
	errorMsg := string(msg)
	return errorMsg[:strings.Index(errorMsg, "\n")]
}
