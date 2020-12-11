package weirdnumber_test

import (
	"testing"

	"practices/weirdNumbers/weirdnumber"
)

func TestCheckWeirdNumber(t *testing.T) {
	var testCases = []struct {
		n      int
		expect string
	}{
		{12, "not weird"},
		{70, "weird"},
	}

	for _, testCase := range testCases {
		result := weirdnumber.CheckWeirdNumber(testCase.n)
		if testCase.expect != result {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
