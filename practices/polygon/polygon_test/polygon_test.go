package polygon_test

import (
	"practices/polygon/polygon"
	"reflect"
	"testing"
)

func TestCheckPolygon(t *testing.T) {
	var testCases = []struct {
		input  [][]int
		expect []int
	}{
		{
			[][]int{
				{36, 30, 36, 30},
				{15, 15, 15, 15},
				{46, 96, 90, 100},
				{86, 86, 86, 86},
				{100, 200, 100, 200},
				{-100, 200, -100, 200},
			},
			[]int{2, 2, 2},
		},
	}

	for _, testCase := range testCases {
		result := polygon.CheckPolygon(testCase.input)
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
