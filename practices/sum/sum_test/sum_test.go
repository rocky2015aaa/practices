package sum_test

import (
	"testing"

	"practices/sum/sum"
)

func TestCheckSum(t *testing.T) {
	var testCases = []struct {
		input  []int
		sum    int
		expect int
	}{
		{
			[]int{18, 11, 21, 28, 31, 38, 40, 55, 60, 62},
			66,
			1,
		},
	}

	for _, testCase := range testCases {
		result := sum.CheckSum(testCase.sum, testCase.input)
		if result != testCase.expect {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
