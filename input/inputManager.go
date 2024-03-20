package input

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten"
)



var instanceInputManager *inputManager

func ResourceManager() *inputManager {
	if instanceInputManager == nil {
		instanceInputManager = newInputManager()
	}

	return instanceInputManager
}
 

func newInputManager() *inputManager {
	i := &inputManager{}
	return i
}

type inputManager struct {
	activeContexts []*InputContext
	//callbackTable  map[int]InputCallback

	mouseEnabled bool
	cursorPos    math2D.Vector2D
}

func (i *inputManager) Update() {

	if i.mouseEnabled {
		x, y := ebiten.CursorPosition()
		i.cursorPos.SetPosition(float64(x), float64(y))
	}
}
