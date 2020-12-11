package controller

import (
	"log"
	"net/http"
	"os"
)

const (
	ERROR_MSG_FORMAT = "[%s] %s"
	INFO_MSG_FORMAT  = "%s"
	WARN_MSG_FORMAT  = "%s"
)

var (
	errorMap = map[int]string{
		http.StatusUnauthorized:        "Unauthorized",
		http.StatusServiceUnavailable:  "Service Unabailable",
		http.StatusUnprocessableEntity: "Unprocessable Entity",
		http.StatusInternalServerError: "Internal Server Error",
		http.StatusBadRequest:          "Bad Request",
	}

	Error *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
)

func init() {
	logFormat := log.Ldate | log.Ltime | log.Lshortfile
	Error = log.New(os.Stderr, "Error: ", logFormat)
	Info = log.New(os.Stdout, "Info: ", logFormat)
	Warn = log.New(os.Stdout, "Warn: ", logFormat)
}
