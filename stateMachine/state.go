package statemachine


type State interface {
	 Enter(owner any)
	 Execute(owner any)
	 Exit(owner any)
}