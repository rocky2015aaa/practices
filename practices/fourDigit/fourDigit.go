package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FourDigit(1, 8, 3, 2))
	fmt.Println(FourDigit(2, 4, 0, 0))
	fmt.Println(FourDigit(3, 0, 7, 0))
	fmt.Println(FourDigit(9, 1, 9, 7))
	fmt.Println(FourDigit(1, 3, 5, 9))
}

func FourDigit(A, B, C, D int) string {
	nums := []int{A, B, C, D}

	indexes := getBiggestNumIndexes(nums, 24)
	if len(indexes) == 0 {
		return "IMPORSSIBLE"
	}
	time := strconv.Itoa(nums[indexes[0]]) + strconv.Itoa(nums[indexes[1]])

	if indexes[0] > indexes[1] {
		nums = append(nums[:indexes[1]], nums[indexes[1]+1:]...)
		nums = append(nums[:indexes[0]-1], nums[indexes[0]:]...)
	} else {
		nums = append(nums[:indexes[0]], nums[indexes[0]+1:]...)
		nums = append(nums[:indexes[1]-1], nums[indexes[1]:]...)
	}

	indexes = getBiggestNumIndexes(nums, 60)
	if len(indexes) == 0 {
		return "IMPOSSIBLE"
	}
	time = time + ":" + strconv.Itoa(nums[indexes[0]]) + strconv.Itoa(nums[indexes[1]])

	return time
}

func getBiggestNumIndexes(nums []int, limit int) []int {
	var max int
	indexes := make([]int, 2, 2)
	for i := range nums {
		for j := range nums {
			if i != j {
				temp := nums[i]*10 + nums[j]
				if limit > temp && max < temp {
					max = temp
					indexes[0] = i
					indexes[1] = j
				}
			}
		}
	}

	if max == 0 {
		return []int{}
	}
	return indexes
}
