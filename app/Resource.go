package main

import (
	"bytes"
	_ "embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	svg "github.com/ajstarks/svgo"
)

var (
	//go:embed "assets/font/Times New Roman Bold.ttf"
	TimesNewRomanBold []byte
	//go:embed "assets/font/Times New Roman Italic.ttf"
	TimesNewRomanItalic []byte

	//go:embed "assets/icon/icon.png"
	Icon []byte
	//go:embed "assets/icon/spinner_dark.gif"
	SpinnerDark []byte
	//go:embed "assets/icon/spinner_light.gif"
	SpinnerLight []byte

	//go:embed "assets/music/EnterGame.MP3"
	EnterGameMusic []byte
	//go:embed "assets/music/NormalMove.MP3"
	NormalMoveMusic []byte
	//go:embed "assets/music/Score.MP3"
	ScoreMusic []byte
	//go:embed "assets/music/TouchButton.MP3"
	TouchButtonMusic []byte
	//go:embed "assets/music/Win.MP3"
	WinMusic []byte
)

var (
	IconResource                = fyne.NewStaticResource("Icon", Icon)
	TimesNewRomanItalicResource = fyne.NewStaticResource("TimesNewRomanItalic", TimesNewRomanItalic)
	TimesNewRomanBoldResource   = fyne.NewStaticResource("TimesNewRomanBold", TimesNewRomanBold)
	SpinnerDarkResource         = fyne.NewStaticResource("SpinnerDark", SpinnerDark)
	SpinnerLightResource        = fyne.NewStaticResource("SpinnerLight", SpinnerLight)
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
