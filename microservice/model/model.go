package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Todo struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}

type TodoList struct {
	UserName string `json:"user name"`
	Todos    []Todo `json:"todo list"`
}

type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}
