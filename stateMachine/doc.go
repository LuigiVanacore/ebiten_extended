// Package statemachine provides a simple state machine for AI or game flow.
//
// StateMachine has an owner (any), current state, previous state, and optional global state.
// States implement the State interface (Execute(owner), Enter(owner), Exit(owner)).
// Update runs the global state (if set) then the current state. ChangeState switches
// states and calls Exit/Enter; RevertToPreviousState restores the previous state.
package statemachine
