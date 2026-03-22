package main

import (
	"log"
	"math"

	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/example/resources"
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
	sprite   *ludum.Sprite
	rotation float64
	engine   *ludum.Engine
}

func NewGame() *Game {
	engine := ludum.NewEngine()
	if err := engine.Resources().AddImage("aircraft", resources.Aircraft); err != nil {
		log.Fatal(err)
	}
	if err := engine.Resources().AddImage("desert", resources.Desert); err != nil {
		log.Fatal(err)
	}

	sprite := ludum.NewSprite("aircraftSprite1", engine.Resources().GetImage(aircraftID), 0, true)
	sprite.SetPosition(screenWidth/2, screenHeight/2)

	sprite2 := ludum.NewSprite("aircraftSprite2", engine.Resources().GetImage(aircraftID), 0, true)
	sprite2.SetPosition(100, 100)

	desertSprite := ludum.NewSprite("desertSprite", engine.Resources().GetImage(desertID), 0, false)

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
