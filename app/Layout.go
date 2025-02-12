package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type CenterLayout struct {
	Min fyne.Size
}

func (l *CenterLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	px := (size.Width - l.Min.Width) / 2
	py := (size.Height - l.Min.Height) / 2
	for _, w := range objects {
		if c, ok := w.(*fyne.Container); ok {
			c.Move(fyne.NewPos(px, py))
			c.Refresh()
		} else if r, ok := w.(*canvas.Rectangle); ok {
			r.Resize(size)
			r.Move(fyne.NewPos(0, 0))
		}
	}
}

func (l *CenterLayout) MinSize([]fyne.CanvasObject) fyne.Size { return l.Min }
