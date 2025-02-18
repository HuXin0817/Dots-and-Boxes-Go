package container

import (
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
	"github.com/stretchr/testify/assert"
)

type EdgeCountOfBox [model.MaxBox]int

func (b *EdgeCountOfBox) Add(e model.Edge) (s int) {
	for _, box := range model.NearBoxes[e] {
		(*b)[box]++
		if config.DEBUG {
			assert.True(nil, (*b)[box] <= 4)
		}
		if (*b)[box] == 4 {
			s++
		}
	}
	return s
}

func (b *EdgeCountOfBox) MaxCount(e model.Edge) (c int) {
	for _, box := range model.NearBoxes[e] {
		c = max(c, b[box])
	}
	return c
}
