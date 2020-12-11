package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	. "test/model"
)

const (
	XML_PATH    = "/home/donggeon/challenge.xml"
	SERVER_PORT = ":8080"
)

var (
	LoggerInfo    *log.Logger
	LoggerWarning *log.Logger
	LoggerError   *log.Logger

	mux     map[string]func(http.ResponseWriter, *http.Request)
	housing Housing
)

type serverHandler struct{}

func init() {
	// set logs
	logProperty := log.Ldate | log.Ltime | log.Lshortfile
	LoggerInfo = log.New(os.Stdout, "INFO: ", logProperty)
	LoggerWarning = log.New(os.Stdout, "WARNING: ", logProperty)
	LoggerError = log.New(os.Stderr, "ERROR: ", logProperty)

	// import xml file
	err := importXml(XML_PATH)
	if err != nil {
		LoggerError.Printf("Failed to import xml file\nerror: %v", err)
		os.Exit(1)
	}
}

func main() {
	server := GetServer()
	LoggerInfo.Printf("Server is started. port number%s", SERVER_PORT)
	server.ListenAndServe()
}

// set router for server api and return it
func GetServer() http.Server {
	server := http.Server{
		Addr:    SERVER_PORT,
		Handler: &serverHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = MainApi
	return server
}

/* call helper function to read xml and convert it to json.
   set ok status code and print json and info log */
func MainApi(w http.ResponseWriter, r *http.Request) {
	housingJson, err := json.Marshal(&housing)
	if err != nil {
		LoggerError.Printf("Failed to marshal json\nerror: %v", err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", string(housingJson))
	LoggerInfo.Println("Successfully called main api")
}

// open xml file, read it and unmarshal it
func importXml(path string) error {
	file, err := os.Open("/home/donggeon/challenge.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	xmlData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	housing = Housing{}
	err = xml.Unmarshal(xmlData, &housing)
	if err != nil {
		return err
	}

	return nil
}

/* run the api by path and if it is not found,
   then print proper code and warning log */
func (s *serverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	w.WriteHeader(404)
	fmt.Fprintf(w, "%s", "Page not found")
	LoggerWarning.Println("Page not found")
}
