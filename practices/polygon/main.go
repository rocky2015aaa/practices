package main

import (
	"fmt"
	"practices/polygon/polygon"
)

func main() {
	var numOfInput int
	fmt.Scanf("%d", &numOfInput)
	inputs := make([][]int, numOfInput, numOfInput)
	for i := 0; i < numOfInput; i++ {
		inputs[i] = make([]int, 4, 4)
		fmt.Scanf("%d %d %d %d", &inputs[i][0], &inputs[i][1], &inputs[i][2], &inputs[i][3])
	}
	outputs := polygon.CheckPolygon(inputs)
	for _, output := range outputs {
		fmt.Printf("%d ", output)
	}
	fmt.Println()
}
