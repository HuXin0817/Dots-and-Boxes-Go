package main

import "fyne.io/fyne/v2/app"

var MainWindow = app.NewWithID("io.github.dotsandboxes").NewWindow("Dots and Boxes")

func main() {
	ShowIntroduceInterface()
	MainWindow.ShowAndRun()
}
