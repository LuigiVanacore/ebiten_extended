package ebiten_extended

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)


const ( 
	FIXED_DELTA float64 = 1.0 / 60.0
)

var gameManager_instance *gameManager


func GameManager() *gameManager {
	if gameManager_instance == nil {
		gameManager_instance = newGameManager()
	}

	return gameManager_instance
}

type gameManager struct {
	clock *Clock
	isRunning bool
	debug Debug
	world *World
	layers *Layers
}

func newGameManager() *gameManager {
	return &gameManager{clock: NewClock(), isRunning:  true, debug: *NewDebug(false), world: NewWorld()}
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
	return g.debug.Enabled()
}

func (g *gameManager) World() *World {
	return g.world
}

func (g *gameManager) SetIsDebug(debugFlag bool ){
	g.debug.SetEnabled(debugFlag)
	fmt.Printf("The debug flag is setted as %t", g.debug.Enabled()) 
}

func (g *gameManager) Update() {

	if g.isRunning {

		InputManager().Update()
		g.world.Update()
	}

}



func (g *gameManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Reset()
	g.world.Draw(target, op)
}