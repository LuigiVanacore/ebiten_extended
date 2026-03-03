package main

import (
	"fmt"
	"log"
	"math"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	aircraftID   = "aircraft"
)

type Game struct {
	sprite   *ebiten_extended.Sprite
	rotation float64
	engine   *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	if err := engine.Resources().AddImage(aircraftID, resources.Aircraft); err != nil {
		log.Fatal(err)
	}

	sprite := ebiten_extended.NewSprite("Aircraft_1", engine.Resources().GetImage(aircraftID), 0, true)
	sprite.SetPosition(screenWidth/2, screenHeight/2)

	sprite2 := ebiten_extended.NewSprite("Aircraft_2", engine.Resources().GetImage(aircraftID), 0, true)
	sprite2.SetPosition(50, 50)

	sprite.AddChildren(sprite2)

	engine.World().AddNodeToLayer(sprite, 0)
	engine.SetIsDebug(true)
	return &Game{sprite: sprite, engine: engine}
}

func (g *Game) Update() error {
	g.engine.Update()
	g.rotation += 0.05
	if g.rotation >= 2*math.Pi {
		g.rotation -= 2 * math.Pi
	}
	g.sprite.SetRotation(g.rotation)
	fmt.Println("Sprite rotation:", g.rotation)
	transform := g.sprite.GetTransform()
	position := transform.GetPosition()
	fmt.Println("Sprite position:", position.String())
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
	ebiten.SetWindowTitle("Sprite Example")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
