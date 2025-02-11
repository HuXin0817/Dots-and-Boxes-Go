package internal

import (
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
	"github.com/HuXin0817/dots-and-boxes/src/model/container"
	"github.com/jefflund/stones/pkg/hjkl/rand"
)

type L3Model struct {
	M          Interface
	SearchTime int
	AuxBoard   board.BoardV2
}

func NewL3Model(searchTime int, M Interface) *L3Model {
	return &L3Model{
		M:          M,
		SearchTime: searchTime,
		AuxBoard:   *board.NewBoardV2(),
	}
}

func DefaultL3Model() *L3Model { return NewL3Model(10000, NewL2Model()) }

func (m *L3Model) Searching(b *board.BoardV2) *container.EdgeScoreMap {
	var EdgeScoreMap container.EdgeScoreMap
	for range m.SearchTime/b.RemainStep() + 1 {
		m.AuxBoard.Reset(&b.BoardV1)
		e := rand.Choice(m.M.BestCandidateEdges(&m.AuxBoard))
		m.AuxBoard.Add(e)
		for m.AuxBoard.Gaming() {
			m.AuxBoard.Add(rand.Choice(m.M.BestCandidateEdges(&m.AuxBoard)))
		}
		EdgeScoreMap.Add(e, m.AuxBoard.Score())
	}
	return &EdgeScoreMap
}

func (m *L3Model) BestCandidateEdges(b *board.BoardV2) []model.Edge {
	if l := NewL2Model().BestCandidateEdges(b); len(l) == 1 {
		return l
	}
	return m.Searching(b).Export()
}
