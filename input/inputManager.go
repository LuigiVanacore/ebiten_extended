package input

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)



func NewInputManager() *InputManager {
	i := &InputManager{}
	i.keySlice = make([]ebiten.Key, 0, 4)

	return i
}

type InputManager struct {
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



func (i *InputManager) SetMouseEnabled(value bool) {
	i.mouseEnabled = value
}

func (i *InputManager) IsMouseEnabled() bool {
	return i.mouseEnabled
}

func (i *InputManager) GetCursorPos() math2D.Vector2D {
	return i.cursorPos
}

func (i *InputManager) Update() {

	if i.mouseEnabled {
		x, y := ebiten.CursorPosition()
		i.cursorPos.SetPosition(float64(x), float64(y))
	}
}



