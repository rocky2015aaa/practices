package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_OF_CLOCKS = 16
	TARGET_TIME   = 12
	CYCLE_TIME    = 12
	QUARTER_TIME  = 3
	INF           = 99999
)

var (
	switches = [][]int{
		{0, 1, 2},
		{3, 7, 9, 11},
		{4, 10, 14, 15},
		{0, 4, 5, 6, 7},
		{6, 7, 8, 10, 12},
		{0, 2, 14, 15},
		{3, 14, 15},
		{4, 5, 7, 14, 15},
		{1, 2, 3, 4, 5},
		{3, 4, 5, 9, 13},
	}
	clocks [][]int
)

func main() {
	var input int
	fmt.Scanf("%d", &input)
	rand.Seed(time.Now().UTC().UnixNano())
	clocks = getClocks(input)
	for i := range clocks {
		fmt.Println(clocks[i])
	}
	for i := range clocks {
		ans := checkClocks(i, 0, 0)
		if ans == INF {
			fmt.Println(-1)
		} else {
			fmt.Println(ans)
		}
	}
}

func getClocks(numOfTry int) [][]int {
	totalClocks := make([][]int, numOfTry, numOfTry)
	setClock := func() []int {
		clocks := make([]int, NUM_OF_CLOCKS, NUM_OF_CLOCKS)
		for i := range clocks {
			clocks[i] = (rand.Intn(4) + 1) * 3
		}
		return clocks
	}

	for j := range totalClocks {
		totalClocks[j] = setClock()
	}

	return totalClocks
}

func isdone(index int) bool {
	for i := 0; i < NUM_OF_CLOCKS; i++ {
		if clocks[index][i] != TARGET_TIME {
			return false
		}
	}
	return true
}

func checkClocks(index, deep, cnt int) int {
	if isdone(index) {
		return cnt
	}
	if deep == len(switches) {
		return INF
	}
	ret := INF

	for i := 0; i < 4; i++ {
		ret = min(ret, checkClocks(index, deep+1, cnt+i))
		for j := 0; j < len(switches[deep]); j++ {
			clocks[index][switches[deep][j]] += QUARTER_TIME
			if clocks[index][switches[deep][j]] > CYCLE_TIME {
				clocks[index][switches[deep][j]] = QUARTER_TIME
			}
		}
	}
	return ret
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
