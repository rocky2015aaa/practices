package main

import (
	"fmt"

	"sort/insertionSort/insertion"
)

func main() {
	//A := []int{4, 5, 2, 7, 3, 5}
	A := []int{4, 5, 7, 2}
	fmt.Println(insertion.InsertionSort(A))
}
