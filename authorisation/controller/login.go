package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginInfo
	err := decodeRequest(r.Body, &loginRequest)
	if err != nil {
		Error.Printf(ERROR_MSG_FORMAT, errorMap[http.StatusBadRequest], err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if loginRequest.getCookie() != "" {
		affect, err := loginRequest.saveCookie()
		if err != nil {
			Error.Printf(ERROR_MSG_FORMAT, errorMap[http.StatusInternalServerError], err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if affect == 0 {
			Warn.Printf(WARN_MSG_FORMAT, errors.New("There is no user to set cookie"))
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	userJs, err := createUserJson(&loginRequest)
	if err == sql.ErrNoRows {
		Warn.Printf(WARN_MSG_FORMAT, errors.New("There is no user"))
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		Error.Printf(ERROR_MSG_FORMAT, errorMap[http.StatusInternalServerError], err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(w, userJs, http.StatusOK)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginInfo
	err := decodeRequest(r.Body, &loginRequest)
	if err != nil {
		Error.Printf(ERROR_MSG_FORMAT, errorMap[http.StatusBadRequest], err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = loginRequest.save()
	if err != nil {
		Error.Printf(ERROR_MSG_FORMAT, errorMap[http.StatusInternalServerError], err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userJs, err := createUserJson(&loginRequest)
	if err != nil {
		Error.Printf(ERROR_MSG_FORMAT, errorMap[http.StatusInternalServerError], err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeResponse(w, userJs, http.StatusCreated)
}

func createUserJson(lr LoginRepo) ([]byte, error) {
	user, err := lr.getUser()
	if err != nil {
		return nil, err
	}
	userJs, err := json.Marshal(&user)
	if err != nil {
		return nil, err
	}

	return userJs, nil
}
