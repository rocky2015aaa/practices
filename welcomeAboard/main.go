package main

import (
	"net/http"

	con "github.com/myBoard/welcomeAboard/controller"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))
	http.HandleFunc("/", con.LoginHandler)
	http.HandleFunc("/save-cache", con.CacheHandler)
	// TODO: make it works only slash not including other letters with slash
	http.ListenAndServe(":3000", nil)
}
