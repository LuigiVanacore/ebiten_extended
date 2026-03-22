package input

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// ActionBinding maps a list of abstract physical inputs to a single virtual string.
type ActionBinding []RawInputButton

type InputManager struct {
	actions      map[string]ActionBinding
	actionMap    *ActionMap
	stateBuf     *StateBuffer
	mouseEnabled bool
	cursorPos    math2D.Vector2D
}

func NewInputManager() *InputManager {
	return &InputManager{
		actions:   make(map[string]ActionBinding),
		actionMap: NewActionMap(),
		stateBuf:  NewStateBuffer(),
	}
}

// AddAction binds a string name to an abstract RawInputButton (Keyboard, Mouse, or Gamepad).
// If the action name already exists, the new button is appended as an alternative trigger.
func (i *InputManager) AddAction(actionName string, button RawInputButton) {
	if _, exists := i.actions[actionName]; !exists {
		i.actions[actionName] = ActionBinding{}
	}
	i.actions[actionName] = append(i.actions[actionName], button)
}

// RemoveAction entirely deletes the action mapping.
func (i *InputManager) RemoveAction(actionName string) {
	delete(i.actions, actionName)
}

// IsActionPressed returns true if ANY of the buttons bound to this action are currently held down.
func (i *InputManager) IsActionPressed(actionName string) bool {
	bindings, exists := i.actions[actionName]
	if !exists {
		return false
	}
	for _, button := range bindings {
		if button.IsPressed() {
			return true
		}
	}
	return false
}

// IsActionJustPressed returns true only on the first frame ANY bound button is pressed.
func (i *InputManager) IsActionJustPressed(actionName string) bool {
	bindings, exists := i.actions[actionName]
	if !exists {
		return false
	}
	for _, button := range bindings {
		if button.IsJustPressed() {
			return true
		}
	}
	return false
}

// IsActionJustReleased returns true only on the exact frame ANY bound button is let go.
func (i *InputManager) IsActionJustReleased(actionName string) bool {
	bindings, exists := i.actions[actionName]
	if !exists {
		return false
	}
	for _, button := range bindings {
		if button.IsJustReleased() {
			return true
		}
	}
	return false
}

// IsActionReleased returns true only if ALL buttons bound to this action are currently unpressed.
func (i *InputManager) IsActionReleased(actionName string) bool {
	bindings, exists := i.actions[actionName]
	if !exists {
		return true // By default, unused actions are considered released
	}
	for _, button := range bindings {
		if button.IsPressed() {
			return false
		}
	}
	return true
}

// IsActionHeld returns true if ANY of the buttons bound to this action have been held down
// for at least the specified number of ticks (frames).
func (i *InputManager) IsActionHeld(actionName string, minFrames int) bool {
	bindings, exists := i.actions[actionName]
	if !exists {
		return false
	}
	for _, button := range bindings {
		if button.PressDuration() >= minFrames {
			return true
		}
	}
	return false
}

// SetMouseEnabled determines whether the cursor position will be polled during Update.
func (i *InputManager) SetMouseEnabled(value bool) {
	i.mouseEnabled = value
}

// IsMouseEnabled returns the current mouse-polling state.
func (i *InputManager) IsMouseEnabled() bool {
	return i.mouseEnabled
}

// GetCursorPos returns the absolute screen coordinates (OS Window) of the Hardware mouse.
// To get the world-mapped coordinates, use Camera.GetCursorCoords(inputMgr) instead.
func (i *InputManager) GetCursorPos() math2D.Vector2D {
	return i.cursorPos
}

// Update polls hardware (keys, mouse, gamepad) and updates cursor if enabled.
func (i *InputManager) Update() {
	i.stateBuf.Update()
	if i.mouseEnabled {
		x, y := ebiten.CursorPosition()
		i.cursorPos.SetPosition(math2D.NewVector2D(float64(x), float64(y)))
	}
}

// SetActionMap replaces the ActionMap used for RegisterAction / IsActionActive.
func (i *InputManager) SetActionMap(m *ActionMap) {
	i.actionMap = m
}

