package main

import (
	"fmt"
	"image/color"
	"runtime"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/HuXin0817/dots-and-boxes/app/assets/gen"
	"github.com/HuXin0817/dots-and-boxes/src/audio"
	"github.com/HuXin0817/dots-and-boxes/src/board"
	"github.com/HuXin0817/dots-and-boxes/src/config"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

type GameInterface struct {
	Board         board.BoardV2
	UserInputEdge model.Edge
	LastEdge      model.Edge
	End           bool
	Paused        bool
	Online        bool
	IsFirst       bool
	TimeOut       int
	Container     *fyne.Container
	Dialog        dialog.Dialog
	DotCanvases   [model.MaxDot]*canvas.Image
	EdgeCanvases  [model.MaxEdge]*canvas.Line
	EdgeAnimation [model.MaxEdge]*fyne.Animation
	EdgeLastColor [model.MaxEdge]color.NRGBA
	BoxCanvases   [model.MaxBox]*canvas.Rectangle
	BoxAnimation  [model.MaxBox]*fyne.Animation
	BoxTipped     [model.MaxBox]bool
	Buttons       [model.MaxEdge]*widget.Button
}

func NewGameInterface() *GameInterface {
	game := &GameInterface{
		Board:    *board.NewBoardV2(),
		LastEdge: -1,
	}
	fyne.CurrentApp().Settings().SetTheme(&GameTheme{
		game: game,
	})
	w := container.NewWithoutLayout()
	for b := range model.MaxBox {
		game.BoxCanvases[b] = NewBoxCanvas(b)
		w.Add(game.BoxCanvases[b])
	}
	for e := range model.MaxEdge {
		game.EdgeCanvases[e] = NewEdgeCanvas(e, EmptyEdgeColor())
		w.Add(game.EdgeCanvases[e])
	}
	for e := range model.MaxEdge {
		game.Buttons[e] = NewButtonCanvas(e, func() {
			if game.UserInputEdge == -1 {
				game.UserInputEdge = e
			}
		})
		w.Add(game.Buttons[e])
	}
	for d := range model.MaxDot {
		game.DotCanvases[d] = NewDotCanvas(d)
		w.Add(game.DotCanvases[d])
	}
	game.Container = container.New(&CenterLayout{
		Min: fyne.NewSize(EdgeWidth*model.DotsHeight+MinMargin, EdgeWidth*model.DotsWidth+MinMargin),
	}, w)
	go func() {
		<-FadeIn(700*time.Millisecond, 0, game.Container)
		game.UpdateTitle()
	}()
	MainWindow.SetContent(game.Container)
	MainWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyR:
			if game.End {
				return
			}
			if game.Dialog != nil {
				return
			}
			game.Paused = true
			game.Dialog = dialog.NewCustomConfirm("Restart?", "Yes", "No", game.ScoreLabel(), func(b bool) {
				if b {
					game.End = true
					restart(game.Online)
				}
				game.Paused = false
			}, MainWindow)
			game.Dialog.SetOnClosed(func() { game.Dialog = nil })
			game.Dialog.Show()
		case fyne.KeyZ:
			if game.End {
				return
			}
			if game.Dialog != nil {
				return
			}
			game.Paused = true
			game.Dialog = dialog.NewCustomConfirm("Return?", "Yes", "No", game.ScoreLabel(), func(b bool) {
				if b {
					game.Container.Hide()
					game.End = true
					ShowStartMenu()
				}
				game.Paused = false
			}, MainWindow)
			game.Dialog.SetOnClosed(func() { game.Dialog = nil })
			game.Dialog.Show()
		case fyne.KeySpace:
			if game.End {
				return
			}
			if game.Dialog != nil {
				game.Dialog.Hide()
				if game.Paused {
					game.Paused = false
				}
			} else {
				game.Dialog = dialog.NewCustom("Paused", "Continue", game.ScoreLabel(), MainWindow)
				game.Dialog.SetOnClosed(func() {
					game.Paused = false
				})
				game.Paused = true
				game.Dialog.SetOnClosed(func() { game.Dialog = nil })
				game.Dialog.Show()
			}
		case fyne.KeyQ:
			if game.Dialog != nil {
				return
			}
			game.Paused = true
			game.Dialog = dialog.NewCustomConfirm("Quit?", "Yes", "No", game.ScoreLabel(), func(b bool) {
				if b {
					MainWindow.Close()
				}
				game.Paused = false
			}, MainWindow)
			game.Dialog.SetOnClosed(func() { game.Dialog = nil })
			game.Dialog.Show()
		case fyne.KeyM:
			audio.On = !audio.On
		}
	})
	return game
}

