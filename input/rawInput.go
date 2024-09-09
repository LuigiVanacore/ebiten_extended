package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type InputButtonTypes int

const (
	KEYBOARD InputButtonTypes = iota
	MOUSE
	GAMEPAD
)

type RawInputButton struct {
	keyButton ebiten.Key
	mouseButton ebiten.MouseButton
	gamePadButton ebiten.GamepadButton
	gamepadID ebiten.GamepadID
	buttonType InputButtonTypes
}

func NewKeyRawInputButton(key ebiten.Key) RawInputButton {
	return RawInputButton{
		keyButton: key,
		buttonType: KEYBOARD,
	}
}

func NewMouseRawInputButton(mouse ebiten.MouseButton) RawInputButton {
	return RawInputButton{
		mouseButton: mouse,
		buttonType: MOUSE,
	}
}

func NewGamePadRawInputButton(gamePad ebiten.GamepadButton, gamePadId ebiten.GamepadID) RawInputButton {
	return RawInputButton{
		gamePadButton: gamePad,
		gamepadID: gamePadId,
		buttonType: GAMEPAD,
	}
}

func (r *RawInputButton) IsKey(key ebiten.Key) bool {
	return r.keyButton == key
}

func (r *RawInputButton) IsMouseButton(mouse ebiten.MouseButton) bool {
	return r.mouseButton == mouse
}

func (r *RawInputButton) IsGamePadButton(gamePad ebiten.GamepadButton) bool {
	return r.gamePadButton == gamePad
}

func (r *RawInputButton) IsKeyType() bool {
	return r.buttonType == KEYBOARD
}

func (r *RawInputButton) IsMouseType() bool {
	return r.buttonType == MOUSE
}

func (r *RawInputButton) IsGamepadType() bool {
	return r.buttonType == GAMEPAD
}


func (r *RawInputButton) IsGamepadID(gamepadID ebiten.GamepadID) bool {
	return r.gamepadID == gamepadID
}

func (r *RawInputButton) IsEqual(other RawInputButton) bool {
	if r.buttonType != other.buttonType {
		return false
	}

	switch r.buttonType {
	case KEYBOARD:
		return r.keyButton == other.keyButton
	case MOUSE:
		return r.mouseButton == other.mouseButton
	case GAMEPAD:
		return r.gamePadButton == other.gamePadButton && r.gamepadID == other.gamepadID
	}

	return false
}

func (r *RawInputButton) IsEqualKey(key ebiten.Key) bool {
	if r.buttonType != KEYBOARD {
		return false
	}

	return r.keyButton == key
}

func (r *RawInputButton) IsEqualMouseButton(mouse ebiten.MouseButton) bool {
	if r.buttonType != MOUSE {
		return false
	}

	return r.mouseButton == mouse
}

func (r *RawInputButton) IsEqualGamepadButton(gamePad ebiten.GamepadButton) bool {
	if r.buttonType != GAMEPAD {
		return false
	}

	return r.gamePadButton == gamePad
}

func (r *RawInputButton) IsEqualType(buttonType InputButtonTypes) bool {
	return r.buttonType == buttonType
}

func (r *RawInputButton) IsPressed() bool {
	switch r.buttonType {
	case KEYBOARD:
		return ebiten.IsKeyPressed(r.keyButton)
	case MOUSE:
		return ebiten.IsMouseButtonPressed(r.mouseButton)
	case GAMEPAD:
		return ebiten.IsGamepadButtonPressed(r.gamepadID, r.gamePadButton)
	}

	return false
}

func (r *RawInputButton) IsJustPressed() bool {
	switch r.buttonType {
	case KEYBOARD:
		return inpututil.IsKeyJustPressed(r.keyButton)
	case MOUSE:
		return inpututil.IsMouseButtonJustPressed(r.mouseButton)
	case GAMEPAD:
		return inpututil.IsGamepadButtonJustPressed(r.gamepadID, r.gamePadButton)
	}

	return false
}


func (r *RawInputButton) IsJustReleased() bool {
	switch r.buttonType {
	case KEYBOARD:
		return inpututil.IsKeyJustReleased(r.keyButton)
	case MOUSE:
		return inpututil.IsMouseButtonJustReleased(r.mouseButton)
	case GAMEPAD:
		return inpututil.IsGamepadButtonJustReleased(r.gamepadID, r.gamePadButton)
	}

	return false
}

func (r *RawInputButton) IsReleased() bool {
	switch r.buttonType {
	case KEYBOARD:
		return !ebiten.IsKeyPressed(r.keyButton)
	case MOUSE:
		return !ebiten.IsMouseButtonPressed(r.mouseButton)
	case GAMEPAD:
		return !ebiten.IsGamepadButtonPressed(r.gamepadID, r.gamePadButton)
	}

	return false
}

func (r *RawInputButton) PressDuration() int {
	switch r.buttonType {
	case KEYBOARD:
		return inpututil.KeyPressDuration(r.keyButton)
	case MOUSE:
		return inpututil.MouseButtonPressDuration(r.mouseButton)
	case GAMEPAD:
		return inpututil.GamepadButtonPressDuration(r.gamepadID, r.gamePadButton)
	}

	return 0
}



type RawInputAxis struct {
	axis int
	gamepadID ebiten.GamepadID
	axisValues     [8]float64
	prevAxisValues [8]float64
}