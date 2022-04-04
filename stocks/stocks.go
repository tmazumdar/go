package stocks

import "fyne.io/fyne/v2/widget"

func Load() *widget.Card {
	return widget.NewCard("Stocks", "This card will show stock information",
		widget.NewAccordion(
			widget.NewAccordionItem("AAPL", widget.NewLabel("Price: $100")),
			widget.NewAccordionItem("NVIDIA", widget.NewLabel("Price: $300")),
		),
	)
}
