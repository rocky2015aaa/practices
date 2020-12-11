package main

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	go func() {
		resp, err := http.Get("http://localhost:9002")
		if err != nil {
			log.Printf("%s\n", err)
			return
		}
		log.Printf("%s\n", resp)
	}()
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func index(resp http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("view.html")
	t.Execute(resp, "")
}
