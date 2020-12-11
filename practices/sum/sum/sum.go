package sum

import ()

func CheckSum(sum int, inputs []int) int {
	inputMap := make(map[int]bool)

	for i := range inputs {
		if _, ok := inputMap[inputs[i]]; ok {
			inputMap[inputs[i]] = true
		} else {
			inputMap[inputs[i]] = false
		}
	}

	for i := range inputs {
		temp := sum - inputs[i]
		if val, ok := inputMap[temp]; ok {
			if temp == inputs[i] && !val {
				continue
			}
			return 1
		}
	}

	return 0
}
