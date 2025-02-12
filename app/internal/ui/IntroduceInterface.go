package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/HuXin0817/dots-and-boxes/app/internal/audio"
	"github.com/HuXin0817/dots-and-boxes/app/internal/gen"
)

func ShowIntroduceInterface() {
	Container := container.NewWithoutLayout()
	MenuSize := fyne.NewSize(612, 600)
	Container.Resize(MenuSize)
	IconCanvas := canvas.NewImageFromResource(Icon)
	IconCanvas.Resize(fyne.NewSize(150, 150))
	IconCanvas.Move(fyne.NewPos(100, 225))
	Container.Add(IconCanvas)
	TitleText := canvas.NewText("Dots and Boxes", TextColor())
	TitleText.Alignment = fyne.TextAlignCenter
	TitleText.TextSize = 33
	TitleText.FontSource = TimesNewRomanBold
	TitleText.Resize(fyne.NewSize(250, 70))
	TitleText.Move(fyne.NewPos(280, 265))
	Container.Add(TitleText)
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{
		title: TitleText,
	})
	MainContainer := container.New(&CenterLayout{
		Min: MenuSize,
	}, Container)
	MainContainer.Resize(MenuSize)
	c := fadeIn(time.Second, time.Second, MainContainer)
	MainWindow.SetContent(MainContainer)
	go func() {
		time.Sleep(800 * time.Millisecond)
		audio.Play(gen.EnterGame)
		<-c
		<-fadeOut(time.Second, 700*time.Millisecond, MainContainer)
		ShowStartMenu()
	}()
}
