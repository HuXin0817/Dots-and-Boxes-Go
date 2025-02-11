package ai

import (
	"github.com/HuXin0817/dots-and-boxes/src/ai/internal"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/HuXin0817/dots-and-boxes/src/model/board"
	"github.com/jefflund/stones/pkg/hjkl/rand"
)

func New(name string) (func(v2 *board.BoardV2) model.Edge, error) {
	M, err := internal.NewInterface(name)
	if err != nil {
		return nil, err
	}
	return func(v2 *board.BoardV2) model.Edge {
		return rand.Choice(M.BestCandidateEdges(v2))
	}, nil
}
