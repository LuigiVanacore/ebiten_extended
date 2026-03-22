package input

import (
	"github.com/LuigiVanacore/ludum/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// ActionMode is a bit mask for action trigger modes.
type ActionMode = utils.ByteSet

const (
	ActionHold ActionMode = 1 << iota
	ActionPressOnce
	ActionReleaseOnce
)

// ActionType identifies the kind of input an Action uses.
type ActionType int

const (
	KeyAction ActionType = iota
	MouseButtonAction
	GamePadButtonAction
	JoystickAxisAction
	EventAction
)

// Action represents a single input trigger (key, mouse button, gamepad button, or axis).
type Action struct {
	actionType    ActionType
	mode          ActionMode
	key           ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton GamePadButton
	gamepadAxis   GamePadAxis
}

// NewKeyAction creates an action bound to a keyboard key.
func NewKeyAction(key ebiten.Key, mode ActionMode) Action {
	return Action{
		actionType: KeyAction,
		mode:       mode,
		key:        key,
	}
}

// NewMouseButtonAction creates an action bound to a mouse button.
func NewMouseButtonAction(button ebiten.MouseButton, mode ActionMode) Action {
	return Action{
		actionType:  MouseButtonAction,
		mode:        mode,
		mouseButton: button,
	}
}

// NewGamePadButtonAction creates an action bound to a gamepad button.
func NewGamePadButtonAction(button GamePadButton, mode ActionMode) Action {
	return Action{
		actionType:    GamePadButtonAction,
		mode:          mode,
		gamepadButton: button,
	}
}

// NewJoystickAxisAction creates an action bound to a gamepad axis with threshold.
func NewJoystickAxisAction(axis GamePadAxis, mode ActionMode) Action {
	return Action{
		actionType:  JoystickAxisAction,
		mode:        mode,
		gamepadAxis: axis,
	}
}

// GetActionType returns the action's input type.
func (a Action) GetActionType() ActionType {
	return a.actionType
}

// GetMode returns the action's trigger mode.
func (a Action) GetMode() utils.ByteSet {
	return a.mode
}

// GetKey returns the keyboard key (only valid for KeyAction).
func (a Action) GetKey() ebiten.Key {
	return a.key
}

// GetMouseButton returns the mouse button (only valid for MouseButtonAction).
func (a Action) GetMouseButton() ebiten.MouseButton {
	return a.mouseButton
}

// GetGamePadButton returns the gamepad button (only valid for GamePadButtonAction).
func (a Action) GetGamePadButton() GamePadButton {
	return a.gamepadButton
}

// GetGamePadAxis returns the gamepad axis (only valid for JoystickAxisAction).
func (a Action) GetGamePadAxis() GamePadAxis {
	return a.gamepadAxis
}
