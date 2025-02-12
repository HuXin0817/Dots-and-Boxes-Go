package board

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/container"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type BoardV3 struct {
	BoardV1
	scoreableEdges container.EdgeQueue
}

func NewBoardV3() *BoardV3 {
	return &BoardV3{
		BoardV1: *NewBoardV1(),
	}
}

func (b *BoardV3) Reset(nb *BoardV1) {
	b.BoardV1 = *nb
	b.scoreableEdges.Clear()
}

func (b *BoardV3) Add(e model.Edge) (s int) {
	s = b.BoardV1.Add(e)
	for _, box := range model.NearBoxes[e] {
		if b.EdgeCountOfBox[box] == 3 {
			b.scoreableEdges.Append(b.findNotContainsEdgeInBox(box))
		}
	}
	return s
}

func (b *BoardV3) MaxObtainableScore(minScore int) (s int) {
	for b.Gaming() {
		if b.scoreableEdges.Empty() {
			if e := b.findScoreableEdge(); e != -1 {
				b.scoreableEdges.Append(e)
			} else {
				break
			}
		}
		e := b.scoreableEdges.Pop()
		if b.Contains(e) {
			continue
		}
		s0 := b.Add(e)
		if config.DEBUG {
			assert.True(nil, s0 > 0)
		}
		s += s0
		if s >= minScore {
			break
		}
	}
	return s
}
