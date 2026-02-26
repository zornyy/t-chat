package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {

	// var i int

	// fmt.Print("Start go-chat Server (0) or Client (1) 0/1: ")
	// fmt.Scan(&i)

	// if i == 0 {
	// 	server.Start()
	// } else {
	// 	client.Start()
	// }

	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	inputField := tview.NewInputField().
		SetLabel("-> ").
		SetFieldWidth(0).
		SetFieldTextColor(tcell.ColorBlack).
		SetFieldBackgroundColor(tcell.ColorWhite).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

	msgIput := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(inputField, 0, 1, true)

	sideBar := newPrimitive("Current connection")
	messagesContainer := newPrimitive("Messages will show up here.")

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 100).
		SetBorders(true).
		AddItem(newPrimitive("Welcome to T-CHAT ! A TCP Chat app running in the terminal !"), 0, 0, 1, 3, 0, 0, false).
		AddItem(msgIput, 2, 0, 1, 3, 0, 0, false)

	// Layout for screens narrower than 100 cells (menu and side bar are hidden).
	grid.AddItem(sideBar, 1, 0, 1, 1, 1, 0, false).
		AddItem(messagesContainer, 1, 0, 1, 3, 0, 0, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
}
