package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/server"

	"goMicroservice/auth/proto"
	"goMicroservice/db/proto"
	. "goMicroservice/util"
)

type App struct{}

/*
	- GetTodoList
	* Parameters : context, auth service request, db service response
	* Return Value : error
	The function to return data
	1. Check authentication from auth service
	2. Get data from db service
*/
func (a *App) GetTodoList(ctx context.Context, req *auth.AuthRequest, res *db.DbResponse) error {
	log.Println(APP_SERVER + " Get request")
	// Check authentication from auth service
	authRes := &auth.AuthResponse{}
	err := client.Call(ctx, client.NewRequest("Auth", "Auth.Authenticate", req), authRes)
	if err != nil {
		log.Println(APP_SERVER + " Failed user authentication")
		return err
	}
	// Get data from db service
	todoReq := client.NewRequest("Db", "Db.LoadTodoList", authRes)
	err = client.Call(ctx, todoReq, res)
	if err != nil {
		log.Println(APP_SERVER + " Failed to fetch data")
		return err
	}

	log.Println(APP_SERVER + " Fetching data is successul")
	return nil
}

func main() {
	server.Init(
		server.Name("App"),
		server.Version("1.0"),
	)

	server.Handle(server.NewHandler(&App{}))

	log.Println(APP_SERVER + " Server is started")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
