package fsm

// StateMachine manages the state transitions and execution for a specific owner of type T.
type StateMachine[T any] struct {
	owner         T
	currentState  State[T]
	previousState State[T]
	globalState   State[T]
}

// NewStateMachine initializes a new StateMachine with the given owner.
func NewStateMachine[T any](owner T) *StateMachine[T] {
	return &StateMachine[T]{owner: owner}
}

// GetCurrentState returns the active state.
func (sm *StateMachine[T]) GetCurrentState() State[T] {
	return sm.currentState
}

// GetGlobalState returns the global state that executes alongside the current state.
func (sm *StateMachine[T]) GetGlobalState() State[T] {
	return sm.globalState
}

// GetPreviousState returns the state that was active before the current one.
func (sm *StateMachine[T]) GetPreviousState() State[T] {
	return sm.previousState
}

// SetCurrentState forcibly sets the current state without triggering Exit/Enter sequences.
func (sm *StateMachine[T]) SetCurrentState(state State[T]) {
	sm.currentState = state
}

// SetGlobalState sets a global state that updates every tick regardless of the current state.
func (sm *StateMachine[T]) SetGlobalState(state State[T]) {
	sm.globalState = state
}

// SetPreviousState manually sets the previous state.
func (sm *StateMachine[T]) SetPreviousState(state State[T]) {
	sm.previousState = state
}

// Update evaluates the global state and then the current state.
func (sm *StateMachine[T]) Update() {
	if sm.globalState != nil {
		sm.globalState.Execute(sm.owner)
	}

	if sm.currentState != nil {
		sm.currentState.Execute(sm.owner)
	}
}

// ChangeState transitions from the current state to the given state, triggering Exit and Enter.
func (sm *StateMachine[T]) ChangeState(state State[T]) {
	sm.previousState = sm.currentState

	if sm.currentState != nil {
		sm.currentState.Exit(sm.owner)
	}

	sm.currentState = state

	if sm.currentState != nil {
		sm.currentState.Enter(sm.owner)
	}
}

// RevertToPreviousState transitions back to the immediately preceding state.
func (sm *StateMachine[T]) RevertToPreviousState() {
	sm.ChangeState(sm.previousState)
}

// IsInState returns true if the specified state is the current active state.
func (sm *StateMachine[T]) IsInState(state State[T]) bool {
	return sm.currentState == state
}
