package bubble_test

import (
	"reflect"
	"testing"

	"sort/bubbleSort/bubble"
)

func TestBubbleSort(t *testing.T) {
	problem := []int{4, 1, 6, 9, 2}
	problem2 := make([]int, len(problem))
	copy(problem2, problem)
	answer := []int{1, 2, 4, 6, 9}

	compareResult := reflect.DeepEqual(answer, bubble.BubbleSort(problem))
	if !compareResult {
		t.Error("expected", true, "got", compareResult)
	}

	compareResult2 := reflect.DeepEqual(answer, bubble.ImprovedBubbleSort(problem2))
	if !compareResult {
		t.Error("expected", true, "got", compareResult2)
	}

}
