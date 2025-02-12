package ui

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/HuXin0817/dots-and-boxes/app/internal/gen"
	svg "github.com/ajstarks/svgo"
)

var (
	Icon                = fyne.NewStaticResource("Icon", gen.Icon)
	TimesNewRomanItalic = fyne.NewStaticResource("TimesNewRomanItalic", gen.TimesNewRomanItalic)
	TimesNewRomanBold   = fyne.NewStaticResource("TimesNewRomanBold", gen.TimesNewRomanBold)
	SpinnerDark         = fyne.NewStaticResource("SpinnerDark", gen.SpinnerDark)
	SpinnerLight        = fyne.NewStaticResource("SpinnerLight", gen.SpinnerLight)
)

func CircleResource() fyne.Resource {
	var buf bytes.Buffer
	canvas := svg.New(&buf)
	canvas.Start(200, 200)
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		canvas.Circle(100, 100, 80, "fill:rgb(202, 202, 202)")
	} else {
		canvas.Circle(100, 100, 80, "fill:rgb(255, 255, 255)")
	}
	canvas.End()
	return fyne.NewStaticResource("Icon", buf.Bytes())
}

func SpinnerGIFResource() fyne.Resource {
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		return SpinnerDark
	} else {
		return SpinnerLight
	}
}
