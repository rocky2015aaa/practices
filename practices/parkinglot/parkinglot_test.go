package main

import (
	"testing"
)

func TestParkingLot(t *testing.T) {
	result := CheckParkingFee("10:00", "13:21")
	if result != 17 {
		t.Error("expected", 17, "got", result)
	}
}
