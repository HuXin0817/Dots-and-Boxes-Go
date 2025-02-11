package internal

import (
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
	"github.com/HuXin0817/dots-and-boxes/src/model/container"
	"github.com/jefflund/stones/pkg/hjkl/rand"
)

type L3Model struct {
	M            Interface
	SearchTime   int
	auxBoard     board.BoardV2
	EdgeScoreMap container.EdgeScoreMap
}

func NewL3Model(searchTime int, M Interface) *L3Model {
	return &L3Model{
		M:          M,
		SearchTime: searchTime,
		auxBoard:   *board.NewBoardV2(),
	}
}

func DefaultL3Model() *L3Model { return NewL3Model(10000, NewL2Model()) }

func (m *L3Model) BestCandidateEdges(b *board.BoardV2) []model.Edge {
	if l := NewL2Model().BestCandidateEdges(b); len(l) == 1 {
		return l
	}
	m.EdgeScoreMap = container.EdgeScoreMap{}
	for range m.SearchTime/b.RemainStep() + 1 {
		m.auxBoard.Reset(&b.BoardV1)
		e := rand.Choice(m.M.BestCandidateEdges(&m.auxBoard))
		m.auxBoard.Add(e)
		for m.auxBoard.Gaming() {
			m.auxBoard.Add(rand.Choice(m.M.BestCandidateEdges(&m.auxBoard)))
		}
		m.EdgeScoreMap.Add(e, m.auxBoard.Score())
	}
	return m.EdgeScoreMap.Export()
}
