package internal

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
	"github.com/HuXin0817/dots-and-boxes/src/model/container"
	"github.com/stretchr/testify/assert"
)

type L2Model struct {
	L1          L1Model
	AuxBoard    board.BoardV2
	SearchEdges container.EdgeList
}

func NewL2Model() *L2Model {
	return &L2Model{
		L1:       *NewL1Model(),
		AuxBoard: *board.NewBoardV2(),
	}
}

func (m *L2Model) BestCandidateEdges(b *board.BoardV2) []model.Edge {
	if l := m.L1.BestCandidateEdges(b); !m.L1.L0.EnemyUnscoreableEdges.Empty() {
		return l
	}
	m.SearchEdges.Clear()
	maxs := -int(model.MaxBox + 1)
	for _, e := range b.EmptyEdges() {
		m.AuxBoard.Reset(&b.BoardV1)
		m.AuxBoard.Add(e)
		for m.AuxBoard.Gaming() {
			edge := m.L1.BestCandidateEdges(&m.AuxBoard)[0]
			if config.DEBUG {
				haveUpper1 := false
				for _, box := range model.NearBoxes[edge] {
					if m.AuxBoard.EdgeCountOfBox[box] > 1 {
						haveUpper1 = true
					}
				}
				assert.True(nil, haveUpper1)
			}
			m.AuxBoard.Add(edge)
		}
		s := m.AuxBoard.Score()
		if s > maxs {
			maxs = s
			m.SearchEdges.Reset(e)
		} else if s == maxs {
			m.SearchEdges.Append(e)
		}
	}
	return m.SearchEdges.Export()
}
