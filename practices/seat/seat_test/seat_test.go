package seat_test

import (
	"practices/seat/seat"
	"testing"
)

func TestPossibleSeats(t *testing.T) {
	var testCases = []struct {
		n      int
		s      string
		expect int
	}{
		{2, "1A 2F 1C", 4},
		{1, "", 3},
		{10, "10G", 30},
	}

	for _, tCase := range testCases {
		result := seat.PossibleSeats(tCase.n, tCase.s)
		if result != tCase.expect {
			t.Error("expected", tCase.expect, "got", result)
		}
	}
}
