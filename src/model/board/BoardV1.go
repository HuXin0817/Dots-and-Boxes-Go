package board

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/container"
	"github.com/stretchr/testify/assert"
)

type BoardV1 struct {
	BoardV0
	container.EdgeCountOfBox
}

func NewBoardV1() *BoardV1 {
	return &BoardV1{
		BoardV0: *NewBoardV0(),
	}
}

func (b *BoardV1) Add(e model.Edge) int {
	b.BoardV0.Add(e)
	return b.EdgeCountOfBox.Add(e)
}

func (b *BoardV1) findNotContainsEdgeInBox(box model.Box) model.Edge {
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

func (b *BoardV1) findScoreableEdge() model.Edge {
	for box := range model.MaxBox {
		if b.EdgeCountOfBox[box] == 3 {
			return b.findNotContainsEdgeInBox(box)
		}
	}
	return -1
}
