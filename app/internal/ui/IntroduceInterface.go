package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/HuXin0817/dots-and-boxes/app/internal/audio"
	"github.com/HuXin0817/dots-and-boxes/app/internal/gen"
)

var IntroduceInterfaceIconCanvas = func() *canvas.Image {
	IconCanvas := canvas.NewImageFromResource(Icon)
	IconCanvas.Resize(fyne.NewSize(150, 150))
	return IconCanvas
}()

var IntroduceInterfaceTitleText = func() *canvas.Text {
	TitleText := canvas.NewText("Dots and Boxes", TextColor())
	TitleText.Alignment = fyne.TextAlignCenter
	TitleText.TextSize = 33
	TitleText.FontSource = TimesNewRomanBold
	TitleText.Resize(fyne.NewSize(250, 70))
	return TitleText
}()

func ShowIntroduceInterface() {
	fyne.CurrentApp().Settings().SetTheme(&StartMenuTheme{})
	Container := container.NewWithoutLayout()
	Container.Resize(StartMenuSize)
	IntroduceInterfaceIconCanvas.Move(fyne.NewPos(100, 225))
	Container.Add(IntroduceInterfaceIconCanvas)
	IntroduceInterfaceTitleText.Move(fyne.NewPos(280, 265))
	Container.Add(IntroduceInterfaceTitleText)
	MainContainer := container.New(&CenterLayout{
		Min: StartMenuSize,
	}, Container)
	MainContainer.Resize(StartMenuSize)
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
