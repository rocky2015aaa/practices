package main

import "testing"

var ans int

func TestGetClock(t *testing.T) {
	input := 2
	clocks := getClocks(input)
	clockLen := len(clocks)
	if clockLen != input {
		t.Errorf("Expected len is %d, but it was %d instead.", input, clockLen)
	}
	elementLen := 16
	for _, clock := range clocks {
		clockElementLen := len(clock)
		if clockElementLen != elementLen {
			t.Errorf("Expected element len is %d, but it was %d instead.", clockElementLen, elementLen)
		}
	}
}

func TestCheckClocks(t *testing.T) {
	clocks = [][]int{
		{12, 6, 6, 6, 6, 6, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12},
		{12, 9, 3, 12, 6, 6, 9, 3, 12, 9, 12, 9, 12, 12, 6, 6},
	}
	answers := make([]int, len(clocks), len(clocks))
	for i := range clocks {
		ans = checkClocks(i, 0, 0)
		if ans == INF {
			answers[i] = -1
		} else {
			answers[i] = ans
		}
	}

	for _, answer := range answers {
		if answer != 2 && answer != 9 {
			t.Errorf("Expected time is 2 or 9, but it was %d instead.", ans)
		}
	}
}
