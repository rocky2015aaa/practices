#!/bin/sh
go run ./app/main.go &
go run ./auth/main.go &
go run ./router/main.go &
cd db
go run main.go &
cd ..
