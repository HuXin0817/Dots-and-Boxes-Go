package ai

import (
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/container"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type L2Model struct {
	L1          L1Model
	auxBoard    board.V2
	SearchEdges container.EdgeList
}

func NewL2Model() *L2Model {
	return &L2Model{
		L1:       *NewL1Model(),
		auxBoard: *board.NewV2(),
	}
}

func (m *L2Model) BestCandidateEdges(b *board.V2) []model.Edge {
	if l := m.L1.BestCandidateEdges(b); !m.L1.L0.EnemyUnscoreableEdges.Empty() {
		return l
	}
	m.SearchEdges.Clear()
	maxs := -int(model.MaxBox + 1)
	for _, e := range b.EmptyEdges() {
		m.auxBoard.Reset(&b.V1)
		m.auxBoard.Add(e)
		for m.auxBoard.Gaming() {
			edge := m.L1.BestCandidateEdges(&m.auxBoard)[0]
			if config.DEBUG {
				assert.True(nil, b.MaxCount(edge) > 1)
			}
			m.auxBoard.Add(edge)
		}
		if s := m.auxBoard.Score(); s > maxs {
			maxs = s
			m.SearchEdges.Reset(e)
		} else if s == maxs {
			m.SearchEdges.Append(e)
		}
	}
	return m.SearchEdges.Export()
}
