package main

import (
	"sync/atomic"
	"testing"
)

func TestDataRaceConditions(t *testing.T) {
	var state int32
	// var mu sync.RWMutex

	// atomic value
	for i := 0; i < 10; i++ {
		go func(i int) {
			// state += int32(i)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}
}
