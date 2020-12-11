package main

import "testing"

func TestCountWords(t *testing.T) {
	expectedInput := `banana, apple,
	melon
	ice creamEOF`

	expectedCount := 5
	realCount := getCount(expectedInput)

	if expectedCount != realCount {
		t.Error("Expected", expectedCount, "Got", realCount)
	}
}
