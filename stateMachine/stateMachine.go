package statemachine


type StateMachine struct {
	owner any
	currentState State
	previousState State
	globalState State
}
  


func NewStateMachine(owner any) *StateMachine {
	return &StateMachine{owner: owner}
}

func (stateMachine *StateMachine) GetCurrentState() State {
	return stateMachine.currentState
}

func (stateMachine *StateMachine) GetGlobalState() State {
	return stateMachine.globalState
}

func (stateMachine *StateMachine) GetPreviousState() State {
	return stateMachine.previousState
}  

func (stateMachine *StateMachine) SetCurrentState(state State) {
	stateMachine.currentState = state
}

func (stateMachine *StateMachine) SetGlobalState(state State) {
	stateMachine.globalState = state
}

func (stateMachine *StateMachine) SetPreviousState(state State) {
	stateMachine.previousState = state
}

func (stateMachine *StateMachine) Update() {
	if stateMachine.globalState != nil {
		stateMachine.globalState.Execute(stateMachine.owner)
	}

	if stateMachine.currentState != nil {
		stateMachine.currentState.Execute(stateMachine.owner)
	}
}


func (stateMachine *StateMachine) ChangeState(state State) {
	stateMachine.previousState = stateMachine.currentState

	stateMachine.currentState.Exit(stateMachine.owner)

	stateMachine.currentState = state

	stateMachine.currentState.Enter(stateMachine.owner)
}



func (stateMachine *StateMachine) RevertToPreviousState() {
	stateMachine.ChangeState(stateMachine.previousState)
}


func (stateMachine *StateMachine) IsInState(state State) bool {
	return stateMachine.currentState == state
}

