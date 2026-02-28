package input

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

// ActionBinding maps a list of abstract physical inputs to a single virtual string.
type ActionBinding []RawInputButton

type InputManager struct {
	actions map[string]ActionBinding

	mouseEnabled bool
	cursorPos    math2D.Vector2D
}

func NewInputManager() *InputManager {
	return &InputManager{
		actions: make(map[string]ActionBinding),
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

// GetCursorPos returns the World mapped Vector2D of the Hardware mouse.
func (i *InputManager) GetCursorPos() math2D.Vector2D {
	return i.cursorPos
}

// Update polls hardware limits based on features (like mouse motion).
// Native buttons are automatically polled by ebiten via ebiten.IsKeyPressed etc.
func (i *InputManager) Update() {
	if i.mouseEnabled {
		x, y := ebiten.CursorPosition()
		i.cursorPos.SetPosition(float64(x), float64(y))
	}
}
