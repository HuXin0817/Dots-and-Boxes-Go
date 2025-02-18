package board

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/container"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type V1 struct {
	V0
	container.EdgeCountOfBox
}

func NewV1() *V1 {
	return &V1{
		V0: *NewV0(),
	}
}

func (b *V1) Add(e model.Edge) int {
	b.V0.Add(e)
	return b.EdgeCountOfBox.Add(e)
}

func (b *V1) findNotContainsEdgeInBox(box model.Box) model.Edge {
	if config.DEBUG {
		assert.Equal(nil, b.EdgeCountOfBox[box], 3)
	}
	for _, e := range model.BoxEdges[box] {
		if b.NotContains(e) {
			return e
		}
	}
	panic("unreachable")
}

func (b *V1) findScoreableEdge() model.Edge {
	for box := range model.MaxBox {
		if b.EdgeCountOfBox[box] == 3 {
			return b.findNotContainsEdgeInBox(box)
		}
	}
	return -1
}
