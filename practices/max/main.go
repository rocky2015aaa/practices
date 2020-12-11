package main

import (
	"fmt"
	"practices/max/max"
)

func main() {
	//A := []int{4, 6, 2, 2, 6, 6, 4}
	A := []int{4, 2, 3, 4, 4, 4, 4}
	fmt.Println(max.Solution(A))
	fmt.Println(max.Solution2(A))
}
