package max

func Solution(A []int) int {
	N := len(A)
	var result int
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if A[i] != A[j] {
				if result < (j - i) {
					result = (j - i)
				}
			}
		}
	}
	return result
}

func Solution2(A []int) int {
	N := len(A) - 1
	var result int
	for i := 0; i <= N/2; i++ {
		for j := N; j >= N/2; j-- {
			if A[i] != A[j] {
				if result < (j - i) {
					result = (j - i)
				}
			}
		}
	}
	return result
}
