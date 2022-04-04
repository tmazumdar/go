package main

import (
	"orphans/clocks"
	"orphans/notes"
	"orphans/stocks"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()
	window := app.NewWindow("Orphans O|o")

	tabs := container.NewAppTabs(
		container.NewTabItem("Notes", notes.Load()),
		container.NewTabItem("Stocks", stocks.Load()),
		container.NewTabItem("Clock", clocks.Load()),
	)

	window.SetContent(tabs)
	window.Resize(fyne.NewSize(1280, 720))

	window.Show()
	app.Run()
}
