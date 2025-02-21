package ai

import (
	"runtime"
	"sync"

	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/container"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

type L4Model []L3Model

func NewL4Model(newM func() *L3Model) *L4Model {
	var m L4Model
	for range runtime.NumCPU() {
		m = append(m, *newM())
	}
	return &m
}

func DefaultL4Model() *L4Model { return NewL4Model(DefaultL3Model) }

func (m *L4Model) BestCandidateEdges(b *board.V2) []model.Edge {
	if l := NewL2Model().BestCandidateEdges(b); len(l) == 1 {
		return l
	}
	T := len(*m) - 1
	var wg sync.WaitGroup
	wg.Add(T)
	for i := range T {
		go func(i int) {
			(*m)[i].BestCandidateEdges(b)
			wg.Done()
		}(i)
	}
	(*m)[T].BestCandidateEdges(b)
	wg.Wait()
	var r container.EdgeScoreMap
	for i := range *m {
		r.Plus(&(*m)[i].EdgeScoreMap)
	}
	return r.Export()
}
