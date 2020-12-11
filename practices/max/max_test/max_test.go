package max_test

import (
	"practices/max/max"
	"testing"
)

func TestSolution(t *testing.T) {
	var testCases = []struct {
		input  []int
		expect int
	}{
		{[]int{4, 6, 2, 2, 6, 6, 4}, 5},
	}

	for _, testCase := range testCases {
		result := max.Solution2(testCase.input)
		if result != testCase.expect {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
