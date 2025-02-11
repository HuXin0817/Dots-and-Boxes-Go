package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var ThemeVariant = fyne.CurrentApp().Settings().ThemeVariant()

type fyneDefaultTheme struct{}

func (t *fyneDefaultTheme) Icon(icon fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(icon)
}

func (t *fyneDefaultTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (t *fyneDefaultTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

type GameTheme struct {
	fyneDefaultTheme
}

func (t *GameTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameButton:
		return &color.NRGBA{}
	case theme.ColorNameBackground:
		return BackGroundColor
	}
	return theme.DefaultTheme().Color(name, variant)
}

type StartMenuTheme struct {
	fyneDefaultTheme
}

func (s *StartMenuTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return StartMenuBackgroundColor
	}
	if name == theme.ColorNameButton {
		return EmptyEdgeColor
	}
	return theme.DefaultTheme().Color(name, variant)
}
