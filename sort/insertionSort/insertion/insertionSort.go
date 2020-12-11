package insertion

func InsertionSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		temp := A[i]
		aux := i - 1
		for {
			if aux >= 0 && A[aux] < temp {
				A[aux+1] = A[aux]
				aux--
			} else {
				break
			}
		}
		A[aux+1] = temp
	}

	return A
}
