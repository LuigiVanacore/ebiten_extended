package fsm

// State represents a generic state interface for the Finite State Machine.
// Type T represents the owner (entity) of the state.
type State[T any] interface {
	Enter(owner T)
	Execute(owner T)
	Exit(owner T)
}
