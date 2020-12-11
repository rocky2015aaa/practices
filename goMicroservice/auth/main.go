package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/server"

	"goMicroservice/auth/proto"
	. "goMicroservice/util"
)

type Auth struct{}

/*
	- Authenticate
	* Parameters : context, auth service request, auth service response
	* Return Value : error
	The function to check authentication with token
	1. Check token
	2. Set username
*/
func (a *Auth) Authenticate(ctx context.Context, req *auth.AuthRequest, res *auth.AuthResponse) error {
	log.Println(AUTH_SERVER + " Get request")

	if req.Token != TOKEN {
		log.Println(AUTH_SERVER + " Failed to Autheticate the user")
		return errors.New(AUTH_SERVER + " Token is not valid")
	}

	res.Username = USERNAME

	log.Println(AUTH_SERVER + " Authentication is successful")
	return nil
}

func main() {
	server.Init(
		server.Name("Auth"),
		server.Version("1.0"),
	)

	server.Handle(server.NewHandler(&Auth{}))

	log.Println(AUTH_SERVER + " Server is started")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
