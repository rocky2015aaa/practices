package main

import (
	"testing"
)

func TestSwap(t *testing.T) {
	A := []int{1, 5, 3, 3, 7}
	B := []int{1, 3, 5, 3, 4}
	C := []int{1, 3, 5}
	var swapTest = []struct {
		n        []int
		expected bool
	}{
		{A, true},
		{B, false},
		{C, true},
	}
	for _, testCase := range swapTest {
		result := CheckOneSwapPossible(testCase.n)
		if result != testCase.expected {
			t.Error("expected", testCase.expected, "got", result)
		}
	}
}
