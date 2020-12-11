package main

import (
	"practices/letterTransform/lTransform"
	"testing"
)

func TestTransform(t *testing.T) {
	var testCases = []struct {
		input  string
		expect string
	}{
		{"ABBCC", "AC"},
	}

	for _, testCase := range testCases {
		result := lTransform.Transform(testCase.input)
		if result != testCase.expect {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
