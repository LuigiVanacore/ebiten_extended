package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/LuigiVanacore/ludum"
	exampleresources "github.com/LuigiVanacore/ludum/example/resources"
	"github.com/LuigiVanacore/ludum/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	defaultFontSize = 24
)

type Game struct {
	engine *ludum.Engine
}

func loadDefaultFont() text.Face {
	tt, err := text.NewGoTextFaceSource(bytes.NewReader(exampleresources.DefaultFont))
	if err != nil {
		log.Fatal(err)
	}
	return &text.GoTextFace{
		Source: tt,
		Size:   defaultFontSize,
	}
}

func NewGame() *Game {
	engine := ludum.NewEngine()
	im := engine.Input()
	im.SetMouseEnabled(true) // Required for Button logic

	gameFont := loadDefaultFont()

	// 1. Create a background Panel
	panel := ui.NewPanelNode("background_panel", 300, 400)
	panel.SetBackgroundColor(color.RGBA{R: 50, G: 50, B: 80, A: 255})
	panel.SetPosition(screenWidth/2-150, 40) // Center horizontally

	// 2. Create a Text label at the top of the panel
	titleLabel := ludum.NewTextNode("title_lbl", "UI Example", gameFont, color.White)
	titleLabel.SetPosition(40, 20)
	panel.AddChildren(titleLabel)

	// 3. Create an interactive Button
	clickCount := 0
	counterLabel := ludum.NewTextNode("counter_lbl", "Clicks: 0", gameFont, color.RGBA{255, 200, 0, 255})
	counterLabel.SetPosition(40, 60)
	panel.AddChildren(counterLabel)

	button := ui.NewButtonNode("my_button", 200, 50, im)
	button.SetPosition(50, 200) // Relative to panel parent
	button.SetText("Click Me!", gameFont, color.White)

	// 4. Create a ProgressBar
	progressBar := ui.NewProgressBarNode("progress", 200, 20)
	progressBar.SetPosition(50, 260)
	panel.AddChildren(progressBar)

	// Bind callbacks to the button
	button.OnClick = func() {
		clickCount++
		counterLabel.SetMessage(fmt.Sprintf("Clicks: %d", clickCount))

		// Fill progress bar (loops around)
		prog := progressBar.GetProgress()
		prog += 0.1
		if prog > 1 {
			prog = 0
		}
		progressBar.SetProgress(prog)
	}
	button.OnMouseEnter = func() {
		titleLabel.SetMessage("Hovering!")
	}
	button.OnMouseExit = func() {
		titleLabel.SetMessage("UI Example")
	}

	panel.AddChildren(button)

	tooltip := ui.NewTooltipNode("btn_tooltip", button, "Increments the counter and progress bar", gameFont, im)
	tooltip.SetPosition(0, 0)
	panel.AddChildren(tooltip)

	// 5. Create a Checkbox
	checkbox := ui.NewCheckboxNode("theme_checkbox", 30, im)
	checkbox.SetPosition(50, 300)
	panel.AddChildren(checkbox)

	chkLabel := ludum.NewTextNode("chk_lbl", "Toggle Red Progress", gameFont, color.White)
	chkLabel.SetPosition(90, 300)
	panel.AddChildren(chkLabel)

	// 6. Text input field
	textInput := ui.NewTextInputNode("name_input", 200, 30, gameFont, im)
	textInput.SetPosition(50, 340)
	textInput.SetPlaceholder("Enter name...")
	textInput.SetMaxLength(20)
	textInput.OnSubmit = func(text string) {
		titleLabel.SetMessage("Hello, " + text + "!")
	}
	panel.AddChildren(textInput)

	// 7. Dropdown
	dropdown := ui.NewDropdownNode("theme_dropdown", 200, 30, gameFont, im)
	dropdown.SetPosition(50, 380)
	dropdown.SetItems([]string{"Option A", "Option B", "Option C"})
	dropdown.OnSelectionChanged = func(idx int, text string) {
		titleLabel.SetMessage("Selected: " + text)
	}
	panel.AddChildren(dropdown)

	checkbox.OnToggle = func(checked bool) {
		if checked {
			progressBar.SetFillColor(color.RGBA{200, 0, 0, 255})
		} else {
			progressBar.SetFillColor(color.RGBA{0, 200, 0, 255})
		}
	}

	// Attach the root panel to the world
	engine.World().AddNodeToDefaultLayer(panel)

	return &Game{engine: engine}
}

func (g *Game) Update() error {
	g.engine.Update() // Evaluates UI interactables automatically via World
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Interactive UI Panel Demo")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
