package main

import (
	"fmt"

	"practices/weirdNumbers/weirdnumber"
)

func main() {
	var numOfInput int
	fmt.Scanf("%d", &numOfInput)
	inputs := make([]int, numOfInput, numOfInput)
	outputs := make([]string, numOfInput, numOfInput)
	for i := 0; i < numOfInput; i++ {
		fmt.Scanf("%d", &inputs[i])
		outputs[i] = weirdnumber.CheckWeirdNumber(inputs[i])
	}
	for _, output := range outputs {
		fmt.Println(output)
	}
}
