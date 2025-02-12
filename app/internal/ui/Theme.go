package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	widgetx "fyne.io/x/fyne/widget"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

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
	game    *GameInterface
	variant fyne.ThemeVariant
	fyneDefaultTheme
}

func (t *GameTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if t.variant != variant {
		t.variant = variant
		for d := range model.MaxDot {
			if t.game.DotCanvases[d] == nil {
				break
			}
			t.game.DotCanvases[d].Resource = CircleResource()
			t.game.DotCanvases[d].Refresh()
		}
		for e := range model.MaxEdge {
			if t.game.EdgeCanvases[e] == nil {
				break
			}
			if t.game.Board.NotContains(e) {
				t.game.EdgeCanvases[e].StrokeColor = EmptyEdgeColor()
			}
		}
	}

	switch name {
	case theme.ColorNameButton:
		return &color.NRGBA{}
	case theme.ColorNameBackground:
		return BackGroundColor()
	}
	return theme.DefaultTheme().Color(name, variant)
}

type StartMenuTheme struct {
	fyneDefaultTheme
	variant fyne.ThemeVariant
	gif     *widgetx.AnimatedGif
	title   *canvas.Text
}

func (t *StartMenuTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if t.variant != variant && t.gif != nil {
		t.variant = variant
		t.gif.Stop()
		if err := t.gif.LoadResource(SpinnerGIFResource()); err != nil {
			dialog.NewError(err, MainWindow).Show()
		}
		t.gif.Start()
		t.gif.Refresh()

		t.title.Color = TextColor()
		t.title.Refresh()
	}
	if name == theme.ColorNameBackground {
		return StartMenuBackgroundColor()
	}
	if name == theme.ColorNameButton {
		return EmptyEdgeColor()
	}
	return theme.DefaultTheme().Color(name, variant)
}
