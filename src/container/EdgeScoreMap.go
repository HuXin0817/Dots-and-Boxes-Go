package container

import "github.com/HuXin0817/dots-and-boxes/src/model"

type EdgeScoreMap struct {
	time      [model.MaxEdge]int
	score     [model.MaxEdge]int
	bestEdges EdgeList
}

func (r *EdgeScoreMap) Add(edge model.Edge, score int) {
	r.time[edge]++
	r.score[edge] += score
}

func (r *EdgeScoreMap) Plus(m *EdgeScoreMap) {
	for e := range model.MaxEdge {
		r.time[e] += m.time[e]
		r.score[e] += m.score[e]
	}
}

func (r *EdgeScoreMap) Export() []model.Edge {
	var maxs float64
	for e := range model.MaxEdge {
		if r.time[e] > 0 {
			s := float64(r.score[e]) / float64(r.time[e])
			if s > maxs || r.bestEdges.Empty() {
				maxs = s
				r.bestEdges.Reset(e)
			} else if s == maxs {
				r.bestEdges.Append(e)
			}
		}
	}
	return r.bestEdges.Export()
}
