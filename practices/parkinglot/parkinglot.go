package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(CheckParkingFee("10:00", "13:21"))
	fmt.Println(CheckParkingFee("09:42", "11:42"))
}

func CheckParkingFee(E, L string) int {
	eTimeStr := strings.Split(E, ":")
	lTimeStr := strings.Split(L, ":")

	eTimeHour, _ := strconv.Atoi(eTimeStr[0])
	eTimeMin, _ := strconv.Atoi(eTimeStr[1])
	lTimeHour, _ := strconv.Atoi(lTimeStr[0])
	lTimeMin, _ := strconv.Atoi(lTimeStr[1])

	hourDiff := lTimeHour - eTimeHour
	if eTimeMin <= lTimeMin && (lTimeMin-eTimeMin) != 0 {
		hourDiff++
	}

	return (2 + 3*1 + 4*(hourDiff-1))
}
