package bubble

import ()

func BubbleSort(numbers []int) []int {
	lastIndex := len(numbers) - 1
	for i := lastIndex; i > 0; i-- {
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func ImprovedBubbleSort(numbers []int) []int {
	lastIndex := len(numbers) - 1
	proceed := true
	for i := lastIndex; proceed; i-- {
		proceed = false
		for j := 0; j < i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				proceed = true
			}
		}
	}
	return numbers
}
