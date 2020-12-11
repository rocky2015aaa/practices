package insertionSort_test

import (
	"reflect"
	"testing"

	"sort/insertionSort/insertion"
)

func TestInsertionSort(t *testing.T) {
	var testCases = []struct {
		n      []int
		expect []int
	}{
		{
			[]int{4, 5, 2, 7, 3, 5, 10},
			[]int{10, 7, 5, 5, 4, 3, 2},
		},
	}

	for _, testCase := range testCases {
		result := insertion.InsertionSort(testCase.n)
		if !reflect.DeepEqual(result, testCase.expect) {
			t.Error("expect", testCase.expect, "got", result)
		}
	}
}
