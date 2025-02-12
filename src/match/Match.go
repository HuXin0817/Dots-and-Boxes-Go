package match

import (
	"runtime"
	"sync/atomic"
)

var (
	matcher *uint64
	l       uintptr
)

func lock() {
	for !atomic.CompareAndSwapUintptr(&l, 0, 1) {
		runtime.Gosched()
	}
}

func unlock() { atomic.StoreUintptr(&l, 0) }

func Match(id uint64) (success bool, player uint64) {
	lock()
	defer unlock()

	success = matcher != nil
	if success {
		player = *matcher
		matcher = nil
	} else {
		matcher = &id
	}

	return success, player
}
