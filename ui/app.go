package ui

import (
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application
}

func NewApp() *App {
	app := &App{}
	app.Application = tview.NewApplication()
	layout := tview.NewFlex()

	mainPanel := NewMainPanel()
	sidePanel := tview.NewBox().SetBorder(true).SetTitle("side panel")
	//bottomPanel := tview.NewBox().SetBorder(true).SetTitle("Git commands")
	input := tview.NewInputField().SetFieldBackgroundColor(tcell.ColorBlack)
	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			cmd := exec.Command("git", strings.Split(input.GetText(), " ")...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Run()
			input.SetText("")
			mainPanel.Refresh()
		}
	})
	input.SetBorder(true).SetTitle("Git commands")

	layout.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).AddItem(mainPanel, 0, 5, false).AddItem(input, 0, 1, true), 0, 5, true)
	layout.AddItem(sidePanel, 0, 1, false)

	app.SetRoot(layout, true)

	return app
}
