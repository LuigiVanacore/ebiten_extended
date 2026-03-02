package inputv3

import "github.com/hajimehoshi/ebiten/v2"

type InputState struct {
	isPressed  bool
	wasPressed bool
}

func (s InputState) IsPressed() bool {
    return s.isPressed
}

func (s InputState) WasPressed() bool {
    return s.wasPressed
}

func (s InputState) IsReleased() bool {
    return !s.isPressed && s.wasPressed
}

func (s InputState) IsHold() bool {
    return s.isPressed && s.wasPressed
}

func (s InputState) IsJustPressed() bool {
    return s.isPressed && !s.wasPressed
}

func (s InputState) IsJustReleased() bool {
    return !s.isPressed && s.wasPressed
}


type InputBuffer struct {
	keyStates     map[ebiten.Key]InputState
	mouseStates   map[ebiten.MouseButton]InputState
	gamepadStates map[GamePadButton]InputState
}

func NewInputBuffer() *InputBuffer {
	return &InputBuffer{
		keyStates:     make(map[ebiten.Key]InputState),
		mouseStates:   make(map[ebiten.MouseButton]InputState),
		gamepadStates: make(map[GamePadButton]InputState),
	}
}

func (i *InputBuffer) Update() {
    // Update previous states
    for k, v := range i.keyStates {
        v.wasPressed = v.isPressed
        i.keyStates[k] = v
    }
    for k, v := range i.mouseStates {
        v.wasPressed = v.isPressed
        i.mouseStates[k] = v
    }

    // Update current states
    for key := ebiten.Key(0); key <= ebiten.KeyMax; key++ {
        i.keyStates[key] = InputState{
            isPressed:  ebiten.IsKeyPressed(key),
            wasPressed: i.keyStates[key].wasPressed,
        }
    }

    for btn := ebiten.MouseButtonLeft; btn <= ebiten.MouseButtonRight; btn++ {
        i.mouseStates[btn] = InputState{
            isPressed:  ebiten.IsMouseButtonPressed(btn),
            wasPressed: i.mouseStates[btn].wasPressed,
        }
    }
}

func (i *InputBuffer) IsActionActive(action Action) bool {
	 switch action.GetActionType() {
    case KeyAction:
        state, exists := i.keyStates[action.GetKey()]
        if !exists {
            return false
        }
        
        isActive := false
        if action.GetMode().Has(Hold) {
            isActive = isActive || state.IsPressed()
        }
        if action.GetMode().Has(PressOnce) {
            isActive = isActive || (state.IsPressed() && !state.WasPressed())
        }
        if action.GetMode().Has(ReleaseOnce) {
            isActive = isActive || (!state.IsPressed() && state.WasPressed())
        }
        return isActive

    case MouseButtonAction:
        state, exists := i.mouseStates[action.GetMouseButton()]
        if !exists {
            return false
        }
        
        isActive := false
        if action.GetMode().Has(Hold) {
            isActive = isActive || state.IsPressed()
        }
        if action.GetMode().Has(PressOnce) {
            isActive = isActive || (state.IsPressed() && !state.WasPressed())
        }
        if action.GetMode().Has(ReleaseOnce) {
            isActive = isActive || (!state.IsPressed() && state.WasPressed())
        }
        return isActive
    }
    return false
}