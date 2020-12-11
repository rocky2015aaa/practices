package main

import (
	"fmt"
	"sort"
)

/*
type List []int

func (l List) Len() int           { return len(l) }
func (l List) Less(i, j int) bool { return l[i] < l[j] }
func (l List) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
*/
func main() {
	list := []int{1, 4, 7, 3, 3, 5}
	fmt.Println("Sorted: ", Solution(list))
}

func Solution(A []int) int {
	// write your code in Go 1.4
	aMap := make(map[int][]int)
	for i, a := range A {
		aMap[a] = append(aMap[a], i)
	}

	var keys []int
	for k := range aMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var max int
	for i := 0; i < len(keys)-1; i++ {
		for _, v1 := range aMap[keys[i]] {
			for _, v2 := range aMap[keys[i+1]] {
				temp := v1 - v2
				if temp < 0 {
					temp *= -1
				}
				if temp > max {
					max = temp
				}
			}
		}
	}

	return max
}
