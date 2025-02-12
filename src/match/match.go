package match

import (
	"runtime"
	"sync/atomic"
)

var (
	matcher *uint64
	lock    uintptr
)

func Lock() {
	for !atomic.CompareAndSwapUintptr(&lock, 0, 1) {
		runtime.Gosched()
	}
}

func Unlock() { atomic.StoreUintptr(&lock, 0) }

func Match(id uint64) (success bool, player uint64) {
	Lock()
	defer Unlock()

	success = matcher != nil
	if success {
		player = *matcher
		matcher = nil
	} else {
		matcher = &id
	}

	return success, player
}
