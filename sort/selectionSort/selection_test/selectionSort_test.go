package selection_test

import (
	"reflect"
	"testing"

	"sort/selectionSort/selection"
)

func TestSelectionSort(t *testing.T) {
	var testCases = []struct {
		n      []int
		expect []int
	}{
		{
			[]int{3, 6, 4, 2, 9, 1},
			[]int{9, 6, 4, 3, 2, 1},
		},
	}

	for _, testCase := range testCases {
		result := selection.SelectionSort(testCase.n)
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Error("expected", testCase.expect, "got", result)
		}
	}
}
