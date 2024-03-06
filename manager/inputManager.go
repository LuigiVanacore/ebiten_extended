package manager


type InputManager struct {
	activeContexts []*input.InputContext
	callbackTable  map[int]InputCallback
}