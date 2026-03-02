<<<<<<< HEAD
// Package statemachine provides a simple state machine for AI or game flow.
//
// StateMachine has an owner (any), current state, previous state, and optional global state.
// States implement the State interface (Execute(owner), Enter(owner), Exit(owner)).
// Update runs the global state (if set) then the current state. ChangeState switches
// states and calls Exit/Enter; RevertToPreviousState restores the previous state.
=======
// Package statemachine contains an alternative state machine used by example code.
>>>>>>> 153f371edcb4dcf68c2d6633071e13a31c6b0c07
package statemachine
