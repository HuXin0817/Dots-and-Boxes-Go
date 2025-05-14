package ai

import (
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/container"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/jefflund/stones/pkg/hjkl/rand"
)

type L3Model struct {
	M            L2Model
	SearchTime   int
	auxBoard     board.V2
	EdgeScoreMap container.EdgeScoreMap
}

func NewL3Model() *L3Model {
	return &L3Model{
		M:          *NewL2Model(),
		SearchTime: 10000,
		auxBoard:   *board.NewV2(),
	}
}

func (m *L3Model) BestCandidateEdges(b *board.V2) []model.Edge {
	if l := NewL2Model().BestCandidateEdges(b); len(l) == 1 {
		return l
	}
	m.EdgeScoreMap = container.EdgeScoreMap{}
	for range m.SearchTime/b.RemainStep() + 1 {
		m.auxBoard.Reset(&b.V1)
		e := rand.Choice(m.M.BestCandidateEdges(&m.auxBoard))
		m.auxBoard.Add(e)
		for m.auxBoard.Gaming() {
			m.auxBoard.Add(rand.Choice(m.M.BestCandidateEdges(&m.auxBoard)))
		}
		m.EdgeScoreMap.Add(e, m.auxBoard.Score())
	}
	return m.EdgeScoreMap.Export()
}
