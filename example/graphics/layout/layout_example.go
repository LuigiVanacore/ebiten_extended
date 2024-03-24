package main 

import (
	"log"
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 320
	screenHeight = 240
	AircraftID = 0
	DesertID = 1
	BackgroundLayer = 2
	AircraftLayer = 3
)

type Game struct {
	sprite *ebiten_extended.Sprite
	rotation int
}

func NewGame() *Game {
	ebiten_extended.ResourceManager().LoadImage(resources.Aircraft)
	ebiten_extended.ResourceManager().LoadImage(resources.Desert)

	sprite := ebiten_extended.NewSprite(ebiten_extended.ResourceManager().GetTexture(AircraftID), true)
	sprite.SetPosition(screenWidth/2, screenHeight/2)

	sprite2:= ebiten_extended.NewSprite(ebiten_extended.ResourceManager().GetTexture(AircraftID), true)
	sprite2.SetPosition(100,100)

	desertSprite := ebiten_extended.NewSprite(ebiten_extended.ResourceManager().GetTexture(DesertID),false)
	

	sprite.AddChildren(sprite2)


	backgroundLayer := ebiten_extended.NewLayer(BackgroundLayer, 0)
	aircraftLayer := ebiten_extended.NewLayer(AircraftLayer, 0)

	backgroundLayer.AddNode(desertSprite)

	aircraftLayer.AddNode(sprite)

	//ebiten_extended.SceneManager().AddEntityToDefaultLayer(sprite, "SpriteNode")
	ebiten_extended.SceneManager().AddLayer(backgroundLayer)
	ebiten_extended.SceneManager().AddLayer(aircraftLayer)
	ebiten_extended.GameManager().SetIsDebug(true)
	return &Game{ sprite: sprite}
}

func (g *Game) Update() error {
	ebiten_extended.GameManager().Update()
	g.rotation= g.rotation+1%360
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