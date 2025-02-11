package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	widgetx "fyne.io/x/fyne/widget"
	"github.com/HuXin0817/dots-and-boxes/app/internal/audio"
	"github.com/HuXin0817/dots-and-boxes/app/internal/config"
	"github.com/HuXin0817/dots-and-boxes/app/internal/gen"
	"github.com/HuXin0817/dots-and-boxes/src/ai"
)

func ShowStartMenu() {
	MainWindow.SetTitle("Dots and Boxes")
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{})
	Container := container.NewWithoutLayout()
	MenuSize := fyne.NewSize(612, 600)
	Container.Resize(MenuSize)
	IconCanvas := canvas.NewImageFromResource(Icon)
	IconCanvas.Resize(fyne.NewSize(200, 200))
	IconCanvas.Move(fyne.NewPos(206, 50))
	Container.Add(IconCanvas)
	TitleText := canvas.NewText("Dots and Boxes", TextColor())
	TitleText.Alignment = fyne.TextAlignCenter
	TitleText.TextSize = 26
	TitleText.FontSource = TimesNewRomanBold
	TitleText.Resize(fyne.NewSize(200, 50))
	TitleText.Move(fyne.NewPos(206, 250))
	Container.Add(TitleText)
	GameLink := canvas.NewText("https://github.com/HuXin0817/Dots-and-Boxes", LinkColor)
	GameLink.Resize(fyne.NewSize(200, 50))
	GameLink.Alignment = fyne.TextAlignCenter
	GameLink.TextSize = 17
	GameLink.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	GameLink.FontSource = TimesNewRomanItalic
	GameLink.Move(fyne.NewPos(206, 285))
	Container.Add(GameLink)
	SpinnerCanvas, err := widgetx.NewAnimatedGifFromResource(SpinnerGIFResource)
	if err != nil {
		dialog.NewError(err, MainWindow).Show()
	}
	SpinnerCanvas.Start()
	SpinnerCanvas.Resize(fyne.NewSize(70, 70))
	SpinnerCanvas.Move(fyne.NewPos(271, 335))
	Container.Add(SpinnerCanvas)
	started := false
	PlayOnlineButton := widget.NewButton("Play Online", func() {
		if started {
			return
		}
		started = true
		audio.Play(gen.TouchButton)
		restart(true)
	})
	PlayOnlineButton.Resize(fyne.NewSize(200, 50))
	PlayOnlineButton.Move(fyne.NewPos(206, 435))
	Container.Add(PlayOnlineButton)
	PlayLocalButton := widget.NewButton("Play Local", func() {
		TmpConfig := config.Conf
		Player1RadioGroup := widget.NewRadioGroup([]string{
			"People",
			"Computer",
		}, func(selected string) {
			if selected == "People" {
				TmpConfig.AI1 = false
			} else if selected == "Computer" {
				TmpConfig.AI1 = true
			}
		})
		if config.Conf.AI1 {
			Player1RadioGroup.SetSelected("Computer")
		} else {
			Player1RadioGroup.SetSelected("People")
		}
		Player2RadioGroup := widget.NewRadioGroup([]string{
			"People",
			"Computer",
		}, func(selected string) {
			if selected == "People" {
				TmpConfig.AI2 = false
			} else if selected == "Computer" {
				TmpConfig.AI2 = true
			}
		})
		if config.Conf.AI2 {
			Player2RadioGroup.SetSelected("Computer")
		} else {
			Player2RadioGroup.SetSelected("People")
		}
		Entry1 := widget.NewEntry()
		Entry1.SetText(config.Conf.AI1Name)
		Entry2 := widget.NewEntry()
		Entry2.SetText(config.Conf.AI2Name)
		PlayLocalButtonSettingDialog := dialog.NewForm("Play Local", "Start", "Cancel", []*widget.FormItem{
			widget.NewFormItem("Player1:", Player1RadioGroup),
			widget.NewFormItem("Player2:", Player2RadioGroup),
			widget.NewFormItem("AI1:", Entry1),
			widget.NewFormItem("AI2:", Entry2),
		}, func(b bool) {
			if !b {
				return
			}
			if _, err := ai.New(Entry1.Text); err != nil {
				dialog.NewError(err, MainWindow).Show()
				return
			} else {
				TmpConfig.AI1Name = Entry1.Text
			}
			if _, err := ai.New(Entry2.Text); err != nil {
				dialog.NewError(err, MainWindow).Show()
				return
			} else {
				TmpConfig.AI2Name = Entry2.Text
			}
			config.Conf = TmpConfig
			if err := config.Conf.Save(); err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			if started {
				return
			}
			started = true
			audio.Play(gen.TouchButton)
			restart(false)
		}, MainWindow)
		PlayLocalButtonSettingDialog.Resize(fyne.NewSize(300, 360))
		PlayLocalButtonSettingDialog.Show()
	})
	PlayLocalButton.Resize(fyne.NewSize(200, 50))
	PlayLocalButton.Move(fyne.NewPos(206, 495))
	Container.Add(PlayLocalButton)
	MainContainer := container.New(&CenterLayout{
		Min: MenuSize,
	}, Container)
	MainContainer.Resize(MenuSize)
	fadeIn(time.Second, 300*time.Millisecond, MainContainer)
	MainWindow.SetContent(MainContainer)
}
