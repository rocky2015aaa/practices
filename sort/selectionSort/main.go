package main

import (
	"fmt"

	"sort/selectionSort/selection"
)

func main() {
	A := []int{4, 7, 2, 7, 3, 9}
	fmt.Println(selection.SelectionSort(A))
}
