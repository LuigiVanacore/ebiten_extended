package input

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
