package ebiten_extended

type ActionMap struct {
	actionMap map[int]Action
}

func NewActionMap() *ActionMap {
	return &ActionMap{actionMap: make(map[int]Action)}
}

func (a *ActionMap) Add(actionKey int, action Action) {
	a.actionMap[actionKey] = action
}

func (a *ActionMap) Get(actionKey int) Action {
	return a.actionMap[actionKey]
}
