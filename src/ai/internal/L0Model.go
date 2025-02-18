package internal

import (
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/container"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

type L0Model struct {
	EnemyUnscoreableEdges container.EdgeList
	ScoreableEdge         container.EdgeList
}

func NewL0Model() *L0Model {
	return &L0Model{}
}

func (m *L0Model) BestCandidateEdges(b *board.V2) []model.Edge {
	m.ScoreableEdge.Clear()
	m.EnemyUnscoreableEdges.Clear()
	for _, edge := range b.EmptyEdges() {
		maxCount := 0
		for _, box := range model.NearBoxes[edge] {
			maxCount = max(maxCount, b.EdgeCountOfBox[box])
		}
		if maxCount == 3 {
			m.ScoreableEdge.Append(edge)
		} else if maxCount < 2 {
			m.EnemyUnscoreableEdges.Append(edge)
		}
	}
	if !m.ScoreableEdge.Empty() {
		return m.ScoreableEdge.Export()
	} else if !m.EnemyUnscoreableEdges.Empty() {
		return m.EnemyUnscoreableEdges.Export()
	} else {
		return b.EmptyEdges()
	}
}
