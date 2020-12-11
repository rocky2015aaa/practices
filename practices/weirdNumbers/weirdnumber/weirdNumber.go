package weirdnumber

func CheckWeirdNumber(n int) string {
	divisorsChecker := make(map[int]bool)
	half := n / 2
	divisors := []int{}
	divisors2 := []int{}
	for i := 1; i < half; i++ {
		if n%i == 0 {
			if _, ok := divisorsChecker[i]; !ok {
				divisors = append(divisors, i)
				divisorsChecker[i] = true
			} else {
				continue
			}

			q := n / i
			if _, ok := divisorsChecker[q]; !ok {
				divisors2 = append([]int{q}, divisors2...)
				divisorsChecker[q] = true
			}
		}
	}

	divisors = append(divisors, divisors2...)
	leng := (uint32(len(divisors)) - 1)
	for j := 0; j < (1 << leng); j++ {
		var sum int
		var k uint32
		for k = 0; k < leng; k++ {
			if (j & (1 << k)) > 0 {
				sum += divisors[k]
				if sum == n {
					return "not weird"
				}
				if sum > n {
					break
				}

			}
		}
	}

	return "weird"
}
