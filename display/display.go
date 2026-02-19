package display

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func StartChatUI() error {
	app := tview.NewApplication()

	inputField := tview.NewInputField().
		SetLabel("Message: ").
		SetFieldWidth(50).
		SetPlaceholder("John Smith").
		SetAcceptanceFunc(tview.InputFieldMaxLength(150)).
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})
	if err := app.SetRoot(inputField, true).SetFocus(inputField).Run(); err != nil {
		panic(err)
	}

	return nil
}
