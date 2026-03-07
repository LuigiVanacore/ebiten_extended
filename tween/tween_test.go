package tween

import "testing"

func TestTween_Set(t *testing.T) {
	tw := NewTween("t", 0, 100, 10, Linear)
	val, done := tw.Set(0)
	if val != 0 || done {
		t.Errorf("Set(0): val=%v done=%v", val, done)
	}
	val, done = tw.Set(10)
	if val != 100 || !done {
		t.Errorf("Set(10): val=%v done=%v", val, done)
	}
	val, done = tw.Set(5)
	if val != 50 || done {
		t.Errorf("Set(5): val=%v done=%v", val, done)
	}
}

func TestTween_StepDelta(t *testing.T) {
	tw := NewTween("t", 0, 10, 1, Linear)
	val, done := tw.StepDelta(0.5)
	if val != 5 || done {
		t.Errorf("StepDelta(0.5): val=%v done=%v", val, done)
	}
	val, done = tw.StepDelta(0.5)
	if val != 10 || !done {
		t.Errorf("StepDelta(0.5) again: val=%v done=%v", val, done)
	}
}

func TestTween_Reset(t *testing.T) {
	tw := NewTween("t", 0, 100, 10, Linear)
	tw.Set(5)
	tw.Reset()
	val, _ := tw.Set(0)
	if val != 0 {
		t.Errorf("Reset: val=%v", val)
	}
}
