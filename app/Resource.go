package main

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/HuXin0817/dots-and-boxes/app/assets/gen"
	svg "github.com/ajstarks/svgo"
)

var (
	IconResource                = fyne.NewStaticResource("Icon", gen.Icon)
	TimesNewRomanItalicResource = fyne.NewStaticResource("TimesNewRomanItalic", gen.TimesNewRomanItalic)
	TimesNewRomanBoldResource   = fyne.NewStaticResource("TimesNewRomanBold", gen.TimesNewRomanBold)
	SpinnerDarkResource         = fyne.NewStaticResource("SpinnerDark", gen.SpinnerDark)
	SpinnerLightResource        = fyne.NewStaticResource("SpinnerLight", gen.SpinnerLight)
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
		return SpinnerDarkResource
	} else {
		return SpinnerLightResource
	}
}
