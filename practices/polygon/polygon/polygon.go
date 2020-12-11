package polygon

func CheckPolygon(inputs [][]int) []int {
	counts := make([]int, 3, 3)
	for i := range inputs {
		analysisForm(counts, inputs[i])
	}
	return counts
}

func analysisForm(counts []int, input []int) {
	if input[0] <= 0 || input[1] <= 0 || input[2] <= 0 || input[3] <= 0 {
		counts[2]++
	} else if input[0] == input[2] && input[1] == input[3] {
		if input[0] == input[1] {
			counts[0]++
		} else {
			counts[1]++
		}
	} else {
		counts[2]++
	}
}
