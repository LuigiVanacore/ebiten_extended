package ebiten_extended

import "github.com/hajimehoshi/ebiten/v2"

// Input state constants.
const (
	REAL_TIME = iota
	PRESSED
	RELEASED
)

// Hardware source constants for input actions.
const (
	KEY_BUTTON = iota
	MOUSE_BUTTON
	GAMEPAD_BUTTON
)

// Action encapsulates an input event mapping (keyboard, mouse, or gamepad) to be tested.
type Action struct {
	inputType     int
	buttonType    int
	keyButton     ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton ebiten.GamepadButton
	gamepadID     ebiten.GamepadID
}

// NewActionKey creates a keyboard-based Action.
func NewActionKey(keyButton ebiten.Key, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: KEY_BUTTON, keyButton: keyButton}
}

// NewActionMouse creates a mouse-based Action.
func NewActionMouse(mouseButton ebiten.MouseButton, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: MOUSE_BUTTON, mouseButton: mouseButton}
}

// NewActionGamepad creates a gamepad-based Action tied to a specific gamepad ID.
func NewActionGamepad(gamepadButton ebiten.GamepadButton, gamepadID ebiten.GamepadID, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: GAMEPAD_BUTTON, gamepadButton: gamepadButton, gamepadID: gamepadID}
}

// Test evaluates the current input state against the action configuration and returns true if triggered.
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
