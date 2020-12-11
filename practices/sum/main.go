package main

import (
	"fmt"

	"practices/sum/sum"
)

func main() {
	var total int
	fmt.Scanf("%d", &total)

	var numOfInputs int
	fmt.Scanf("%d", &numOfInputs)

	inputs := make([]int, numOfInputs, numOfInputs)
	for i := 0; i < numOfInputs; i++ {
		fmt.Scanf("%d", &inputs[i])
	}

	fmt.Println(sum.CheckSum(total, inputs))
}
