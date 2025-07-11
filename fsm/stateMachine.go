package fsm

type Trigger func() bool

type Transition struct {
	Origin  StateID
	Target  StateID
	Trigger Trigger
}

type StateID int

type StateMachine struct {
	states         map[StateID]Stateable
	transitions    map[StateID][]Transition
	currentStateID StateID
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		states:      make(map[StateID]Stateable),
		transitions: make(map[StateID][]Transition),
	}
}

func (sm *StateMachine) AddState(id StateID, state Stateable) {
	sm.states[id] = state
}

func (sm *StateMachine) RemoveState(id StateID) {
	delete(sm.states, id)
}

func (sm *StateMachine) SetState(id StateID) {
	if state, exists := sm.states[id]; exists {
		if curr, ok := sm.states[sm.currentStateID]; ok {
			curr.Exit()
		}
		sm.currentStateID = id
		state.Enter()
	}
}

func (sm *StateMachine) GetCurrentState() Stateable {
	return sm.states[sm.currentStateID]
}

func (sm *StateMachine) Update() {

	if sm.IsEmpty() {
		return
	}
	
	sm.checkTransitions()

	if state, ok := sm.states[sm.currentStateID]; ok {
		state.Update()
	}
}

func (sm *StateMachine) AddTransitions(transitions ...Transition) {
	for _, transition := range transitions {
		_, originExists := sm.states[transition.Origin]
		_, targetExists := sm.states[transition.Target]
		if originExists && targetExists {
			sm.transitions[transition.Origin] = append(sm.transitions[transition.Origin], transition)
		}
	}
}

func (sm *StateMachine) RemoveTransition(origin StateID, target StateID) {
	if transitions, exists := sm.transitions[origin]; exists {
		for i, transition := range transitions {
			if transition.Target == target {
				sm.transitions[origin] = append(transitions[:i], transitions[i+1:]...)
				break
			}
		}
	}
}

func (sm *StateMachine) GetTransitions() map[StateID][]Transition {
	return sm.transitions
}

func (sm *StateMachine) GetStates() map[StateID]Stateable {
	return sm.states
}

// Reset resets the state machine to its initial state.
func (sm *StateMachine) Reset() {
	if curr, ok := sm.states[sm.currentStateID]; ok {
		curr.Exit()
	}
	sm.currentStateID = 0
	for _, state := range sm.states {
		state.Exit()
	}
	sm.states = make(map[StateID]Stateable)
}

// Clear clears all states in the state machine.
func (sm *StateMachine) Clear() {
	if curr, ok := sm.states[sm.currentStateID]; ok {
		curr.Exit()
	}
	sm.currentStateID = 0
	sm.states = make(map[StateID]Stateable)
}

// IsEmpty checks if the state machine has no states.
func (sm *StateMachine) IsEmpty() bool {
	return len(sm.states) == 0
}

// IsState checks if a state with the given ID exists in the state machine.
func (sm *StateMachine) IsState(id StateID) bool {
	_, exists := sm.states[id]
	return exists
}

func (sm *StateMachine) checkTransitions() {
	if _, ok := sm.states[sm.currentStateID]; !ok {
		return
	}
	if transitions, ok := sm.transitions[sm.currentStateID]; ok {
		for _, transition := range transitions {
			if transition.Trigger() {
				sm.SetState(transition.Target)
				break
			}
		}
	}
}
