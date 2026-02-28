package main

import (
	"bytes"
	"image/color"
	"log"
	"strconv"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	defaultFontSize = 24
	defaultFontDPI  = 72
)

type Game struct {
	counter   int
	textLabel *ebiten_extended.TextNode
	engine    *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	gameFont := loadDefaultFont()
	textLabel := ebiten_extended.NewTextNode("labelTest", "test label text", gameFont, color.White)
	textLabel.SetPosition(0, 0)
	engine.World().AddNodeToDefaultLayer(textLabel)
	return &Game{textLabel: textLabel, engine: engine}
}

func (g *Game) Update() error {
	if g.counter%ebiten.TPS() == 0 {
		g.counter = 0
	}
	message := "test label text " + strconv.Itoa(g.counter)
	g.textLabel.SetMessage(message)
	g.counter++
	g.engine.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadDefaultFont() text.Face {

	tt, err := text.NewGoTextFaceSource(bytes.NewReader(resources.DefaultFont))
	if err != nil {
		log.Fatal(err)
	}

	gamefont := &text.GoTextFace{
		Source: tt,
		Size:   defaultFontSize,
	}
	return gamefont
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
