package main

import (
	"fmt"
	"net/url"
)

func main() {
	var inputNum int
	fmt.Scanf("%d", &inputNum)
	inputStrings := make([]string, inputNum, inputNum)
	for index := range inputStrings {
		fmt.Scanf("%s", &inputStrings[index])
		inputStrings[index], _ = url.QueryUnescape(inputStrings[index])
	}
	for _, inputString := range inputStrings {
		fmt.Printf("%s\n", inputString)
	}
}
