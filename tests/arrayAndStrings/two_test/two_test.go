package two_test

import (
	"test/arrayAndStrings/two"
	"testing"
)

func TestReverses(t *testing.T) {
	var testCases = []struct {
		input  string
		expect string
	}{
		{"destruction", "noitcurtsed"},
		{"christmas", "samtsirhc"},
	}

	for _, testCase := range testCases {
		result := two.Reverse(testCase.input)
		if testCase.expect != result {
			t.Error("input", testCase.input, "expect", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := two.Reverse2(testCase.input)
		if testCase.expect != result {
			t.Error("input", testCase.input, "expect", testCase.expect, "got", result)
		}
	}

}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		two.Reverse("destruction")
	}
}

func BenchmarkReverse2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		two.Reverse2("destruction")
	}
}
