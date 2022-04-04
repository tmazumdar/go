package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	app := app.New()
	window := app.NewWindow("Orphans O|o")

	notes := widget.NewEntry()
	stocks := widget.NewCard("Stocks", "This card will show stock information", widget.NewAccordion(
		widget.NewAccordionItem("AAPL", widget.NewLabel("Price: $100")),
		widget.NewAccordionItem("NVIDIA", widget.NewLabel("Price: $300")),
	))
	clock := widget.NewLabel("")

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	tabs := container.NewAppTabs(
		container.NewTabItem("Notes", notes),
		container.NewTabItem("Stocks", stocks),
		container.NewTabItem("Clock", clock),
	)

	window.SetContent(tabs)
	window.Resize(fyne.NewSize(1280, 720))

	window.Show()
	app.Run()
}
