package ui

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

const (
	EdgeWidth    = 80
	DotWidth     = EdgeWidth / 5
	HalfDotWidth = DotWidth / 2
	BoxWidth     = EdgeWidth - DotWidth
	MinMargin    = EdgeWidth / 3 * 2
)

func convertPosition(x int) float32 { return MinMargin + float32(x)*EdgeWidth }

func NewDotCanvas(d model.Dot) *canvas.Image {
	r := canvas.NewImageFromResource(CircleResource())
	r.Resize(fyne.NewSize(DotWidth*math.Sqrt2, DotWidth*math.Sqrt2))
	x := convertPosition(d.X()) - DotWidth*(math.Sqrt2-1)/2
	y := convertPosition(d.Y()) - DotWidth*(math.Sqrt2-1)/2
	p := fyne.NewPos(x, y)
	r.Move(p)
	return r
}

func NewEdgeCanvas(e model.Edge, c color.Color) *canvas.Line {
	x1 := convertPosition(e.Dot1().X()) + HalfDotWidth
	y1 := convertPosition(e.Dot1().Y()) + HalfDotWidth
	x2 := convertPosition(e.Dot2().X()) + HalfDotWidth
	y2 := convertPosition(e.Dot2().Y()) + HalfDotWidth
	nc := *c.(*color.NRGBA)
	l := canvas.NewLine(&nc)
	l.Position1 = fyne.NewPos(x1, y1)
	l.Position2 = fyne.NewPos(x2, y2)
	l.StrokeWidth = DotWidth
	return l
}

func NewBoxCanvas(b model.Box) *canvas.Rectangle {
	d := b.LeftTopDot()
	x := convertPosition(d.X()) + DotWidth - 1
	y := convertPosition(d.Y()) + DotWidth - 1
	r := canvas.NewRectangle(&color.NRGBA{})
	p := fyne.NewPos(x, y)
	r.Move(p)
	s := fyne.NewSize(BoxWidth+1, BoxWidth+1)
	r.Resize(s)
	return r
}

func NewButtonCanvas(e model.Edge, tapped func()) *widget.Button {
	b := widget.NewButton("", tapped)
	var s fyne.Size
	if e.Dot1().X() == e.Dot2().X() {
		s = fyne.NewSize(DotWidth, EdgeWidth)
	} else {
		s = fyne.NewSize(EdgeWidth, DotWidth)
	}
	b.Resize(s)
	x1 := convertPosition(e.Dot1().X())
	x2 := convertPosition(e.Dot2().X())
	y1 := convertPosition(e.Dot1().Y())
	y2 := convertPosition(e.Dot2().Y())
	mx := (x1 + x2) / 2
	my := (y1 + y2) / 2
	x := mx - s.Width/2 + HalfDotWidth
	y := my - s.Height/2 + HalfDotWidth
	p := fyne.NewPos(x, y)
	b.Move(p)
	return b
}
