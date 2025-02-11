package internal

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
	"github.com/stretchr/testify/assert"
)

type L1Model struct {
	L0       L0Model
	AuxBoard board.BoardV3
}

func NewL1Model() *L1Model {
	return &L1Model{
		L0:       *NewL0Model(),
		AuxBoard: *board.NewBoardV3(),
	}
}

func (m *L1Model) BestCandidateEdges(b *board.BoardV2) []model.Edge {
	if l := m.L0.BestCandidateEdges(b); !m.L0.EnemyUnscoreableEdges.Empty() || !m.L0.ScoreableEdge.Empty() {
		return l
	}
	mins := int(model.MaxBox + 1)
	Candidate := &m.L0.EnemyUnscoreableEdges
	if config.DEBUG {
		assert.True(nil, Candidate.Empty())
	}
	for _, e := range b.EmptyEdges() {
		m.AuxBoard.Reset(&b.BoardV1)
		if config.DEBUG {
			assert.Equal(nil, m.AuxBoard.Add(e), 0)
		} else {
			m.AuxBoard.Add(e)
		}
		s := m.AuxBoard.MaxObtainableScore(mins)
		if s < mins {
			mins = s
			Candidate.Reset(e)
		} else if s == mins {
			Candidate.Append(e)
		}
	}
	return Candidate.Export()
}
