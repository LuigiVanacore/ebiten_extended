package ebiten_extended

import (

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
	scenes []*Node
	isRunning bool
}

func newGameManager() *gameManager {
	return &gameManager{isRunning: true}
}

func (g *gameManager) AddNode(node *Node) {
	g.scenes = append(g.scenes, node)
}

func (g *gameManager) Update() {
	
	if g.isRunning {
		g.updateScene(fixedDelta)
	}

}

func (g *gameManager) updateScene(dt float64) {
	for i := range g.scenes {
		g.scenes[i].Update(dt)
	}
}

func (g *gameManager) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	for i := range g.scenes {
		g.scenes[i].Draw(target, op)
	}
}