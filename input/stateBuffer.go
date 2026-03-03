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
}

// NewStateBuffer creates a new StateBuffer.
func NewStateBuffer() *StateBuffer {
	return &StateBuffer{
		keyStates:     make(map[ebiten.Key]InputState),
		mouseStates:   make(map[ebiten.MouseButton]InputState),
		gamepadStates: make(map[GamePadButton]InputState),
	}
}

// Update polls ebiten and updates all button states.
func (s *StateBuffer) Update() {
	for k, v := range s.keyStates {
		v.WasPressed = v.IsPressed
		s.keyStates[k] = v
	}
	for k, v := range s.mouseStates {
		v.WasPressed = v.IsPressed
		s.mouseStates[k] = v
	}
	for k, v := range s.gamepadStates {
		v.WasPressed = v.IsPressed
		s.gamepadStates[k] = v
	}

	for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
		s.keyStates[key] = InputState{
			IsPressed:  ebiten.IsKeyPressed(key),
			WasPressed: s.keyStates[key].WasPressed,
		}
	}

	for btn := ebiten.MouseButtonLeft; btn <= ebiten.MouseButtonRight; btn++ {
		s.mouseStates[btn] = InputState{
			IsPressed:  ebiten.IsMouseButtonPressed(btn),
			WasPressed: s.mouseStates[btn].WasPressed,
		}
	}

	ids := ebiten.AppendGamepadIDs(nil)
	for _, id := range ids {
		for btn := ebiten.GamepadButton(0); btn <= ebiten.GamepadButtonMax; btn++ {
			gp := NewGamePadButton(id, btn)
			s.gamepadStates[gp] = InputState{
				IsPressed:  ebiten.IsGamepadButtonPressed(id, btn),
				WasPressed: s.gamepadStates[gp].WasPressed,
			}
		}
	}
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
