package ebiten_extended

import "github.com/hajimehoshi/ebiten/v2"

// Input type constants used to determine how an action should be evaluated.
const (
	REAL_TIME = iota // REAL_TIME actions are evaluated every frame
	PRESSED          // PRESSED actions trigger when the button is down
	RELEASED         // RELEASED actions trigger when the button is released
)

// Button type constants specifying the input device to test.
const (
	KEY_BUTTON = iota
	MOUSE_BUTTON
	GAMEPAD_BUTTON
)

// Action represents a single input trigger such as a key, mouse button,
// or gamepad button.
type Action struct {
	inputType     int
	buttonType    int
	keyButton     ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton ebiten.GamepadButton
	gamepadID     ebiten.GamepadID
}

// NewActionKey creates an action bound to a keyboard key.
func NewActionKey(keyButton ebiten.Key, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: KEY_BUTTON, keyButton: keyButton}
}

// NewActionMouse creates an action bound to a mouse button.
func NewActionMouse(mouseButton ebiten.MouseButton, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: MOUSE_BUTTON, mouseButton: mouseButton}
}

// NewActionGamepad creates an action bound to a gamepad button.
func NewActionGamepad(gamepadButton ebiten.GamepadButton, gamepadID ebiten.GamepadID, inputType int) *Action {
	return &Action{inputType: inputType, buttonType: GAMEPAD_BUTTON, gamepadButton: gamepadButton, gamepadID: gamepadID}
}

// Test reports whether the action is active according to its input and button type.
// For REAL_TIME and PRESSED inputs it returns true while the button is held.
// For RELEASED inputs it returns true when the button is not pressed.
func (a *Action) Test() bool {
	switch a.buttonType {
	case KEY_BUTTON:
		return testKey(a)
	case MOUSE_BUTTON:
		return testMouse(a)
	case GAMEPAD_BUTTON:
		return testGamepad(a)
	}
	return false
}

func testKey(a *Action) bool {
	pressed := ebiten.IsKeyPressed(a.keyButton)
	switch a.inputType {
	case PRESSED, REAL_TIME:
		return pressed
	case RELEASED:
		return !pressed
	}
	return false
}

func testMouse(a *Action) bool {
	pressed := ebiten.IsMouseButtonPressed(a.mouseButton)
	switch a.inputType {
	case PRESSED, REAL_TIME:
		return pressed
	case RELEASED:
		return !pressed
	}
	return false
}

func testGamepad(a *Action) bool {
	pressed := ebiten.IsGamepadButtonPressed(a.gamepadID, a.gamepadButton)
	switch a.inputType {
	case PRESSED, REAL_TIME:
		return pressed
	case RELEASED:
		return !pressed
	}
	return false
}
