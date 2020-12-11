package four_test

import (
	"test/arrayAndStrings/four"
	"testing"
)

func TestEncodingSpace(t *testing.T) {
	var testCases = []struct {
		input  string
		expect string
	}{
		{"Mr John Smith    ", "Mr%20John%20Smith"},
		{"lala    la la", "lala%20%20%20%20la%20la"},
	}
	/*
		for _, testCase := range testCases {
			result := four.EncodingSpace(testCase.input)
			if testCase.expect != result {
				t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
			}
		}
	*/
	for _, testCase := range testCases {
		result := four.EncodingSpace2(testCase.input)
		if testCase.expect != result {
			t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := four.EncodingSpace3(testCase.input)
		if testCase.expect != result {
			t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
		}
	}

}

func BenchmarkEncodingSpace(b *testing.B) {
	for n := 0; n < b.N; n++ {
		four.EncodingSpace("Mr John Smith    ")
	}
}

func BenchmarkEncodingSpace2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		four.EncodingSpace2("Mr John Smith    ")
	}
}

func BenchmarkEncodingSpace3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		four.EncodingSpace3("Mr John Smith    ")
	}
}
