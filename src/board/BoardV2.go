package board

import "github.com/HuXin0817/dots-and-boxes/src/model"

type V2 struct {
	V1
	model.ScoreMap
}

func NewV2() *V2 {
	b := &V2{V1: *NewV1()}
	b.ScoreMap.Reset()
	return b
}

func (b *V2) Reset(nb *V1) {
	b.V1 = *nb
	b.ScoreMap.Reset()
}

func (b *V2) Add(e model.Edge) int {
	s := b.V1.Add(e)
	b.ScoreMap.Add(s)
	return s
}

func (b *V2) NotOver() bool { return b.ScoreMap.NotOver() && b.Gaming() }
