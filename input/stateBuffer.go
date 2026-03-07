package input

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// InputState holds pressed state for one frame and the previous frame.
type InputState struct {
	IsPressed  bool
	WasPressed bool
}

// StateBuffer tracks keyboard, mouse, and gamepad button states for Action-based input.
type StateBuffer struct {
	keyStates     map[ebiten.Key]InputState
	mouseStates   map[ebiten.MouseButton]InputState
	gamepadStates map[GamePadButton]InputState
	gamepadIDs    []ebiten.GamepadID // reused each frame to avoid per-frame allocation
}

// NewStateBuffer creates a new StateBuffer.
func NewStateBuffer() *StateBuffer {
	return &StateBuffer{
		keyStates:     make(map[ebiten.Key]InputState),
		mouseStates:   make(map[ebiten.MouseButton]InputState),
		gamepadStates: make(map[GamePadButton]InputState),
		gamepadIDs:    make([]ebiten.GamepadID, 0, 4),
	}
}

// Update polls ebiten and updates all button states.
// Uses deterministic loops (no map range) to avoid undefined behavior when updating.
func (s *StateBuffer) Update() {
	// Keyboard: single pass, WasPressed = previous frame's IsPressed
	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		old := s.keyStates[key]
		s.keyStates[key] = InputState{
			IsPressed:  ebiten.IsKeyPressed(key),
			WasPressed: old.IsPressed,
		}
	}

	// Mouse: single pass
	for btn := ebiten.MouseButtonLeft; btn <= ebiten.MouseButtonRight; btn++ {
		old := s.mouseStates[btn]
		s.mouseStates[btn] = InputState{
			IsPressed:  ebiten.IsMouseButtonPressed(btn),
			WasPressed: old.IsPressed,
		}
	}

	// Gamepad: update connected, then age disconnected entries (avoid map range during modify)
	s.gamepadIDs = ebiten.AppendGamepadIDs(s.gamepadIDs[:0])
	connected := make(map[ebiten.GamepadID]bool, len(s.gamepadIDs))
	for _, id := range s.gamepadIDs {
		connected[id] = true
	}
	for _, id := range s.gamepadIDs {
		for btn := ebiten.GamepadButton(0); btn <= ebiten.GamepadButtonMax; btn++ {
			gp := NewGamePadButton(id, btn)
			old := s.gamepadStates[gp]
			s.gamepadStates[gp] = InputState{
				IsPressed:  ebiten.IsGamepadButtonPressed(id, btn),
				WasPressed: old.IsPressed,
			}
		}
	}
	// Age disconnected gamepad entries (WasPressed = last IsPressed)
	var toAge []GamePadButton
	for gp := range s.gamepadStates {
		if !connected[gp.GamepadID] {
			toAge = append(toAge, gp)
		}
	}
	for _, gp := range toAge {
		old := s.gamepadStates[gp]
		s.gamepadStates[gp] = InputState{IsPressed: old.IsPressed, WasPressed: old.IsPressed}
	}
}

// ConnectedGamepadIDs returns the slice of connected gamepad IDs populated during the last Update.
// The slice is reused across frames; do not store a reference to it.
func (s *StateBuffer) ConnectedGamepadIDs() []ebiten.GamepadID {
	return s.gamepadIDs
}

// IsActionActive returns whether the action is currently active according to its mode.
func (s *StateBuffer) IsActionActive(action Action) bool {
	switch action.GetActionType() {
	case KeyAction:
		state := s.keyStates[action.GetKey()]
		return s.checkState(state, action.GetMode())

	case MouseButtonAction:
		state := s.mouseStates[action.GetMouseButton()]
		return s.checkState(state, action.GetMode())

	case GamePadButtonAction:
		gp := action.GetGamePadButton()
		state := s.gamepadStates[gp]
		return s.checkState(state, action.GetMode())

	case JoystickAxisAction:
		axis := action.GetGamePadAxis()
		val := axisValue(axis.GamepadID, axis.Axis)
		pressed := (axis.Above && val > axis.Threshold) || (!axis.Above && val < -axis.Threshold)
		if action.GetMode().Has(ActionHold) || action.GetMode().Has(ActionPressOnce) {
			return pressed
		}
		if action.GetMode().Has(ActionReleaseOnce) {
			return !pressed
		}
		return false
	}
	return false
}

func (s *StateBuffer) checkState(state InputState, mode ActionMode) bool {
	isActive := false
	if mode.Has(ActionHold) {
		isActive = isActive || state.IsPressed
	}
	if mode.Has(ActionPressOnce) {
		isActive = isActive || (state.IsPressed && !state.WasPressed)
	}
	if mode.Has(ActionReleaseOnce) {
		isActive = isActive || (!state.IsPressed && state.WasPressed)
	}
	return isActive
}

func axisValue(id ebiten.GamepadID, axis ebiten.GamepadAxisType) float64 {
	if ebiten.IsStandardGamepadLayoutAvailable(id) {
		return ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxis(axis))
	}
	return ebiten.GamepadAxisValue(id, axis)
}
