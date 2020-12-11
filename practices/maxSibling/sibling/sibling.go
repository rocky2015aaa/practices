package sibling

import (
	"sort"
)

func GetMaxSibling(N int) int {
	num := []int{}
	for {
		if N != 0 {
			num = append(num, N%10)
		} else {
			break
		}
		N /= 10
	}

	sort.Ints(num)
	var max int
	for i := range num {
		temp := 1
		for j := 0; j < i; j++ {
			temp *= 10
		}
		max += (num[i] * temp)
	}

	return max
}
