package sibling_test

import (
	"practices/maxSibling/sibling"
	"testing"
)

func TestGetMaxSibling(t *testing.T) {
	var siblingTest = []struct {
		n        int
		expected int
	}{
		{213, 321},
		{553, 553},
	}

	for _, testCase := range siblingTest {
		max := sibling.GetMaxSibling(testCase.n)
		if max != testCase.expected {
			t.Error("expected", testCase.expected, "got", max)
		}
	}
}
