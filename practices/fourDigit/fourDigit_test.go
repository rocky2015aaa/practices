package main

import "testing"

func TestFourDigit(t *testing.T) {
	answer := "23:18"
	result := FourDigit(1, 8, 3, 2)
	if result != answer {
		t.Error("expected", answer, "got", result)
	}
}
