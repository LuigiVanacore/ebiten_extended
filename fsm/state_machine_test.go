package fsm

import "testing"

type mockState[T any] struct {
	enterCount int
	exitCount  int
	execCount  int
}

func (s *mockState[T]) Enter(owner T)   { s.enterCount++ }
func (s *mockState[T]) Exit(owner T)    { s.exitCount++ }
func (s *mockState[T]) Execute(owner T) { s.execCount++ }

func TestStateMachine_ChangeState(t *testing.T) {
	type owner struct{}
	sm := NewStateMachine[owner](owner{})
	s1 := &mockState[owner]{}
	s2 := &mockState[owner]{}

	sm.ChangeState(s1)
	if sm.GetCurrentState() != s1 {
		t.Error("expected s1 as current")
	}
	if s1.enterCount != 1 || s1.exitCount != 0 {
		t.Errorf("s1 enter=%d exit=%d", s1.enterCount, s1.exitCount)
	}

	sm.ChangeState(s2)
	if sm.GetCurrentState() != s2 {
		t.Error("expected s2 as current")
	}
	if s1.exitCount != 1 || s2.enterCount != 1 {
		t.Errorf("s1 exit=%d s2 enter=%d", s1.exitCount, s2.enterCount)
	}
}

func TestStateMachine_RevertToPreviousState(t *testing.T) {
	type owner struct{}
	sm := NewStateMachine[owner](owner{})
	s1 := &mockState[owner]{}
	s2 := &mockState[owner]{}

	sm.ChangeState(s1)
	sm.ChangeState(s2)
	sm.RevertToPreviousState()
	if sm.GetCurrentState() != s1 {
		t.Error("expected revert to s1")
	}
}

func TestStateMachine_IsInState(t *testing.T) {
	type owner struct{}
	sm := NewStateMachine[owner](owner{})
	s1 := &mockState[owner]{}
	sm.ChangeState(s1)
	if !sm.IsInState(s1) {
		t.Error("expected IsInState(s1)")
	}
}
