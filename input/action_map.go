package input

import "github.com/hajimehoshi/ebiten/v2"

// ActionID identifies an action in the ActionMap (e.g. a const or enum value).
type ActionID int

// ActionMap maps ActionID to Action for the state-based input system.
type ActionMap struct {
	actions map[ActionID]Action
}

// NewActionMap creates an empty ActionMap.
func NewActionMap() *ActionMap {
	return &ActionMap{
		actions: make(map[ActionID]Action),
	}
}

// AddAction binds an action to the given ID.
func (am *ActionMap) AddAction(id ActionID, action Action) {
	am.actions[id] = action
}

// RemoveAction removes the action for the given ID.
func (am *ActionMap) RemoveAction(id ActionID) {
	delete(am.actions, id)
}

// GetAction returns the action for the given ID, and whether it exists.
func (am *ActionMap) GetAction(id ActionID) (Action, bool) {
	action, exists := am.actions[id]
	return action, exists
}

// ClearActions removes all actions.
func (am *ActionMap) ClearActions() {
	am.actions = make(map[ActionID]Action)
}

// SetKeyBinding rebinds the action for the given ID to a new keyboard key.
// Returns an error if the action does not exist. Only KeyAction types are modified;
// other action types are replaced with a new KeyAction.
func (am *ActionMap) SetKeyBinding(id ActionID, key ebiten.Key) error {
	action, exists := am.actions[id]
	if !exists {
		return ErrActionNotFound
	}
	am.actions[id] = NewKeyAction(key, action.GetMode())
	return nil
}

// GetKeyForAction returns the keyboard key bound to the action, and true if the action exists and is a KeyAction.
func (am *ActionMap) GetKeyForAction(id ActionID) (ebiten.Key, bool) {
	action, exists := am.actions[id]
	if !exists || action.GetActionType() != KeyAction {
		return 0, false
	}
	return action.GetKey(), true
}

// ErrActionNotFound is returned when an action ID is not found in the map.
var ErrActionNotFound = &ActionNotFoundError{}

// ActionNotFoundError indicates the action ID was not found.
type ActionNotFoundError struct{}

func (e *ActionNotFoundError) Error() string {
	return "action not found"
}
