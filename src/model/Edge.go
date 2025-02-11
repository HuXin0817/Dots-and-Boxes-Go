package model

import "github.com/HuXin0817/dots-and-boxes/src/config"

type Edge int

const MaxEdge Edge = config.BoardHeight*(config.BoardWidth+1) + (config.BoardHeight+1)*config.BoardWidth

func edgeDotMapper() ([MaxDot][MaxDot]Edge, [MaxEdge]Dot, [MaxEdge]Dot) {
	var dTe [MaxDot][MaxDot]Edge
	var eTd1, eTd2 [MaxEdge]Dot
	e := Edge(0)
	for x := range DotsHeight {
		for y := range DotsWidth {
			d1 := newDot(x, y)
			if x1 := x + 1; x1 < DotsHeight {
				d2 := newDot(x1, y)
				dTe[d1][d2] = e
				eTd1[e] = d1
				eTd2[e] = d2
				e++
			}
			if y1 := y + 1; y1 < DotsWidth {
				d2 := newDot(x, y1)
				dTe[d1][d2] = e
				eTd1[e] = d1
				eTd2[e] = d2
				e++
			}
		}
	}
	return dTe, eTd1, eTd2
}

var dotsToEdges, edgeToDot1, edgeToDot2 = edgeDotMapper()

func (e Edge) Dot1() Dot { return edgeToDot1[e] }

func (e Edge) Dot2() Dot { return edgeToDot2[e] }

func (e Edge) nearBoxes() []Box {
	var b []Box
	x := e.Dot2().X() - 1
	y := e.Dot2().Y() - 1
	if x >= 0 && y >= 0 {
		b = append(b, newBox(x, y))
	}
	x = e.Dot1().X()
	y = e.Dot1().Y()
	if x < config.BoardHeight && y < config.BoardWidth {
		b = append(b, newBox(x, y))
	}
	return b
}

var NearBoxes = func() (b [MaxEdge][]Box) {
	for e := range MaxEdge {
		b[e] = e.nearBoxes()
	}
	return b
}()
