package main

import (
	"testing"
)

func TestCheckingMatchBracket(t *testing.T) {
	matched := "NO"
	brackets := "{{{{(([["

	result := checkMatchBracket(brackets)
	if matched != result {
		t.Error("Expected", true,
			"Got", false)
	}
}
