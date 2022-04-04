package clocks

import (
	"time"

	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func Load() *widget.Label {

	clock := widget.NewLabel("")

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	return clock
}
