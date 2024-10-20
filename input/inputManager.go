package input

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)



var instanceInputManager *inputManager

func InputManager() *inputManager {
	if instanceInputManager == nil {
		instanceInputManager = newInputManager()
	}

	return instanceInputManager
}
 

func newInputManager() *inputManager {
	i := &inputManager{}
	i.keySlice = make([]ebiten.Key, 0, 4)

	return i
}

type inputManager struct {
	activeContexts []*InputContext

	keySlice        []ebiten.Key


	mouseEnabled bool
	mouseHasDrag          bool
	mouseDragging         bool
	mouseJustHadDrag      bool
	mouseJustReleasedDrag bool
	mousePressed          bool
	mouseStartPos         math2D.Vector2D
	mouseDragPos          math2D.Vector2D
	cursorPos    math2D.Vector2D
}



func (i *inputManager) SetMouseEnabled(value bool) {
	i.mouseEnabled = value
}

func (i *inputManager) IsMouseEnabled() bool {
	return i.mouseEnabled
}

func (i *inputManager) GetCursorPos() math2D.Vector2D {
	return i.cursorPos
}

func (i *inputManager) Update() {

	if i.mouseEnabled {
		x, y := ebiten.CursorPosition()
		i.cursorPos.SetPosition(float64(x), float64(y))
	}
}



