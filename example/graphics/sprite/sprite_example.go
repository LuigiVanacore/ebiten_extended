package main

import (
	"fmt" 
	"log"
	"math"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	AircraftID   = "Aircraft_1"
)

type Game struct {
	sprite   *ebiten_extended.SpriteNode
	rotation float64
	engine   *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	engine.Resources().AddImage(AircraftID, resources.Aircraft)

	sprite := ebiten_extended.NewSprite("Aircraft_1", engine.Resources().GetImage(AircraftID), true)
	sprite.SetPosition(0, 0)

	sprite2 := ebiten_extended.NewSprite("Aircraft_2", engine.Resources().GetImage(AircraftID), true)
	sprite2.SetPosition(50, 50)

	sprite.AddChild(sprite2)

	gameLayer := ebiten_extended.NewLayer(2, 2, "GameLayer")
	gameLayer.AddNode(sprite)

	engine.World().AddLayer(gameLayer)
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
