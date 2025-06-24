package inputv3

import (
	"github.com/LuigiVanacore/ebiten_extended/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type InputMode = utils.ByteSet

const (
	Hold InputMode = 1 << iota
	PressOnce
	ReleaseOnce
)
 
type ActionType int

const (
	KeyAction ActionType = iota
	MouseButtonAction
	GamePadButtonAction
	JoystickAxisAction
	EventAction
)

type Action struct {
	actionType    ActionType
	mode          InputMode
	key           ebiten.Key
	mouseButton   ebiten.MouseButton
	gamepadButton GamePadButton
	gamepadAxis  GamePadAxis
	//eventType     EventType
}

func NewKeyAction(key ebiten.Key, mode InputMode) Action {
	return Action{
		actionType: KeyAction,
		mode:       mode,
		key:        key,
	}
}

func NewMouseButtonAction(button ebiten.MouseButton,  mode InputMode) Action {
	return Action{
		actionType:  MouseButtonAction,
		mode:        mode,
		mouseButton: button,
	}
}

func NewGamePadButtonAction(button GamePadButton, mode InputMode) Action {
	return Action{
		actionType:    GamePadButtonAction,
		mode:          mode,
		gamepadButton: button,
	}
}

func NewJoystickAxisAction(axis GamePadAxis, mode InputMode) Action {
	return Action{
		actionType:   JoystickAxisAction,
		mode:         mode,
		gamepadAxis: axis,
	}
}

 
func (a Action) GetActionType() ActionType {
	return a.actionType
}

func (a Action) GetMode() utils.ByteSet {
	return a.mode
}

func (a Action) GetKey() ebiten.Key {
	return a.key
}

func (a Action) GetMouseButton() ebiten.MouseButton {
	return a.mouseButton
}

func (a Action) GetGamePadButton() GamePadButton {
	return a.gamepadButton
}

func (a Action) GetGamePadAxis() GamePadAxis {
	return a.gamepadAxis
}

func (a Action) GetActionInfo() ( ActionType, utils.ByteSet, ebiten.Key, ebiten.MouseButton, GamePadButton, GamePadAxis) {
	return a.actionType, a.mode, a.key, a.mouseButton, a.gamepadButton, a.gamepadAxis
}