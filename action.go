package ebiten_extended

import "github.com/hajimehoshi/ebiten/v2"

const (
	REAL_TIME = iota
	PRESSED
	RELEASED
)

const (
	KEY_BUTTON = iota
	MOUSE_BUTTON
	GAMEPAD_BUTTON
)

type Action struct {
	inputType     int
	buttonType    int
	keyButton     ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton ebiten.GamepadButton
	gamepadID     ebiten.GamepadID
}

func NewActionKey(keyButton ebiten.Key, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: KEY_BUTTON, keyButton: keyButton}
}

func NewActionMouse(mouseButton ebiten.MouseButton, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: MOUSE_BUTTON, mouseButton: mouseButton}
}

func NewActionGamepad(gamepadButton ebiten.GamepadButton, gamepadID ebiten.GamepadID, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: GAMEPAD_BUTTON, gamepadButton: gamepadButton, gamepadID: gamepadID}
}

func (a *Action) Test() bool {
	res := false

	if a.buttonType == KEY_BUTTON {
		if a.inputType == PRESSED {
			res = ebiten.IsKeyPressed(a.keyButton)
		}
	} else if a.buttonType == MOUSE_BUTTON {
		if a.inputType == PRESSED {
			res = ebiten.IsMouseButtonPressed(a.mouseButton)
		}
	}
	return res
}
