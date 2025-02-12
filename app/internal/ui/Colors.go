package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

var (
	LightThemeGameInterfaceBackGroundColor = &color.NRGBA{R: 242, G: 242, B: 242, A: 255}
	DarkThemeGameInterfaceBackGroundColor  = &color.NRGBA{R: 43, G: 43, B: 43, A: 255}
	LightThemeEmptyEdgeColor               = &color.NRGBA{R: 217, G: 217, B: 217, A: 255}
	DarkThemeEmptyEdgeColor                = &color.NRGBA{R: 65, G: 65, B: 65, A: 255}
	Player1EdgeHighlightColor              = &color.NRGBA{R: 64, G: 64, B: 255, A: 255}
	Player2EdgeHighlightColor              = &color.NRGBA{R: 255, G: 64, B: 64, A: 255}
	LightThemeScoreableTipsColor           = &color.NRGBA{R: 250, G: 250, B: 200, A: 255}
	DarkThemeScoreableTipsColor            = &color.NRGBA{R: 65, G: 65, B: 15, A: 255}
	Player1BoxFilledColor                  = &color.NRGBA{R: 64, G: 64, B: 255, A: 64}
	Player2BoxFilledColor                  = &color.NRGBA{R: 255, G: 64, B: 64, A: 64}
	DarkThemeStartMenuBackgroundColor      = &color.NRGBA{R: 16, G: 16, B: 16, A: 255}
	LightThemeStartMenuBackgroundColor     = &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	LinkColor                              = &color.NRGBA{R: 73, G: 148, B: 236, A: 255}
)

func InterpolationColor(start, end color.NRGBA, f float32) *color.NRGBA {
	return &color.NRGBA{
		R: uint8(float32(start.R) + f*(float32(end.R)-float32(start.R))),
		G: uint8(float32(start.G) + f*(float32(end.G)-float32(start.G))),
		B: uint8(float32(start.B) + f*(float32(end.B)-float32(start.B))),
		A: uint8(float32(start.A) + f*(float32(end.A)-float32(start.A))),
	}
}

func GetColor(lightColor, darkColor *color.NRGBA) *color.NRGBA {
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		return darkColor
	} else {
		return lightColor
	}
}

func BackGroundColor() *color.NRGBA {
	return GetColor(LightThemeGameInterfaceBackGroundColor, DarkThemeGameInterfaceBackGroundColor)
}

func EmptyEdgeColor() *color.NRGBA {
	return GetColor(LightThemeEmptyEdgeColor, DarkThemeEmptyEdgeColor)
}

func BoxTipsColor() *color.NRGBA {
	return GetColor(LightThemeScoreableTipsColor, DarkThemeScoreableTipsColor)
}

func StartMenuBackgroundColor() *color.NRGBA {
	return GetColor(LightThemeStartMenuBackgroundColor, DarkThemeStartMenuBackgroundColor)
}

func TextColor() color.Color {
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		return color.White
	} else {
		return color.Black
	}
}

func BoxFilledColor(turn int) *color.NRGBA {
	if turn == model.Player1Turn {
		return Player1BoxFilledColor
	} else {
		return Player2BoxFilledColor
	}
}

func EdgeHighlightColor(turn int) *color.NRGBA {
	if turn == model.Player1Turn {
		return Player1EdgeHighlightColor
	} else {
		return Player2EdgeHighlightColor
	}
}
