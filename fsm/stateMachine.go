package fsm



type StateID int

type StateMachine struct {
	states      map[StateID]Stateable
	currentState Stateable
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		states:      make(map[StateID]Stateable),
		currentState: nil,
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
		if sm.currentState != nil {
			sm.currentState.Exit()
		}
		sm.currentState = state
		sm.currentState.Enter()
	}
  }

  func (sm *StateMachine) GetCurrentState() Stateable {
	return sm.currentState
  }

  func (sm *StateMachine) Update() {
	 if sm.currentState != nil {
   sm.currentState.Update()
	 }	

	   }

	     func (sm *StateMachine) GetStates() map[StateID]Stateable {
			return sm.states
	  }

	  // Reset resets the state machine to its initial state.
	  func (sm *StateMachine) Reset() {
		   if sm.currentState != nil {
	 sm.currentState.Exit()
   }

   sm.currentState = nil
   for _, state := range sm.states {
	 state.Exit()
   }

   sm.states = make(map[StateID]Stateable)
   }
   // Clear clears all states in the state machine.
   func (sm *StateMachine) Clear() {
	 if sm.currentState != nil {
													sm.currentState.Exit()
  }						

	 sm.currentState = nil
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

 