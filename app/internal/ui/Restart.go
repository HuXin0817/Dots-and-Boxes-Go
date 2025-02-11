package ui

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/x/fyne/widget"
	"github.com/HuXin0817/dots-and-boxes/app/internal/config"
	"github.com/HuXin0817/dots-and-boxes/server/api"
	"github.com/HuXin0817/dots-and-boxes/src/ai"
	"github.com/HuXin0817/dots-and-boxes/src/mock"
	"github.com/HuXin0817/dots-and-boxes/src/model"
)

var cli = api.New(config.ServerAddr)

func restart(Online bool) {
	go func() {
		Container := MainWindow.Content().(*fyne.Container)
		<-fadeOut(700*time.Millisecond, 0, Container)
		if Online {
			SpinnerCanvas, err := widget.NewAnimatedGifFromResource(SpinnerGIFResource)
			if err != nil {
				dialog.NewError(err, MainWindow).Show()
				return
			}
			SpinnerCanvas.Start()
			SpinnerCanvas.Resize(fyne.NewSize(70, 70))
			wait := true
			Size := Container.Size()
			MatchingTimeText := canvas.NewText("Matching", LinkColor)
			MatchingTimeText.Resize(fyne.NewSize(200, 50))
			MatchingTimeText.Alignment = fyne.TextAlignCenter
			MatchingTimeText.TextSize = 20
			MatchingTimeText.FontSource = TimesNewRomanBold
			Size = Container.Size()
			SpinnerCanvas.Move(fyne.NewPos(Size.Width/2-35, Size.Height/2-55))
			MatchingTimeText.Move(fyne.NewPos(Size.Height/2-100, Size.Height/2+15))
			Container.Add(SpinnerCanvas)
			Container.Add(MatchingTimeText)
			go func() {
				startTime := time.Now()
				for range time.Tick(10 * time.Millisecond) {
					if !wait {
						return
					}
					Size = Container.Size()
					SpinnerCanvas.Move(fyne.NewPos(Size.Width/2-35, Size.Height/2-55))
					MatchingTimeText.Move(fyne.NewPos(Size.Width/2-100, Size.Height/2+15))
					t := time.Since(startTime)
					MatchingTimeText.Text = fmt.Sprintf("Matching: %d min %d s", int(t.Minutes()), int(t.Seconds())%60)
					SpinnerCanvas.Refresh()
					MatchingTimeText.Refresh()
				}
			}()
			id, err := cli.StartGame()
			if err != nil {
				d := dialog.NewError(err, MainWindow)
				d.SetOnClosed(ShowStartMenu)
				d.Show()
				return
			}
			isFirst, err := cli.WaitJoin(id)
			if err != nil {
				d := dialog.NewError(err, MainWindow)
				d.SetOnClosed(ShowStartMenu)
				d.Show()
				return
			}
			wait = false
			UI := NewGameInterface()
			userAdd := func() model.Edge {
				e := UI.GetUserEdge()
				timeOut, err := cli.AddEdge(id, e)
				if err != nil {
					d := dialog.NewError(err, MainWindow)
					d.SetOnClosed(ShowStartMenu)
					d.Show()
					return 0
				}
				if timeOut != 0 {
					UI.TimeOut = timeOut
					UI.Close()
				}
				return e
			}
			UI.IsFirst = isFirst
			UI.Online = Online
			enemyAdd := func() model.Edge {
				edge, timeOut, err := cli.GetOnlinePlayerEdge(id, UI.CurrentBoard().Step)
				if err != nil {
					d := dialog.NewError(err, MainWindow)
					d.SetOnClosed(ShowStartMenu)
					d.Show()
				}
				if timeOut != 0 {
					UI.TimeOut = timeOut
					UI.Close()
				}
				return edge
			}
			if isFirst {
				go func() {
					mock.Run(UI.CurrentBoard(), UI.Gaming, UI.Add, userAdd, enemyAdd)
					UI.Close()
				}()
			} else {
				go func() {
					mock.Run(UI.CurrentBoard(), UI.Gaming, UI.Add, enemyAdd, userAdd)
					UI.Close()
				}()
			}
		} else {
			AI1, err := ai.New(config.Conf.AI1Name)
			if err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			AI2, err := ai.New(config.Conf.AI2Name)
			if err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			UI := NewGameInterface()
			func1 := UI.GetUserEdge
			func2 := UI.GetUserEdge
			if config.Conf.AI1 {
				func1 = func() model.Edge { return AI1(UI.CurrentBoard()) }
			}
			if config.Conf.AI2 {
				func2 = func() model.Edge { return AI2(UI.CurrentBoard()) }
			}
			go func() {
				mock.Run(UI.CurrentBoard(), UI.Gaming, UI.Add, func1, func2)
				UI.Close()
			}()
		}
	}()
}