// RegisterAction binds an Action to an ActionID for the state-based API.
func (i *InputManager) RegisterAction(id ActionID, action Action) {
	i.actionMap.AddAction(id, action)
}

// IsActionActive returns true if the action for the given ID is currently active.
func (i *InputManager) IsActionActive(id ActionID) bool {
	action, ok := i.actionMap.GetAction(id)
	if !ok {
		return false
	}
	return i.stateBuf.IsActionActive(action)
}

// GamepadIDs returns the IDs of connected gamepads as of the last Update call.
// The returned slice is reused across frames; do not store a reference to it.
func (i *InputManager) GamepadIDs() []ebiten.GamepadID {
	return i.stateBuf.ConnectedGamepadIDs()
}

// StickDeadzone is applied to axis values; inputs below this magnitude are treated as zero.
const StickDeadzone = 0.15

func applyDeadzone(v, deadzone float64) float64 {
	if v < -deadzone {
		return (v + deadzone) / (1 - deadzone)
	}
	if v > deadzone {
		return (v - deadzone) / (1 - deadzone)
	}
	return 0
}

// GetLeftStick returns the left stick (x, y) for the given gamepad in [-1, 1] with deadzone.
func (i *InputManager) GetLeftStick(id ebiten.GamepadID) (x, y float64) {
	if ebiten.IsStandardGamepadLayoutAvailable(id) {
		x = applyDeadzone(ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickHorizontal), StickDeadzone)
		y = applyDeadzone(ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisLeftStickVertical), StickDeadzone)
		return x, y
	}
	if ebiten.GamepadAxisCount(id) >= 2 {
		x = applyDeadzone(ebiten.GamepadAxisValue(id, 0), StickDeadzone)
		y = applyDeadzone(ebiten.GamepadAxisValue(id, 1), StickDeadzone)
		return x, y
	}
	return 0, 0
}

// GetRightStick returns the right stick (x, y) for the given gamepad.
func (i *InputManager) GetRightStick(id ebiten.GamepadID) (x, y float64) {
	if ebiten.IsStandardGamepadLayoutAvailable(id) {
		x = applyDeadzone(ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickHorizontal), StickDeadzone)
		y = applyDeadzone(ebiten.StandardGamepadAxisValue(id, ebiten.StandardGamepadAxisRightStickVertical), StickDeadzone)
		return x, y
	}
	if ebiten.GamepadAxisCount(id) >= 4 {
		x = applyDeadzone(ebiten.GamepadAxisValue(id, 2), StickDeadzone)
		y = applyDeadzone(ebiten.GamepadAxisValue(id, 3), StickDeadzone)
		return x, y
	}
	return 0, 0
}

// IsGamepadButtonPressed reports whether the button is currently held.
func (i *InputManager) IsGamepadButtonPressed(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	return ebiten.IsGamepadButtonPressed(id, button)
}

// IsGamepadButtonJustPressed reports whether the button was pressed this frame.
func (i *InputManager) IsGamepadButtonJustPressed(id ebiten.GamepadID, button ebiten.GamepadButton) bool {
	return inpututil.IsGamepadButtonJustPressed(id, button)
}

// IsStandardGamepadButtonPressed reports whether the standard button is held.
func (i *InputManager) IsStandardGamepadButtonPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return ebiten.IsStandardGamepadButtonPressed(id, button)
}

// IsStandardGamepadButtonJustPressed reports whether the standard button was pressed this frame.
func (i *InputManager) IsStandardGamepadButtonJustPressed(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return inpututil.IsStandardGamepadButtonJustPressed(id, button)
}

// IsStandardGamepadButtonJustReleased reports whether the standard button was released this frame.
func (i *InputManager) IsStandardGamepadButtonJustReleased(id ebiten.GamepadID, button ebiten.StandardGamepadButton) bool {
	return inpututil.IsStandardGamepadButtonJustReleased(id, button)
}
