package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	result := Solution("a0Ba")
	if result != 2 {
		t.Error("expected", 2, "got", result)
	}
}
