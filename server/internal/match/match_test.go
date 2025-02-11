package match

import (
	"sync"
	"testing"
)

func BenchmarkMatch(b *testing.B) {
	const N = uint64(1e9)

	for i := range N {
		Match(i)
	}
}

func BenchmarkMatchParallel(b *testing.B) {
	const T = 8
	const N = 1e8

	var wg sync.WaitGroup
	wg.Add(T)
	for i := range T {
		go func(i uint64) {
			defer wg.Done()
			for j := uint64(0); j < N; j += T {
				Match(i + j)
			}
		}(uint64(i))
	}
	wg.Wait()
}
