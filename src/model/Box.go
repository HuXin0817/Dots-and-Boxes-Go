package model

import "github.com/HuXin0817/dots-and-boxes/src/config"

type Box int

const MaxBox Box = config.BoardHeight * config.BoardWidth

func newBox(x, y int) Box { return Box(x*config.BoardWidth + y) }

func (b Box) LeftTopDot() Dot {
	x := int(b / config.BoardWidth)
	y := int(b % config.BoardWidth)
	return newDot(x, y)
}

func (b Box) edges() [4]Edge {
	x := int(b / config.BoardWidth)
	y := int(b % config.BoardWidth)
	d00 := newDot(x, y)
	d01 := newDot(x+1, y)
	d10 := newDot(x, y+1)
	d11 := newDot(x+1, y+1)
	return [4]Edge{
		dotsToEdges[d00][d01],
		dotsToEdges[d00][d10],
		dotsToEdges[d10][d11],
		dotsToEdges[d01][d11],
	}
}

var BoxEdges = func() (m [MaxBox][4]Edge) {
	for box := range MaxBox {
		m[box] = box.edges()
	}
	return m
}()
