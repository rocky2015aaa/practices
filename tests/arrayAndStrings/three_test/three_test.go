package three_test

import (
	"test/arrayAndStrings/three"
	"testing"
)

func TestCheckPermutations(t *testing.T) {
	var testCases = []struct {
		inputs []string
		expect bool
	}{
		{[]string{"abcd", "cbda"}, true},
		{[]string{"23rwef", "rwef34"}, false},
	}

	for _, testCase := range testCases {
		result := three.CheckPermutation(testCase.inputs[0], testCase.inputs[1])
		if testCase.expect != result {
			t.Error("inputs", testCase.inputs, "expected", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := three.CheckPermutationAscii(testCase.inputs[0], testCase.inputs[1])
		if testCase.expect != result {
			t.Error("inputs", testCase.inputs, "expected", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := three.CheckPermutationAscii2(testCase.inputs[0], testCase.inputs[1])
		if testCase.expect != result {
			t.Error("inputs", testCase.inputs, "expected", testCase.expect, "got", result)
		}
	}

}

func BenchmarkCheckPermutation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		three.CheckPermutation("abcd", "cbda")
	}
}

func BenchmarkCheckPermutationAscii(b *testing.B) {
	for n := 0; n < b.N; n++ {
		three.CheckPermutationAscii("abcd", "cbda")
	}
}

func BenchmarkCheckPermutationAscii2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		three.CheckPermutationAscii2("abcd", "cbda")
	}
}
