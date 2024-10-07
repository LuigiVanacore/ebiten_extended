package input

//import "github.com/hajimehoshi/ebiten/v2"

type InputContext struct {
	actionMap map[RawInputButton]Action

}

func NewInputContext() *InputContext {
	return &InputContext{
		actionMap: make(map[RawInputButton]Action),
	}
}

func (i *InputContext) MapButtonToAction(button RawInputButton) Action {
	return i.actionMap[button]
}

func (i *InputContext) MapActionToButton(action Action) (RawInputButton, bool) {
	for button, act := range i.actionMap {
		if act == action {
			return button, true
		}
	}
	return RawInputButton{}, false
}

