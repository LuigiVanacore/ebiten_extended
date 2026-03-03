package main

import (
	"log"
	"math"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	aircraftID      = "aircraft"
	desertID        = "desert"
	BackgroundLayer = 0
	AircraftLayer   = 1
)

type Game struct {
	sprite   *ebiten_extended.Sprite
	rotation float64
	engine   *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	if err := engine.Resources().AddImage("aircraft", resources.Aircraft); err != nil {
		log.Fatal(err)
	}
	if err := engine.Resources().AddImage("desert", resources.Desert); err != nil {
		log.Fatal(err)
	}

	sprite := ebiten_extended.NewSprite("aircraftSprite1", engine.Resources().GetImage(aircraftID), 0, true)
	sprite.SetPosition(screenWidth/2, screenHeight/2)

	sprite2 := ebiten_extended.NewSprite("aircraftSprite2", engine.Resources().GetImage(aircraftID), 0, true)
	sprite2.SetPosition(100, 100)

	desertSprite := ebiten_extended.NewSprite("desertSprite", engine.Resources().GetImage(desertID), 0, false)

	sprite.AddChildren(sprite2)

	engine.World().AddNodeToLayer(desertSprite, BackgroundLayer)
	engine.World().AddNodeToLayer(sprite, AircraftLayer)
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
	ebiten.SetWindowTitle("Layout Example")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
