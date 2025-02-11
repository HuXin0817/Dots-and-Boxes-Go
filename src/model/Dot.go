package model

import "github.com/HuXin0817/dots-and-boxes/src/config"

type Dot int

const (
	DotsHeight     = config.BoardHeight + 1
	DotsWidth      = config.BoardWidth + 1
	MaxDot     Dot = DotsHeight * DotsWidth
)

func newDot(x, y int) Dot { return Dot(x*DotsWidth + y) }

func (d Dot) X() int { return int(d) / DotsWidth }

func (d Dot) Y() int { return int(d) % DotsWidth }
