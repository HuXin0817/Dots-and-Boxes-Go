package main

import "github.com/HuXin0817/dots-and-boxes/app/internal/ui"

//go:generate go run cmd/generate.go
func main() {
	ui.ShowIntroduceInterface()
	ui.MainWindow.ShowAndRun()
}
