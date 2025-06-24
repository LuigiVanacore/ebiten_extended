package fsm



type Stateable interface {
	Enter()
	Exit()
	Update()
}