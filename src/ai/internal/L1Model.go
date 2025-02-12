package internal

import (
	board2 "github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type L1Model struct {
	L0       L0Model
	auxBoard board2.BoardV3
}

func NewL1Model() *L1Model {
	return &L1Model{
		L0:       *NewL0Model(),
		auxBoard: *board2.NewBoardV3(),
	}
}

func (m *L1Model) BestCandidateEdges(b *board2.BoardV2) []model.Edge {
	if l := m.L0.BestCandidateEdges(b); !m.L0.EnemyUnscoreableEdges.Empty() || !m.L0.ScoreableEdge.Empty() {
		return l
	}
	mins := int(model.MaxBox + 1)
	Candidate := &m.L0.EnemyUnscoreableEdges
	if config.DEBUG {
		assert.True(nil, Candidate.Empty())
	}
	for _, e := range b.EmptyEdges() {
		m.auxBoard.Reset(&b.BoardV1)
		if config.DEBUG {
			assert.Equal(nil, m.auxBoard.Add(e), 0)
		} else {
			m.auxBoard.Add(e)
		}
		s := m.auxBoard.MaxObtainableScore(mins)
		if s < mins {
			mins = s
			Candidate.Reset(e)
		} else if s == mins {
			Candidate.Append(e)
		}
	}
	return Candidate.Export()
}
