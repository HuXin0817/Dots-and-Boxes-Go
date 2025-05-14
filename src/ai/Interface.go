package ai

import (
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/jefflund/stones/pkg/hjkl/rand"
)

type Interface interface {
	BestCandidateEdges(*board.V2) (edges []model.Edge)
}

func New(s string) func(v2 *board.V2) model.Edge {
	var f func(v2 *board.V2) []model.Edge
	switch s {
	case "L0", "L0()":
		f = NewL0Model().BestCandidateEdges
	case "L1", "L1()":
		f = NewL1Model().BestCandidateEdges
	case "L2", "L2()":
		f = NewL2Model().BestCandidateEdges
	case "L3", "L3()":
		f = NewL3Model().BestCandidateEdges
	case "L4", "L4()":
		f = NewL4Model().BestCandidateEdges
	}
	return func(v2 *board.V2) model.Edge {
		return rand.Choice(f(v2))
	}
}
