package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Hello")

	button := widget.NewButton("Click Me", func() {
		println("Button Clicked")
	})

	window.SetContent(button)
	window.ShowAndRun()
}
