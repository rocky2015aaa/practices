package main

import (
	"reflect"
	"testing"

	. "microservice/model"
)

func TestGenerateData(t *testing.T) {
	todos := []Todo{
		Todo{
			Task:   "finishing homework",
			Status: "completed",
		},
		Todo{
			Task:   "informing the news",
			Status: "not started",
		},
		Todo{
			Task:   "checking daily issue",
			Status: "completed",
		},
		Todo{
			Task:   "washing dishes",
			Status: "completed",
		},
	}
	todoFromFile, err := generateData("./data.json")
	if err != nil {
		todoFromFile = []Todo{}
	}
	isEqual := reflect.DeepEqual(todoFromFile, todos)
	if !isEqual {
		t.Error("Expected", true, "Got", isEqual)
	}
}
