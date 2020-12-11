package util

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	LoggerInfo    *log.Logger
	LoggerWarning *log.Logger
	LoggerError   *log.Logger
)

func init() {
	logProperty := log.Ldate | log.Ltime | log.Lshortfile

	LoggerInfo = log.New(os.Stdout, "INFO: ", logProperty)
	LoggerWarning = log.New(os.Stdout, "WARNING: ", logProperty)
	LoggerError = log.New(os.Stderr, "ERROR: ", logProperty)
}

func PrintLog(logger *log.Logger, res http.ResponseWriter, err error, msg string) {
	if err != nil {
		logger.Println(msg, "\nerror -", err)
		fmt.Fprintln(res, msg, "\nerror -", err)
	} else {
		logger.Println(msg)
		fmt.Fprintln(res, msg)
	}
}
