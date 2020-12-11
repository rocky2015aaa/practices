package selection

func SelectionSort(A []int) []int {
	for i := range A {
		maxIndex := i
		for j := i + 1; j < len(A); j++ {
			if A[j] > A[maxIndex] {
				maxIndex = j
			}
		}
		A[i], A[maxIndex] = A[maxIndex], A[i]
	}

	return A
}