func (game *GameInterface) ScoreLabel() *widget.Label {
	m := fmt.Sprintf("Player 1 Score: %2d\n", game.Board.Player1Score)
	m += fmt.Sprintf("Player 2 Score: %2d\n", game.Board.Player2Score)
	l := widget.NewLabel(m)
	l.Alignment = fyne.TextAlignCenter
	return l
}

func (game *GameInterface) CurrentBoard() *board.BoardV2 { return &game.Board }

func (game *GameInterface) GetUserEdge() model.Edge {
	game.UserInputEdge = -1
	for game.UserInputEdge == -1 {
		runtime.Gosched()
	}
	return game.UserInputEdge
}

func (game *GameInterface) RemoveHighlight() {
	e := game.LastEdge
	if e != -1 {
		game.EdgeAnimation[e].Stop()
		nc := game.EdgeLastColor[e]
		game.EdgeCanvases[e].StrokeColor = &nc
		game.EdgeCanvases[e].Refresh()
		A0 := game.EdgeCanvases[e].StrokeColor.(*color.NRGBA).A
		game.EdgeAnimation[e] = fyne.NewAnimation(300*time.Millisecond, func(f float32) {
			game.EdgeCanvases[e].StrokeColor.(*color.NRGBA).A = A0 - uint8(127.0*f)
			game.EdgeCanvases[e].Refresh()
		})
		game.EdgeAnimation[e].Start()
	}
}

func (game *GameInterface) Add(e model.Edge) {
	defer game.UpdateTitle()
	for game.Paused {
		runtime.Gosched()
		if game.End {
			return
		}
	}
	for _, box := range model.NearBoxes[e] {
		if game.Board.EdgeCountOfBox[box] == 3 {
			game.BoxAnimation[box].Stop()
			BoxStartColor := *game.BoxCanvases[box].FillColor.(*color.NRGBA)
			BoxEndColor := *BoxFilledColor(game.Board.Turn)
			game.BoxAnimation[box] = fyne.NewAnimation(600*time.Millisecond, func(f float32) {
				game.BoxCanvases[box].FillColor = InterpolationColor(BoxStartColor, BoxEndColor, f)
				game.BoxCanvases[box].Refresh()
			})
			game.BoxAnimation[box].Start()
		}
	}
	EdgeStartColor := *game.EdgeCanvases[e].StrokeColor.(*color.NRGBA)
	game.EdgeLastColor[e] = *EdgeHighlightColor(game.Board.Turn)
	game.EdgeAnimation[e] = fyne.NewAnimation(300*time.Millisecond, func(f float32) {
		game.EdgeCanvases[e].StrokeColor = InterpolationColor(EdgeStartColor, game.EdgeLastColor[e], f)
		game.EdgeCanvases[e].Refresh()
	})
	game.EdgeAnimation[e].Start()
	game.Buttons[e].Hide()
	game.RemoveHighlight()
	game.LastEdge = e
	game.EdgeCanvases[e].Refresh()
	score := game.Board.Add(e)
	if !game.End {
		if score > 0 {
			audio.Play(gen.Score)
		} else {
			audio.Play(gen.NormalMove)
		}
	}
	t := false
	for box := range model.MaxBox {
		if game.Board.EdgeCountOfBox[box] == 3 && !game.BoxTipped[box] {
			t = true
			break
		}
	}
	if t {
		for box := range model.MaxBox {
			if game.Board.EdgeCountOfBox[box] == 3 {
				if game.BoxTipped[box] {
					game.BoxAnimation[box].Stop()
				}
				BoxStartColor := *BackGroundColor()
				BoxEndColor := *BoxTipsColor()
				game.BoxAnimation[box] = fyne.NewAnimation(1200*time.Millisecond, func(f float32) {
					game.BoxCanvases[box].FillColor = InterpolationColor(BoxStartColor, BoxEndColor, f)
					game.BoxCanvases[box].Refresh()
				})
				game.BoxAnimation[box].AutoReverse = true
				game.BoxAnimation[box].RepeatCount = -1
				game.BoxAnimation[box].Start()
				game.BoxTipped[box] = true
			}
		}
	}
}

