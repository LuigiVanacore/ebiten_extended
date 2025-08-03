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
	rotation int
}

func NewGame() *Game {
	circle := ebiten_extended.NewDrawnCircle("Circle", math2D.NewVector2D(100, 100), 50, RED_COLOR, true, 0)

	rectangle := ebiten_extended.NewDrawnRectangle("Rectangle", math2D.NewVector2D(200, 200), math2D.NewVector2D(100, 50), RED_COLOR, true, 0)

	ebiten_extended.GameManager().World().AddNode(circle)
	ebiten_extended.GameManager().World().AddNode(rectangle)

	ebiten_extended.GameManager().SetIsDebug(false)
	return &Game{}
}

func (g *Game) Update() error {
	ebiten_extended.GameManager().Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebiten_extended.GameManager().Draw(screen, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite Example")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
