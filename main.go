package main

import (
	"fmt"
	"go-chat/client"
	"go-chat/server"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const IP = "127.0.0.1"
const PORT = 1080

func main() {

	var i int
	fmt.Print("Start go-chat Server (0) or Client (1) 0/1: ")
	fmt.Scan(&i)

	if i == 0 {
		s, err := server.New(8080)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}

		go s.Start()
	} else {
		go client.Start(IP, PORT)
	}

	app := tview.NewApplication()

	newPrimitive := func(text string, alignement int) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(alignement).
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

	sideBar := newPrimitive(fmt.Sprintf("Current connection:\n\nIP: %s\nPORT: %d", IP, PORT), tview.AlignLeft)
	messagesContainer := newPrimitive("Messages will show up here.", tview.AlignCenter)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 100).
		SetBorders(true).
		AddItem(newPrimitive("Welcome to T-CHAT ! A TCP Chat app in the terminal !", tview.AlignCenter), 0, 0, 1, 3, 0, 0, false).
		AddItem(msgIput, 2, 0, 1, 3, 0, 0, false)

	grid.AddItem(sideBar, 1, 0, 1, 1, 1, 0, false).
		AddItem(messagesContainer, 1, 0, 1, 3, 0, 0, false)

	if err := app.SetRoot(grid, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}
}
