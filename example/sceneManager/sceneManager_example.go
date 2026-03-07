package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/LuigiVanacore/ebiten_extended"
	exampleresources "github.com/LuigiVanacore/ebiten_extended/example/resources"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	defaultFontSize = 24
)

type Game struct {
	engine   *ebiten_extended.Engine
	menuRoot *ui.PanelNode
	gameRoot *ebiten_extended.Node

	currentScene string // "menu" or "game"
}

func loadDefaultFont() text.Face {
	tt, err := text.NewGoTextFaceSource(bytes.NewReader(exampleresources.DefaultFont))
	if err != nil {
		log.Fatal(err)
	}
	return &text.GoTextFace{
		Source: tt,
		Size:   defaultFontSize,
	}
}

func buildMenuScene(engine *ebiten_extended.Engine) *ui.PanelNode {
	font := loadDefaultFont()
	panel := ui.NewPanelNode("menu_panel", 400, 200)
	panel.SetBackgroundColor(color.RGBA{40, 40, 60, 255})
	panel.SetPosition(screenWidth/2-200, screenHeight/2-100)

	title := ebiten_extended.NewTextNode("title", "Scene: Menu", font, color.White)
	title.SetPosition(20, 30)
	panel.AddChildren(title)

	hint := ebiten_extended.NewTextNode("hint", "Press Enter to start, Escape to return", font, color.RGBA{180, 180, 180, 255})
	hint.SetPosition(20, 80)
	panel.AddChildren(hint)

	return panel
}

func buildGameScene() *ebiten_extended.Node {
	root := ebiten_extended.NewNode("game_root")

	circle := ebiten_extended.NewDrawnCircle("circle", math2D.NewVector2D(200, 150), 40, color.RGBA{0, 200, 100, 255}, true, 0)
	rect := ebiten_extended.NewDrawnRectangle("rect", math2D.NewVector2D(400, 250), math2D.NewVector2D(80, 60), color.RGBA{200, 50, 50, 255}, true, 0)

	root.AddChildren(circle)
	root.AddChildren(rect)

	return root
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
	engine.Input().SetMouseEnabled(true)

	menuRoot := buildMenuScene(engine)
	gameRoot := buildGameScene()

	// Start with menu
	engine.World().AddNodeToDefaultLayer(menuRoot)

	return &Game{
		engine:       engine,
		menuRoot:     menuRoot,
		gameRoot:     gameRoot,
		currentScene: "menu",
	}
}

func (g *Game) switchToScene(name string) {
	g.engine.World().ClearLayer(0)
	switch name {
	case "menu":
		g.engine.World().AddNodeToDefaultLayer(g.menuRoot)
		g.currentScene = "menu"
	case "game":
		g.engine.World().AddNodeToDefaultLayer(g.gameRoot)
		g.currentScene = "game"
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
		if g.currentScene == "menu" {
			g.switchToScene("game")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.currentScene == "game" {
			g.switchToScene("menu")
		}
	}

	return g.engine.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.engine.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Scene Manager - Enter: start, Escape: menu")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
