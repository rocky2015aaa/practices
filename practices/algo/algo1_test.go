package main

import (
	"testing"
)

func TestResult(t *testing.T) {
	input := "two + three = ivef"
	result := Result(input)
	if result != "Yes" {
		t.Error("expected", "Yes", "got", result)
	}
}
