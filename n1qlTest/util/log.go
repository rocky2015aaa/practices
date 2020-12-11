package util

import (
	"log"
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
