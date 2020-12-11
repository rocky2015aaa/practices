package main

import (
	"fmt"
)

func main() {
	A := []int{60, 80, 40}
	B := []int{2, 3, 5}
	fmt.Println(Solution(A, B, 5, 2, 200))
	fmt.Println(Solution2(A, B, 5, 2, 200))
}

func Solution(A, B []int, M, X, Y int) int {
	var i, totalWeight, preFloor, stopCount int
	for i = 0; i < len(A); i++ {
		if i >= X || totalWeight+A[i] >= Y {
			break
		}
	}

	for j := 0; j < i; j++ {
		if B[j] != preFloor {
			preFloor = B[j]
			stopCount++
		}
	}

	A = A[i:]
	B = B[i:]
	stopCount++

	if len(A) == 0 {
		return stopCount
	}
	return stopCount + Solution(A, B, M, X, Y)
}

func Solution2(A, B []int, M, X, Y int) int {
	var stopCount int
	for {
		var i, preFloor, totalWeight int

		for i = 0; i < len(A); i++ {
			if i >= X || totalWeight+A[i] >= Y {
				break
			}
		}

		for j := 0; j < i; j++ {
			if B[j] != preFloor {
				preFloor = B[j]
				stopCount++
			}
		}

		A = A[i:]
		B = B[i:]
		stopCount++

		if len(A) == 0 {
			break
		}
	}

	return stopCount
}
