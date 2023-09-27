package ebiten_extended

type ActionTarget struct {
	eventRealTimeMap map[int]func(args ...any)
	eventPoolMap     map[int]func(args ...any)
	actionMap        *ActionMap
}

func NewActionTarget(actionMap *ActionMap) *ActionTarget {
	return &ActionTarget{eventRealTimeMap: make(map[int]func(args ...any)), eventPoolMap: make(map[int]func(args ...any)), actionMap: actionMap}
}

func (a *ActionTarget) ProcessEvent(actionId int) bool {
	f := a.eventPoolMap[actionId]
	if f != nil {
		f()
		return true
	}
	return false
}

func (a *ActionTarget) ProcessEvents(args ...any) {
	for actionId, fun := range a.eventRealTimeMap {
		action := a.actionMap.Get(actionId)
		if action.Test() {
			fun(args...)
		}
	}
}

func (a *ActionTarget) Bind(actionId int, callback func(args ...any)) {
	action := a.actionMap.Get(actionId)
	if action.inputType == PRESSED {
		a.eventRealTimeMap[actionId] = callback
	} else {
		a.eventPoolMap[actionId] = callback
	}
}

func (a *ActionTarget) UnBind(actionId int) {
	action := a.actionMap.Get(actionId)
	if action.inputType == PRESSED {
		delete(a.eventRealTimeMap, actionId)
	} else {
		delete(a.eventPoolMap, actionId)
	}
}
