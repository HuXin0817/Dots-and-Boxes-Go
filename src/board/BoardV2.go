package board

import "github.com/HuXin0817/dots-and-boxes/src/model"

type BoardV2 struct {
	BoardV1
	model.ScoreMap
}

func NewBoardV2() *BoardV2 {
	b := &BoardV2{BoardV1: *NewBoardV1()}
	b.ScoreMap.Reset()
	return b
}

func (b *BoardV2) Reset(nb *BoardV1) {
	b.BoardV1 = *nb
	b.ScoreMap.Reset()
}

func (b *BoardV2) Add(e model.Edge) int {
	s := b.BoardV1.Add(e)
	b.ScoreMap.Add(s)
	return s
}

func (b *BoardV2) NotOver() bool { return b.ScoreMap.NotOver() && b.Gaming() }
