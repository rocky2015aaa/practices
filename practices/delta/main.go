package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"practices/delta/delta"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	nums := strings.Fields(input)
	numbers := make([]int, len(nums), len(nums))
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		numbers[i] = number
	}

	fmt.Println(delta.CheckDelta(numbers))
}
