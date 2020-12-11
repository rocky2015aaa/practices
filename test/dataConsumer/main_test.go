package main

import "testing"

func TestAverage(t *testing.T) {
	numberLists := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 3, 5, 7, 9},
		[]int{10, 20, 30, 40, 50, 60, 70},
	}

	averages := []float64{3.0, 5.0, 40.0}

	for i, numberList := range numberLists {
		average := calculator(numberList, len(numberList))
		if average != averages[i] {
			t.Errorf("Expected %.2f, got %.2f\n", average, averages[i])
		}
	}
}
