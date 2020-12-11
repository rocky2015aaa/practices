package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 5, 3, 3, 7}
	B := []int{1, 3, 5, 3, 4}
	C := []int{1, 3, 5}
	fmt.Println(CheckOneSwapPossible(A))
	fmt.Println(CheckOneSwapPossible(B))
	fmt.Println(CheckOneSwapPossible(C))
}

func CheckOneSwapPossible(A []int) bool {
	temp := make([]int, len(A), len(A))
	copy(temp, A)
	sort.Ints(A)

	var count int
	for i := range A {
		if temp[i] != A[i] {
			count++
		}
	}

	if count > 2 {
		return false
	}
	return true
}
