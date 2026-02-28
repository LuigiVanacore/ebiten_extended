package main

import (
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	AircraftID   = "Aircraft_1"
)

type Game struct {
	sprite   *ebiten_extended.SpriteNode
	rotation int
	engine   *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	engine.Resources().AddImage(AircraftID, resources.Aircraft)

	sprite := ebiten_extended.NewSprite("Aircraft_1", engine.Resources().GetImage(AircraftID), true)
	sprite.SetPosition(0, 0)

	sprite2 := ebiten_extended.NewSprite("Aircraft_2", engine.Resources().GetImage(AircraftID), true)
	sprite2.SetPosition(50, 50)

	sprite.AddChildren(sprite2)

	gameLayer := ebiten_extended.NewLayer(2, 2, "GameLayer")
	gameLayer.AddNode(sprite)

	engine.World().AddLayer(gameLayer)
	engine.SetIsDebug(true)
	return &Game{sprite: sprite, engine: engine}
}

func (g *Game) Update() error {
	g.engine.Update()
	g.rotation = g.rotation + 1%360
	g.sprite.SetRotation(g.rotation)
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
