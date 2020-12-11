package one_test

import (
	"test/arrayAndStrings/one"
	"testing"
)

func TestCheckAllUniques(t *testing.T) {
	var testCases = []struct {
		input  string
		expect bool
	}{
		{"asdgwert", true},
		{"oktkwjwej", false},
	}

	for _, testCase := range testCases {
		result := one.CheckAllUnique(testCase.input)
		if result != testCase.expect {
			t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := one.CheckAllUnique2(testCase.input)
		if result != testCase.expect {
			t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := one.CheckAllUniqueAscii(testCase.input)
		if result != testCase.expect {
			t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
		}
	}

	for _, testCase := range testCases {
		result := one.CheckAllUniqueLowerCase(testCase.input)
		if result != testCase.expect {
			t.Error("input", testCase.input, "expected", testCase.expect, "got", result)
		}
	}
}

func BenchmarkAllUnique(b *testing.B) {
	for n := 0; n < b.N; n++ {
		one.CheckAllUnique("asdgwert")
	}
}

func BenchmarkAllUnique2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		one.CheckAllUnique2("asdgwert")
	}
}

func BenchmarkAllUniqueAscii(b *testing.B) {
	for n := 0; n < b.N; n++ {
		one.CheckAllUniqueAscii("asdgwert")
	}
}

func BenchmarkAllUniqueLowerCase(b *testing.B) {
	for n := 0; n < b.N; n++ {
		one.CheckAllUniqueLowerCase("asdgwert")
	}
}
