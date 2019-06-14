package main

import (
	"runtime"
	"sync/atomic"
)

type spinLock struct {
	f     uint32
	i     int
	sched bool
}

func (sl *spinLock) lock() {
	for !sl.tryLock() {
		if sl.sched {
			runtime.Gosched() //allow other goroutines to do stuff.
		}

		continue
	}
}

func (sl *spinLock) unlock() {
	atomic.StoreUint32(&sl.f, 0)
}

func (sl *spinLock) tryLock() bool {
	return atomic.CompareAndSwapUint32(&sl.f, 0, 1)
}

func (sl *spinLock) add() {
	sl.lock()
	sl.i++
	sl.unlock()
}
