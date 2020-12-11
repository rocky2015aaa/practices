package router

import (
	"io"
	"n1qlServerTest/controller"
	"net/http"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}

func RunRouter() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	mux["/"] = controller.AllAirlinesController

	server.ListenAndServe()
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
