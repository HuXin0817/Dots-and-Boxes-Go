package container

import (
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

type EdgeCountOfBox [model.MaxBox]int

func (b *EdgeCountOfBox) Add(e model.Edge) (s int) {
	for _, box := range model.NearBoxes[e] {
		(*b)[box]++
		if (*b)[box] == 4 {
			s++
		}
	}
	return s
}
