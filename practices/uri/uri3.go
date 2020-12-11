package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

/*
 - uri encoding converter
   The program to convert encoded characters to special characters with uri
 #input : number of uris, uris each
 #output : coverted uris with special characters
*/

func main() {
	input := func() string {
		var inputString string
		fmt.Scanf("%s", &inputString)
		return inputString
	}

	numOfInputs, _ := strconv.Atoi(input())
	inputStrings := make([]string, numOfInputs, numOfInputs)
	for i := 0; i < numOfInputs; i++ {
		inputString := input()
		inputStrings = append(inputStrings, inputString)
	}

	newUris := ConvertUri(inputStrings[1:])
	for _, newUri := range newUris {
		fmt.Printf("%s\n", newUri)
	}
}

func ConvertUri(inputStrings []string) []string {
	for index, _ := range inputStrings {
		u, _ := url.Parse(inputStrings[index])
		inputStrings[index] = strings.Replace(u.String(), u.RequestURI(), u.Path, -1)
	}
	return inputStrings
}
