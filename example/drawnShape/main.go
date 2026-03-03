package main

import (
	"image/color"
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var RED_COLOR = color.RGBA{0xf0, 0x31, 0x31, 0xff}

type Game struct {
	engine *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()

	circle := ebiten_extended.NewDrawnCircle("Circle", math2D.NewVector2D(100, 100), 50, RED_COLOR, true, 0)
	rectangle := ebiten_extended.NewDrawnRectangle("Rectangle", math2D.NewVector2D(200, 200), math2D.NewVector2D(100, 50), RED_COLOR, true, 0)

	engine.World().AddNodeToLayer(circle, 0)
	engine.World().AddNodeToLayer(rectangle, 0)
	engine.SetIsDebug(false)

	return &Game{engine: engine}
}

func (g *Game) Update() error {
	return g.engine.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("DrawnShape Example")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
