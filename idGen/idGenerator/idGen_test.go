package idGenerator

import (
	"testing"
)

const INTERVAL = 10

func TestSnowFlake(t *testing.T) {
	idgenerator, _ := New(1)

	firstId := idgenerator.GetNextId()
	var lastId int64

	for i := 0; i < INTERVAL; i++ {
		lastId = idgenerator.GetNextId()
	}

	expectedId := firstId + INTERVAL
	if expectedId != lastId {
		t.Error("Expected", expectedId, "Got", lastId)
	}
}
