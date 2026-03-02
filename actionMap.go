package ebiten_extended

// ActionMap manages a collection of Actions mapped by arbitrary integer identifiers.
type ActionMap struct {
	actionMap map[int]Action
}

// NewActionMap initializes a new empty ActionMap.
func NewActionMap() *ActionMap {
	return &ActionMap{actionMap: make(map[int]Action)}
}

// Add registers an Action into the map under the specified integer key.
func (a *ActionMap) Add(actionKey int, action Action) {
	a.actionMap[actionKey] = action
}

// Get retrieves the Action associated with the specified key.
func (a *ActionMap) Get(actionKey int) Action {
	return a.actionMap[actionKey]
}
