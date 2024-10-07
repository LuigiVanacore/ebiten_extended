package ebiten_extended

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
)


const ( 
	fixedDelta = 1.0 / 60.0
)

var gameManager_instance *gameManager


func GameManager() *gameManager {
	if gameManager_instance == nil {
		gameManager_instance = newGameManager()
	}

	return gameManager_instance
}

type gameManager struct {
	isRunning bool
	isDebug bool
	world *World
}

func newGameManager() *gameManager {
	return &gameManager{isRunning:  true, isDebug: false, world: NewWorld()}
}

func init() {
	InitTextManager()
}

func (g *gameManager) Pause() {
	g.isRunning = false
}

func (g *gameManager) Start() {
	g.isRunning = true
}

func (g *gameManager) IsDebug() bool {
	return g.isDebug
}

func (g *gameManager) World() *World {
	return g.world
}

func (g *gameManager) SetIsDebug(debugFlag bool ){
	g.isDebug = debugFlag
	fmt.Printf("The debug flag is setted as %t", g.isDebug) 
}

func (g *gameManager) Update() {
	
	if g.isRunning {
		input.InputManager().Update()
		g.world.Update()
		
	}

}



func (g *gameManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	g.world.Draw(target, op)
}