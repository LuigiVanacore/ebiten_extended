package main

import (
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	AircraftID      = "Aircraft"
	DesertID        = "Desert"
	BackgroundLayer = 2
	AircraftLayer   = 3
)

type Game struct {
	sprite   *ebiten_extended.Sprite
	rotation int
}

func NewGame() *Game {
	ebiten_extended.ResourceManager().AddImage(AircraftID, resources.Aircraft)
	ebiten_extended.ResourceManager().AddImage(DesertID, resources.Desert)

	sprite := ebiten_extended.NewSprite("aircraftSprite1", ebiten_extended.ResourceManager().GetImage(AircraftID), AircraftLayer, true)
	sprite.SetPosition(math2D.NewVector2D(screenWidth/2, screenHeight/2))

	sprite2 := ebiten_extended.NewSprite("aircraftSprite2", ebiten_extended.ResourceManager().GetImage(AircraftID), AircraftLayer, true)
	sprite2.SetPosition(math2D.NewVector2D(100, 100))

	desertSprite := ebiten_extended.NewSprite("desertSprite", ebiten_extended.ResourceManager().GetImage(DesertID), BackgroundLayer, false)

	sprite.AddChild(sprite2)
 

 
	ebiten_extended.GameManager().World().AddNode(sprite)
	ebiten_extended.GameManager().World().AddNode(desertSprite)
	ebiten_extended.GameManager().SetIsDebug(true)
	return &Game{sprite: sprite}
}

func (g *Game) Update() error {
	ebiten_extended.GameManager().Update()
	g.rotation = g.rotation + 1%360
	g.sprite.SetRotation(g.rotation)
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
