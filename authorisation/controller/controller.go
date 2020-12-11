package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type RequestDecoder interface {
	validate() error
}

func decodeRequest(rBody io.Reader, rd RequestDecoder) error {
	var err error
	switch v := rd.(type) {
	case *LoginInfo:
		err = json.NewDecoder(rBody).Decode(v)
	}
	if err != nil {
		return err
	}
	err = rd.validate()
	if err != nil {
		return err
	}

	return nil
}

func writeResponse(w http.ResponseWriter, js []byte, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
