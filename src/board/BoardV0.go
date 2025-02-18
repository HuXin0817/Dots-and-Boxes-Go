package board

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type V0 struct {
	model.Step
	edges [model.MaxEdge]model.Edge
	index [model.MaxEdge]int
}

func NewV0() *V0 {
	b := &V0{}
	for i := range model.MaxEdge {
		b.index[i] = int(i)
		b.edges[i] = i
	}
	return b
}

func (b *V0) Add(e model.Edge) {
	if config.DEBUG {
		assert.True(nil, b.NotContains(e))
	}
	r := b.edges[b.Step]
	ie := b.index[e]
	ir := int(b.Step)
	b.edges[ie], b.edges[ir] = b.edges[ir], b.edges[ie]
	b.index[e], b.index[r] = ir, ie
	b.Step++
}

func (b *V0) Contains(e model.Edge) bool { return b.index[e] < int(b.Step) }

func (b *V0) NotContains(e model.Edge) bool { return b.index[e] >= int(b.Step) }

func (b *V0) EmptyEdges() []model.Edge { return b.edges[b.Step:] }

func (b *V0) MoveRecord() []model.Edge { return b.edges[:b.Step] }
