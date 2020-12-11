package main

import (
	"net/http"

	con "github.com/myBoard/authorization/controller"
)

func main() {
	http.HandleFunc("/login", con.LoginHandler)
	http.HandleFunc("/sign-up", con.SignUpHandler)
	http.ListenAndServe(":3001", nil)
}
