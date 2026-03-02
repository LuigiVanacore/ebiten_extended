package input

import (
	"time"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)


type InputMode int

const (
	Hold InputMode = iota
	PRESSED
	RELEASED
)

type InputType int

const (
	KEY_BUTTON InputType = iota
	MOUSE_BUTTON
	GAMEPAD_BUTTON
)


type InputEvent struct {
	inputType        InputType
	inputMode    InputMode

	keyButton     ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton ebiten.GamepadButton
	gamepadID     ebiten.GamepadID

	Duration time.Duration
	Pos      math2D.Vector2D
	StartPos math2D.Vector2D
	hasPos      bool 
	hasStartPos bool
}


func NewInputKeyEvent(keyButton ebiten.Key) *InputEvent {
	return &InputEvent{inputType: KEY_BUTTON, keyButton: keyButton}
}

func NewInputMouseEvent(mouseButton ebiten.MouseButton) *InputEvent {
	return &InputEvent{inputType: MOUSE_BUTTON, mouseButton: mouseButton}
}

func NewInputGamepadEvent(gamepadButton ebiten.GamepadButton, gamepadID ebiten.GamepadID) *InputEvent {
	return &InputEvent{inputType: GAMEPAD_BUTTON, gamepadButton: gamepadButton, gamepadID: gamepadID}
}

func (a *InputEvent) Test() bool {
	res := false

	if a.inputType == KEY_BUTTON {
		if a.inputMode == PRESSED {
			res = ebiten.IsKeyPressed(a.keyButton)
		}
	} else if a.inputType == MOUSE_BUTTON {
		if a.inputMode == PRESSED {
			res = ebiten.IsMouseButtonPressed(a.mouseButton)
		}
	}
	return res
}


