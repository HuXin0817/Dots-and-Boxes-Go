package model

import "github.com/HuXin0817/dots-and-boxes/src/config"

type Dot int

const (
	DotSize     = config.BoardSize + 1
	MaxDot  Dot = DotSize * DotSize
)

func newDot(x, y int) Dot { return Dot(x*DotSize + y) }

func (d Dot) X() int { return int(d) / DotSize }

func (d Dot) Y() int { return int(d) % DotSize }
