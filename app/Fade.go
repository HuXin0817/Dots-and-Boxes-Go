package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func fade(d, wait time.Duration, Container *fyne.Container, StartA uint8) <-chan time.Time {
	BackgroundColor := *StartMenuBackgroundColor()
	BackgroundColor.A = StartA
	Rectangle := canvas.NewRectangle(&BackgroundColor)
	Rectangle.Resize(MainWindow.Canvas().Size())
	go func() {
		time.Sleep(wait)
		if StartA == 0 {
			fyne.NewAnimation(d, func(f float32) {
				BackgroundColor.A = uint8(255 * f)
				Rectangle.Refresh()
			}).Start()
		} else {
			fyne.NewAnimation(d, func(f float32) {
				BackgroundColor.A = uint8(255 * (1 - f))
				Rectangle.Refresh()
			}).Start()
		}
	}()
	Container.Add(Rectangle)
	return time.After(d + wait)
}

func fadeOut(d, wait time.Duration, Container *fyne.Container) <-chan time.Time {
	return fade(d, wait, Container, 0)
}

func fadeIn(d, wait time.Duration, Container *fyne.Container) <-chan time.Time {
	return fade(d, wait, Container, 255)
}
