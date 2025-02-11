package mock

import (
	"testing"
)

func BenchmarkL4(b *testing.B) {
	b.Log(RunAILocalMParallel("L4", "L4", 1, 1))
}
