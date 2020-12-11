package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/micro/go-micro/server"

	"goMicroservice/auth/proto"
	"goMicroservice/db/proto"
	. "goMicroservice/util"
)

type Db struct{}

/*
	- LoadTodoList
	* Parameters : context, auth service response (request), db service response (response)
	* Return Value : error
	The function to get todo data
	1. Load data by accessing db service
	2. Set the data format
*/
func (d *Db) LoadTodoList(ctx context.Context, req *auth.AuthResponse, res *db.DbResponse) error {
	log.Println(DB_SERVER + " Get request")
	// Load data by accessing db service
	dbTodoList, err := loadData(DATA_PATH)
	if err != nil {
		log.Println(DB_SERVER + " Failed to load data")
		return err
	}
	// Set the data format
	res.Username = req.Username
	for index, _ := range dbTodoList {
		res.TodoList = append(res.TodoList, &dbTodoList[index])
	}

	log.Println(DB_SERVER + " Loading data is successful")
	return nil
}

func loadData(dataPath string) ([]db.Todo, error) {
	var todoList []db.Todo
	data, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Println(DB_SERVER + " " + err.Error())
		return nil, err
	}
	err = json.Unmarshal(data, &todoList)
	if err != nil {
		log.Println(DB_SERVER + " " + err.Error())
		return nil, err
	}

	return todoList, nil
}

func main() {
	server.Init(
		server.Name("Db"),
		server.Version("1.0"),
	)

	server.Handle(server.NewHandler(&Db{}))

	log.Println(DB_SERVER + " Server is started")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