func (game *GameInterface) UpdateTitle() {
	if game.Online {
		nowStep := game.Board.Step
		var title string
		if game.Board.Turn == model.Player1Turn {
			if game.IsFirst {
				title = "Your Turn"
			} else {
				title = "Opponent's Turn"
			}
		} else {
			if game.IsFirst {
				title = "Opponent's Turn"
			} else {
				title = "Your Turn"
			}
		}
		go func() {
			t := config.PlayerTimeOut / time.Second
			for range time.Tick(time.Second) {
				if game.End || game.Board.Step != nowStep {
					return
				}
				MainWindow.SetTitle(fmt.Sprintf("%s (%ds)", title, t))
				t--
				if t == 0 || game.End {
					return
				}
			}
		}()
	} else {
		if game.Board.Turn == model.Player1Turn {
			MainWindow.SetTitle("Blue Turn")
		} else {
			MainWindow.SetTitle("Red Turn")
		}
	}
}

func (game *GameInterface) Gaming() bool {
	if game.End {
		return false
	}
	if game.TimeOut != 0 {
		return false
	}
	return game.Board.NotOver()
}

func (game *GameInterface) Close() {
	if game.End {
		return
	}
	audio.Play(gen.Win)
	game.End = true
	MainWindow.SetTitle("dots and boxes")
	game.RemoveHighlight()
	title := ""
	if game.Online {
		if game.TimeOut != 0 {
			if game.IsFirst {
				if game.TimeOut == model.Player1Turn {
					title = "Your Time Out!"
				} else if game.TimeOut == model.Player2Turn {
					title = "Opponent's Time Out!"
				}
			} else {
				if game.TimeOut == model.Player1Turn {
					title = "Opponent's Time Out!"
				} else if game.TimeOut == model.Player2Turn {
					title = "Your Time Out!"
				}
			}
		} else {
			if game.Board.Player1Score > game.Board.Player2Score {
				if game.IsFirst {
					title = "You Win!"
				} else {
					title = "You Lose."
				}
			} else if game.Board.Player1Score < game.Board.Player2Score {
				if game.IsFirst {
					title = "You Lose."
				} else {
					title = "You Win!"
				}
			} else {
				title = "Draw!"
			}
		}
	} else {
		if game.Board.Player1Score > game.Board.Player2Score {
			title = "Player 1 Win!"
		} else if game.Board.Player1Score < game.Board.Player2Score {
			title = "Player 2 Win!"
		} else {
			title = "Draw!"
		}
	}
	tapped := false
	banReturn := false
	GameOverDialog := dialog.NewCustomConfirm(title, "", "Return", game.ScoreLabel(), func(b bool) {
		if tapped {
			return
		}
		tapped = true
		if !b && !banReturn {
			ShowStartMenu()
		} else {
			restart(game.Online)
		}
	}, MainWindow)
	go func() {
		t := 5
		for t > 0 && GameOverDialog != nil {
			GameOverDialog.SetConfirmText(fmt.Sprintf("Restart (%ds)", t))
			t--
			time.Sleep(time.Second)
		}
		banReturn = true
		if GameOverDialog != nil {
			GameOverDialog.Hide()
		}
	}()
	GameOverDialog.SetOnClosed(func() {
		GameOverDialog = nil
	})
	GameOverDialog.Show()
}
