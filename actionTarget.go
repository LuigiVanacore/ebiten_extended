package ebiten_extended

// ActionTarget manages event binding, coupling logical input identifiers (from an ActionMap) to callback functions.
type ActionTarget struct {
	eventRealTimeMap map[int]func(args ...any)
	eventPoolMap     map[int]func(args ...any)
	actionMap        *ActionMap
}

// NewActionTarget creates a new ActionTarget bound to the specified ActionMap.
func NewActionTarget(actionMap *ActionMap) *ActionTarget {
	return &ActionTarget{eventRealTimeMap: make(map[int]func(args ...any)), eventPoolMap: make(map[int]func(args ...any)), actionMap: actionMap}
}

// ProcessEvent manually triggers the pooled callback for a specific action ID, returning true if found in the pool.
func (a *ActionTarget) ProcessEvent(actionId int) bool {
	f := a.eventPoolMap[actionId]
	if f != nil {
		f()
		return true
	}
	return false
}

// ProcessEvents iterates over real-time bound actions, evaluating their state and executing active callbacks.
func (a *ActionTarget) ProcessEvents(args ...any) {
	for actionId, fun := range a.eventRealTimeMap {
		action := a.actionMap.Get(actionId)
		if action.Test() {
			fun(args...)
		}
	}
}

// Bind links a callback function to a designated action instance tracked by the given actionId.
func (a *ActionTarget) Bind(actionId int, callback func(args ...any)) {
	action := a.actionMap.Get(actionId)
	if action.inputType == PRESSED {
		a.eventRealTimeMap[actionId] = callback
	} else {
		a.eventPoolMap[actionId] = callback
	}
}

// UnBind removes the previously assigned callback mapping for a specific action ID.
func (a *ActionTarget) UnBind(actionId int) {
	action := a.actionMap.Get(actionId)
	if action.inputType == PRESSED {
		delete(a.eventRealTimeMap, actionId)
	} else {
		delete(a.eventPoolMap, actionId)
	}
}
