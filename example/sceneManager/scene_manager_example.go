package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/LuigiVanacore/ludum"
	exampleresources "github.com/LuigiVanacore/ludum/example/resources"
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/ui"
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
	engine       *ludum.Engine
	sceneManager *ludum.SceneManager
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

// MenuScene implements Scene for the menu screen.
type MenuScene struct {
	engine *ludum.Engine
	panel  *ui.PanelNode
}

func NewMenuScene(engine *ludum.Engine) *MenuScene {
	return &MenuScene{engine: engine}
}

func (s *MenuScene) Enter(engine *ludum.Engine) {
	s.engine = engine
	font := loadDefaultFont()
	s.panel = ui.NewPanelNode("menu_panel", 400, 200)
	s.panel.SetBackgroundColor(color.RGBA{40, 40, 60, 255})
	s.panel.SetPosition(screenWidth/2-200, screenHeight/2-100)

	title := ludum.NewTextNode("title", "Scene: Menu", font, color.White)
	title.SetPosition(20, 30)
	s.panel.AddChildren(title)

	hint := ludum.NewTextNode("hint", "Press Enter to start, Escape to return", font, color.RGBA{180, 180, 180, 255})
	hint.SetPosition(20, 80)
	s.panel.AddChildren(hint)

	engine.World().AddNodeToDefaultLayer(s.panel)
}

func (s *MenuScene) Exit() {
	if s.engine != nil && s.panel != nil {
		s.engine.World().RemoveNode(s.panel)
	}
}

func (s *MenuScene) Update() error {
	return nil
}

func (s *MenuScene) Draw(screen *ebiten.Image) {
	s.engine.World().Draw(screen, nil)
}

// GameScene implements Scene for the game screen.
type GameScene struct {
	engine *ludum.Engine
	root   *ludum.Node
}

func NewGameScene(engine *ludum.Engine) *GameScene {
	return &GameScene{engine: engine}
}

func (s *GameScene) Enter(engine *ludum.Engine) {
	s.engine = engine
	s.root = ludum.NewNode("game_root")

	circle := ludum.NewDrawnCircle("circle", math2d.NewVector2D(200, 150), 40, color.RGBA{0, 200, 100, 255}, true, 0)
	rect := ludum.NewDrawnRectangle("rect", math2d.NewVector2D(400, 250), math2d.NewVector2D(80, 60), color.RGBA{200, 50, 50, 255}, true, 0)

	s.root.AddChildren(circle)
	s.root.AddChildren(rect)

	engine.World().AddNodeToDefaultLayer(s.root)
}

func (s *GameScene) Exit() {
	if s.engine != nil && s.root != nil {
		s.engine.World().RemoveNode(s.root)
	}
}

func (s *GameScene) Update() error {
	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.engine.World().Draw(screen, nil)
}

func NewGame() *Game {
	engine := ludum.NewEngine()
	engine.Input().SetMouseEnabled(true)

	sm := ludum.NewSceneManager(engine)
	sm.SetTransitionDuration(0.3) // enable fade transitions between scenes
	engine.SetSceneManager(sm)

	// Start with menu
	sm.PushScene(NewMenuScene(engine))

	return &Game{
		engine:       engine,
		sceneManager: sm,
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter) {
		if _, ok := g.sceneManager.CurrentScene().(*MenuScene); ok {
			g.sceneManager.ReplaceScene(NewGameScene(g.engine))
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if _, ok := g.sceneManager.CurrentScene().(*GameScene); ok {
			g.sceneManager.ReplaceScene(NewMenuScene(g.engine))
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
