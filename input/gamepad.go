package input

import "github.com/hajimehoshi/ebiten/v2"

// GamePadButton identifies a gamepad button by ID and button index.
type GamePadButton struct {
	GamepadID ebiten.GamepadID
	Button    ebiten.GamepadButton
}

// NewGamePadButton creates a GamePadButton.
func NewGamePadButton(gamepadID ebiten.GamepadID, button ebiten.GamepadButton) GamePadButton {
	return GamePadButton{
		GamepadID: gamepadID,
		Button:    button,
	}
}

// GamePadAxis represents a gamepad axis (stick or trigger) with optional threshold.
// Use ebiten.StandardGamepadAxis values (e.g. StandardGamepadAxisLeftStickHorizontal) for standard layout.
type GamePadAxis struct {
	GamepadID ebiten.GamepadID
	Axis      ebiten.GamepadAxisType
	Threshold float64
	Above     bool // True if position must be above threshold, false if below
}

// NewGamePadAxis creates a GamePadAxis.
func NewGamePadAxis(gamepadID ebiten.GamepadID, axis ebiten.GamepadAxisType, threshold float64, above bool) GamePadAxis {
	return GamePadAxis{
		GamepadID: gamepadID,
		Axis:      axis,
		Threshold: threshold,
		Above:     above,
	}
}
