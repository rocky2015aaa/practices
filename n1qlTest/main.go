package main

import (
	"n1qlServerTest/router"
)

const (
	couchbaseUrl = "couchbase://localhost"
)

func main() {
	router.RunRouter()
}
