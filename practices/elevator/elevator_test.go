package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	A := []int{60, 80, 40}
	B := []int{2, 3, 5}
	result := Solution(A, B, 5, 2, 200)
	if result != 5 {
		t.Error("expected", 5, "got", result)
	}
}
