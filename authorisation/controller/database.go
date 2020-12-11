package controller

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:1q2w3e@/myBoard")
	if err != nil {
		panic(err)
	}
}
