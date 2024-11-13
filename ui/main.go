package main

import "github.com/rivo/tview"

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Search Box"), 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Search Result Metadata"), 0, 1, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Search Result List"), 0, 4, false),
				0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Selected Item Details"), 0, 1, false), 0, 1, false)

	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
