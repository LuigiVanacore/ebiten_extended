package main

import (
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/resources"
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

	sprite := ebiten_extended.NewSprite("Aircraft_1", ebiten_extended.ResourceManager().GetImage(AircraftID), true)
	sprite.SetPosition(math2D.NewVector2D(0, 0))

	sprite2 := ebiten_extended.NewSprite("Aircraft_2", ebiten_extended.ResourceManager().GetImage(AircraftID), true)
	sprite2.SetPosition(math2D.NewVector2D(50, 100))

	sprite.AddChild(sprite2)

	gameLayer := ebiten_extended.NewLayer(2, 2, "GameLayer")
	gameLayer.AddNode(sprite)

	//ebiten_extended.SceneManager().AddEntityToDefaultLayer(sprite, "SpriteNode")
	//ebiten_extended.SceneManager().AddSceneNodeToDefaultLayer(sprite)
	ebiten_extended.GameManager().World().AddLayer(gameLayer)
	ebiten_extended.GameManager().SetIsDebug(false)
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
