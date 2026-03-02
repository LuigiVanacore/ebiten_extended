package main

import (
	"log"
	"math"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth     = 320
	screenHeight    = 240
	AircraftID      = "Aircraft_1"
	DesertID        = "Desert_1"
	BackgroundLayer = 2
	AircraftLayer   = 3
)

type Game struct {
	sprite   *ebiten_extended.SpriteNode
	rotation float64
	engine   *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	engine.Resources().AddImage("aircraft", resources.Aircraft)
	engine.Resources().AddImage("desert", resources.Desert)

	sprite := ebiten_extended.NewSprite("aircraftSprite1", engine.Resources().GetImage(AircraftID), true)
	sprite.SetPosition(screenWidth/2, screenHeight/2)

	sprite2 := ebiten_extended.NewSprite("aircraftSprite2", engine.Resources().GetImage(AircraftID), true)
	sprite2.SetPosition(100, 100)

	desertSprite := ebiten_extended.NewSprite("desertSprite", engine.Resources().GetImage(DesertID), false)

	sprite.AddChildren(sprite2)

	backgroundLayer := ebiten_extended.NewLayer(BackgroundLayer, 0, "backgroundLayer")
	aircraftLayer := ebiten_extended.NewLayer(AircraftLayer, 0, "aircraftLayer")

	backgroundLayer.AddNode(desertSprite)

	aircraftLayer.AddNode(sprite)

	engine.World().AddLayer(backgroundLayer)
	engine.World().AddLayer(aircraftLayer)
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
