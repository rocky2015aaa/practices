package main

import (
	"reflect"
	"testing"

	"goMicroservice/db/proto"
)

func TestLoadData(t *testing.T) {
	todos := []db.Todo{
		db.Todo{
			Task:   "finishing homework",
			Status: "completed",
		},
		db.Todo{
			Task:   "informing the news",
			Status: "not started",
		},
		db.Todo{
			Task:   "checking daily issue",
			Status: "completed",
		},
		db.Todo{
			Task:   "washing dishes",
			Status: "completed",
		},
	}

	todoFromFile, err := loadData("./data.json")
	if err != nil {
		todoFromFile = []db.Todo{}
	}

	isEqual := reflect.DeepEqual(todos, todoFromFile)
	if !isEqual {
		t.Error("Expected", true, "Got", isEqual)
	}
}
