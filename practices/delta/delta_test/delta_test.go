package delta_test

import (
	"reflect"
	"testing"

	"practices/delta/delta"
)

func TestCheckDelta(t *testing.T) {
	var testCases = []struct {
		n      []int
		expect []int
	}{
		{
			[]int{25626, 25757, 24367, 24267, 16, 100, 2, 7277},
			[]int{25626, -128, 131, -128, -1390, -100, -128, -24251, 84, -98, -128, 7275},
		},
	}
	for _, testCase := range testCases {
		result := delta.CheckDelta(testCase.n)
		if !reflect.DeepEqual(testCase.expect, result) {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
