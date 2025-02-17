package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	widgetx "fyne.io/x/fyne/widget"
	"github.com/HuXin0817/dots-and-boxes/src/ai"
	"github.com/HuXin0817/dots-and-boxes/src/audio"
)

func ShowStartMenu() {
	MainWindow.SetTitle("Dots and Boxes")
	Container := container.NewWithoutLayout()
	MenuSize := fyne.NewSize(612, 600)
	Container.Resize(MenuSize)
	Icon := canvas.NewImageFromResource(IconResource)
	Icon.Resize(fyne.NewSize(200, 200))
	Icon.Move(fyne.NewPos(206, 50))
	Container.Add(Icon)
	Title := canvas.NewText("Dots and Boxes", TextColor())
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 26
	Title.FontSource = TimesNewRomanBoldResource
	Title.Resize(fyne.NewSize(200, 50))
	Title.Move(fyne.NewPos(206, 250))
	Container.Add(Title)
	Link := canvas.NewText("https://github.com/HuXin0817/Dots-and-Boxes", LinkColor)
	Link.Resize(fyne.NewSize(200, 50))
	Link.Alignment = fyne.TextAlignCenter
	Link.TextSize = 17
	Link.TextStyle = fyne.TextStyle{Bold: true, Italic: true}
	Link.FontSource = TimesNewRomanItalicResource
	Link.Move(fyne.NewPos(206, 285))
	Container.Add(Link)
	Spinner, err := widgetx.NewAnimatedGifFromResource(SpinnerGIFResource())
	if err != nil {
		dialog.NewError(err, MainWindow).Show()
		return
	}
	Spinner.Start()
	Spinner.Resize(fyne.NewSize(70, 70))
	Spinner.Move(fyne.NewPos(271, 335))
	Container.Add(Spinner)
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{
		gif:   Spinner,
		title: Title,
	})
	started := false
	PlayOnlineButton := widget.NewButton("Play Online", func() {
		if started {
			return
		}
		started = true
		audio.Play(TouchButtonMusic)
		restart(true)
	})
	PlayOnlineButton.Resize(fyne.NewSize(200, 50))
	PlayOnlineButton.Move(fyne.NewPos(206, 435))
	Container.Add(PlayOnlineButton)
	PlayLocalButton := widget.NewButton("Play Local", func() {
		TmpConfig := Conf
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
		if Conf.AI1 {
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
		if Conf.AI2 {
			Player2RadioGroup.SetSelected("Computer")
		} else {
			Player2RadioGroup.SetSelected("People")
		}
		Entry1 := widget.NewEntry()
		Entry1.SetText(Conf.AI1Name)
		Entry2 := widget.NewEntry()
		Entry2.SetText(Conf.AI2Name)
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
			Conf = TmpConfig
			if err := Conf.Save(); err != nil {
				dialog.NewError(err, MainWindow).Show()
			}
			if started {
				return
			}
			started = true
			audio.Play(TouchButtonMusic)
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
	FadeIn(time.Second, 300*time.Millisecond, MainContainer)
	MainWindow.SetContent(MainContainer)
	MainWindow.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyM {
			audio.On = !audio.On
		}
	})
}
