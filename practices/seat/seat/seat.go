package seat

import (
	"strconv"
	"strings"
)

var SEAT_INDEX = map[rune]int{
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
	'E': 4,
	'F': 5,
	'G': 6,
	'H': 7,
	'J': 8,
	'K': 9,
}

func PossibleSeats(N int, S string) int {
	seats := make([][]int, N, N)
	for i := range seats {
		seats[i] = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	}

	if len(S) > 0 {
		seatIndexes := strings.Split(S, " ")
		for i := range seatIndexes {
			setReservedSeat(seats, seatIndexes[i])
		}
	}

	var possibleSeats int
	var val int
	for i := range seats {
		for j := range seats[i] {
			val += seats[i][j]
			if j == 2 || j == 6 || j == 9 {
				if val >= 3 {
					possibleSeats++
				}
				val = 0
			}
		}
	}

	return possibleSeats
}

func setReservedSeat(seats [][]int, S string) {
	index := []rune(S)
	lastIndex := len(index) - 1
	row, _ := strconv.Atoi(string(index[:lastIndex]))
	row--
	col := SEAT_INDEX[index[len(index)-1]]
	if col == 4 || col == 5 {
		seats[row][col] = -1
	} else {
		seats[row][col] = 0
	}
}
