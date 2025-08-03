package main
 

import (
	"bytes"
	"os/exec"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("CLI Runner")

	input := widget.NewEntry()
	output := widget.NewMultiLineEntry()
	output.SetMinRowsVisible(10)

	runBtn := widget.NewButton("Run", func() {
		cmd := exec.Command("cmd", "/C", input.Text)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		cmd.Run()
		output.SetText(out.String())
	})

	content := container.NewVBox(input, runBtn, output)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}