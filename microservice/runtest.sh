#!/bin/sh
cd auth
go test &
cd ../db
go test &
cd .. &
sleep 1
