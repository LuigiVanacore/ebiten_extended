package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	exampleresources "github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/LuigiVanacore/ebiten_extended/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	defaultFontSize = 24
)

type Game struct {
	engine *ebiten_extended.Engine
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
	engine := ebiten_extended.NewEngine()
	im := engine.Input()
	im.SetMouseEnabled(true) // Required for Button logic

	gameFont := loadDefaultFont()

	// 1. Create a background Panel
	panel := ui.NewPanelNode("background_panel", 300, 400)
	panel.SetBackgroundColor(color.RGBA{R: 50, G: 50, B: 80, A: 255})
	panel.SetPosition(screenWidth/2-150, 40) // Center horizontally

	// 2. Create a Text label at the top of the panel
	titleLabel := ebiten_extended.NewTextNode("title_lbl", "UI Example", gameFont, color.White)
	titleLabel.SetPosition(40, 20)
	panel.AddChildren(titleLabel)

	// 3. Create an interactive Button
	clickCount := 0
	counterLabel := ebiten_extended.NewTextNode("counter_lbl", "Clicks: 0", gameFont, color.RGBA{255, 200, 0, 255})
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

	// 5. Create a Checkbox
	checkbox := ui.NewCheckboxNode("theme_checkbox", 30, im)
	checkbox.SetPosition(50, 300)
	panel.AddChildren(checkbox)

	chkLabel := ebiten_extended.NewTextNode("chk_lbl", "Toggle Red Progress", gameFont, color.White)
	chkLabel.SetPosition(90, 300)
	panel.AddChildren(chkLabel)

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
