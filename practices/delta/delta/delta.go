package delta

import ()

func CheckDelta(n []int) []int {
	result := []int{n[0]}
	for i := 1; i < len(n); i++ {
		temp := n[i] - n[i-1]
		if temp > 127 || temp < -127 {
			result = append(result, -128)
		}
		result = append(result, temp)
	}

	return result
}
