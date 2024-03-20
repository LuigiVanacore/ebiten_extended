package ebiten_extended

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended"
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
	scenes []*ebiten_extended.SceneNode
	isRunning bool
	isDebug bool 
}

func newGameManager() *gameManager {
	return &gameManager{scenes: make([]*ebiten_extended.SceneNode, 0), isRunning:  true}
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

func (g *gameManager) SetIsDebug(debugFlag bool ){
	g.isDebug = debugFlag
	fmt.Printf("The debug flag is setted as %t", g.isDebug)
}

func (g *gameManager) AddNode(node *ebiten_extended.SceneNode) {
	g.scenes = append(g.scenes, node)
}

func (g *gameManager) Update() {
	
	if g.isRunning {
		SceneManager().Update()
		
	}

}



func (g *gameManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	SceneManager().Draw(target, op)
}