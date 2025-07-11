package main

import (
	"fmt"
	"log"

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
	sprite   *ebiten_extended.Sprite
	rotation int
}

func NewGame() *Game {
	ebiten_extended.ResourceManager().AddImage(AircraftID, resources.Aircraft)

	sprite := ebiten_extended.NewSprite("Aircraft_1", ebiten_extended.ResourceManager().GetImage(AircraftID), 0, true)
	sprite.SetPosition(math2D.NewVector2D(50, 50))

	sprite2 := ebiten_extended.NewSprite("Aircraft_2", ebiten_extended.ResourceManager().GetImage(AircraftID), 0, true)
	sprite2.SetPosition(math2D.NewVector2D(100, 100))

	sprite.AddChild(sprite2)

 
	ebiten_extended.GameManager().World().AddNode(sprite)
	ebiten_extended.GameManager().SetIsDebug(false)
	return &Game{sprite: sprite}
}

func (g *Game) Update() error {
	ebiten_extended.GameManager().Update()
	g.rotation = g.rotation + 1%360
	g.sprite.SetRotation(g.rotation)
	fmt.Println("Sprite rotation:", g.rotation)
	transform := g.sprite.GetTransform()
	position := transform.GetPosition()
	fmt.Println("Sprite position:", position.String())
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
