package main

import (
	"fmt"

	"sort/bubbleSort/bubble"
)

func main() {
	num := []int{4, 2, 6, 1, 7, 8}
	num2 := make([]int, len(num))
	copy(num2, num)
	fmt.Println(bubble.BubbleSort(num))
	fmt.Println(bubble.ImprovedBubbleSort(num2))
}
