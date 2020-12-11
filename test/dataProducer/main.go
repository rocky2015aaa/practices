package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", Publish)
	if err := http.ListenAndServe(":9002", nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func Publish(resp http.ResponseWriter, request *http.Request) {
	for _ = range time.Tick(time.Second) {
		price := rand.Intn(100)
		_, err := http.PostForm("http://localhost:9003/price", url.Values{"price": {strconv.Itoa(price)}})
		if err != nil {
			log.Printf("%s\n", err)
		}
	}
}
