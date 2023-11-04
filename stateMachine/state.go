package statemachine


type State interface {
	 Enter(owner any)
	 Execute(pwner any)
	 Exit(pwner any)
}