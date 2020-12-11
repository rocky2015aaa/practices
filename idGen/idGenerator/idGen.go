package idGenerator

import (
	"errors"
	"sync"
	"time"
)

type IdGenerator struct {
	time     int64 // 42 bit
	shardId  int64 // 10 bit
	sequence int64 // 12 bit
	lock     *sync.Mutex
}

// Id Generator
func New(shardId int64) (id *IdGenerator, err error) {
	if shardId < 0 && shardId > 1023 { // the shard id rage is from 0 to 1023
		err = errors.New("Shard id is not valid")
		return nil, err
	}
	id = &IdGenerator{
		shardId: shardId,
		lock: new(sync.Mutex),
	}

	return id, nil
}

func (id *IdGenerator) SetTime() int64 {
	return time.Now().UnixNano() / 1000000 // current time in milli second
}

func (id *IdGenerator) GetNextId() int64 {
	id.lock.Lock()
	defer id.lock.Unlock()

	var time int64
	for {
		time = id.SetTime()
		if time == id.time {
			id.sequence = (id.sequence + 1) & 0XFFF // 'AND' op with 12 bit '1'
			if id.sequence == 0 {
				continue
			}
			break
		} else if time > id.time {
			id.sequence = 0
			break
		}
	}
	id.time = time

	return ((time-1477872000000)<<22 | id.shardId<<12 | id.sequence)
	// 10/31/2016 9:00:00(41 bit), time bit move, shard id move
}
