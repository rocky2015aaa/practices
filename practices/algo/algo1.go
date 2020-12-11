package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

var (
	numberLetters = map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
	}
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	inputs := make([]string, num, num)
	outputs := inputs
	reader := bufio.NewReader(os.Stdin)
	for i := range inputs {
		inputs[i], _ = reader.ReadString('\n')
		outputs[i] = Result(inputs[i])
	}

	for j := range outputs {
		fmt.Println(outputs[j])
	}
}

func Result(input string) string {
	elements := strings.Fields(input)
	if check(elements) {
		return "Yes"
	}
	return "No"
}

func check(elements []string) bool {
	var result int
	var resultLetter string

	num1 := numberLetters[elements[0]]
	num2 := numberLetters[elements[2]]

	switch elements[1] {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	}

	for letter := range numberLetters {
		if result == numberLetters[letter] {
			resultLetter = letter
			rLetter := strings.Split(resultLetter, "")
			sort.Strings(rLetter)
			eLetter := strings.Split(elements[4], "")
			sort.Strings(eLetter)
			if reflect.DeepEqual(rLetter, eLetter) {
				return true
			}

			break
		}
	}

	return false
}
