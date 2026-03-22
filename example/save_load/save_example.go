package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/LuigiVanacore/ludum"
	exampleresources "github.com/LuigiVanacore/ludum/example/resources"
	"github.com/LuigiVanacore/ludum/save"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	screenWidth     = 640
	screenHeight    = 480
	defaultFontSize = 24
	saveFilename    = "save_data.json"
)

// GameState is the struct we want to persist
type GameState struct {
	Level int
	Score int
}

type Game struct {
	state     GameState
	textLabel *ludum.TextNode
	engine    *ludum.Engine
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

func NewGame() *Game {
	engine := ludum.NewEngine()
	gameFont := loadDefaultFont()

	textLabel := ludum.NewTextNode("stateLabel", "", gameFont, color.White)
	textLabel.SetPosition(20, 20)
	engine.World().AddNodeToDefaultLayer(textLabel)

	game := &Game{
		engine:    engine,
		textLabel: textLabel,
	}

	game.loadGame()
	return game
}

func (g *Game) loadGame() {
	// Let's try to load the save file using Generics!
	state, err := save.LoadJSON[GameState](saveFilename)
	if err != nil {
		// If fails (e.g. not exists), use default values
		log.Printf("No existing save found or load error: %v. Starting fresh.", err)
		g.state = GameState{Level: 1, Score: 0}
	} else {
		log.Printf("Save loaded successfully!")
		g.state = state
	}
}

func (g *Game) saveGame() {
	// Persist the state atomically
	err := save.SaveJSON(saveFilename, g.state)
	if err != nil {
		log.Printf("Failed to save: %v", err)
	} else {
		log.Printf("Game saved securely to %s", saveFilename)
	}
}

func (g *Game) Update() error {
	// Input logic
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.state.Score += 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		g.state.Level++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.saveGame()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) { // Reload
		g.loadGame()
	}

	// Update UI
	msg := fmt.Sprintf("=== Save/Load Example ===\n\n"+
		"Level: %d\nScore: %d\n\n"+
		"Controls:\n"+
		"[SPACE] Increase Score\n"+
		"[L] Increase Level\n"+
		"[S] SAVE Game\n"+
		"[R] RELOAD File\n\n"+
		"(Close and reopen the game to see persistence!)", g.state.Level, g.state.Score)
	g.textLabel.SetMessage(msg)

	g.engine.Update()
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
	ebiten.SetWindowTitle("Save/Load Demo")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
