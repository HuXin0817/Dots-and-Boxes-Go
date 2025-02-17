package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/HuXin0817/dots-and-boxes/src/audio"
)

func ShowIntroduceInterface() {
	MainWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if event.Name == fyne.KeyM {
			audio.On = !audio.On
		}
	})
	Container := container.NewWithoutLayout()
	MenuSize := fyne.NewSize(612, 600)
	Container.Resize(MenuSize)
	Icon := canvas.NewImageFromResource(IconResource)
	Icon.Resize(fyne.NewSize(150, 150))
	Icon.Move(fyne.NewPos(100, 225))
	Container.Add(Icon)
	Title := canvas.NewText("Dots and Boxes", TextColor())
	Title.Alignment = fyne.TextAlignCenter
	Title.TextSize = 33
	Title.FontSource = TimesNewRomanBoldResource
	Title.Resize(fyne.NewSize(250, 70))
	Title.Move(fyne.NewPos(280, 265))
	Container.Add(Title)
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{
		title: Title,
	})
	MainContainer := container.New(&CenterLayout{
		Min: MenuSize,
	}, Container)
	MainContainer.Resize(MenuSize)
	c := FadeIn(time.Second, time.Second, MainContainer)
	MainWindow.SetContent(MainContainer)
	go func() {
		time.Sleep(800 * time.Millisecond)
		audio.Play(EnterGameMusic)
		<-c
		<-FadeOut(time.Second, 700*time.Millisecond, MainContainer)
		ShowStartMenu()
	}()
}
