package tween

import (
	"math"
	"testing"
)

const floatTolerance float32 = 1e-5

func floatEqual(a, b float32) bool {
	return float32(math.Abs(float64(a-b))) <= floatTolerance
}

func TestNewTween(t *testing.T) {
	tw := NewTween("test", 0, 100, 1.0, Linear)
	if tw == nil {
		t.Fatal("NewTween should not return nil")
	}
	if tw.begin != 0 || tw.end != 100 || tw.duration != 1.0 || tw.change != 100 {
		t.Errorf("NewTween: begin=%v end=%v duration=%v change=%v", tw.begin, tw.end, tw.duration, tw.change)
	}
}

func TestTween_Set_AtStart(t *testing.T) {
	tw := NewTween("test", 10, 110, 1.0, Linear)
	current, finished := tw.Set(0)
	if current != 10 || finished {
		t.Errorf("Set(0): got (%v, %v), want (10, false)", current, finished)
	}
}

func TestTween_Set_AtEnd(t *testing.T) {
	tw := NewTween("test", 0, 100, 1.0, Linear)
	current, finished := tw.Set(1.0)
	if current != 100 || !finished {
		t.Errorf("Set(1.0): got (%v, %v), want (100, true)", current, finished)
	}
}

func TestTween_Set_NegativeTime(t *testing.T) {
	tw := NewTween("test", 5, 15, 1.0, Linear)
	current, finished := tw.Set(-1)
	if current != 5 || finished {
		t.Errorf("Set(-1): got (%v, %v), want (5, false)", current, finished)
	}
}

func TestTween_Set_LinearMidpoint(t *testing.T) {
	tw := NewTween("test", 0, 100, 2.0, Linear)
	current, finished := tw.Set(1.0)
	if current != 50 || finished {
		t.Errorf("Set(1.0) Linear 0->100 over 2s: got (%v, %v), want (50, false)", current, finished)
	}
}

func TestTween_Set_OverDuration(t *testing.T) {
	tw := NewTween("test", 0, 100, 1.0, Linear)
	current, finished := tw.Set(2.0)
	if current != 100 || !finished {
		t.Errorf("Set(2.0) over 1s: got (%v, %v), want (100, true)", current, finished)
	}
}

func TestTween_Reset(t *testing.T) {
	tw := NewTween("test", 0, 100, 1.0, Linear)
	tw.Set(0.5)
	tw.Reset()
	current, _ := tw.Set(0)
	if current != 0 {
		t.Errorf("Reset: Set(0) returned %v, want 0", current)
	}
}

func TestTween_Update(t *testing.T) {
	tw := NewTween("test", 20, 80, 1.0, Linear)
	tw.Set(0.5)
	current, finished := tw.Update()
	if current != 50 || finished {
		t.Errorf("Update after Set(0.5): got (%v, %v), want (50, false)", current, finished)
	}
}

func TestLinear(t *testing.T) {
	// Linear: c*t/d + b
	v := Linear(0.5, 10, 90, 1.0)
	expected := float32(55) // 10 + 90*0.5 = 55
	if v != expected {
		t.Errorf("Linear(0.5, 10, 90, 1) = %v, want %v", v, expected)
	}
}

func TestInQuad(t *testing.T) {
	v := InQuad(0.5, 0, 100, 1.0)
	expected := float32(25)
	if !floatEqual(v, expected) {
		t.Errorf("InQuad(0.5,0,100,1) = %v, want ~25", v)
	}
}

func TestOutQuad(t *testing.T) {
	v := OutQuad(1.0, 0, 100, 1.0)
	if v != 100 {
		t.Errorf("OutQuad(1,0,100,1) = %v, want 100", v)
	}
}

func TestInOutQuad(t *testing.T) {
	v := InOutQuad(1.0, 0, 100, 1.0)
	if v != 100 {
		t.Errorf("InOutQuad(1,0,100,1) = %v, want 100", v)
	}
}

func TestInCubic(t *testing.T) {
	v := InCubic(1.0, 0, 100, 1.0)
	if v != 100 {
		t.Errorf("InCubic(1,0,100,1) = %v, want 100", v)
	}
}

func TestInSine(t *testing.T) {
	v := InSine(1.0, 0, 100, 1.0)
	if !floatEqual(v, 100) {
		t.Errorf("InSine(1,0,100,1) = %v, want ~100", v)
	}
}

func TestInExpo(t *testing.T) {
	v := InExpo(0, 0, 100, 1.0)
	if v != 0 {
		t.Errorf("InExpo(0,...) = %v, want 0", v)
	}
}

func TestOutBounce(t *testing.T) {
	v := OutBounce(1.0, 0, 100, 1.0)
	if v != 100 {
		t.Errorf("OutBounce(1,0,100,1) = %v, want 100", v)
	}
}

func TestInBack(t *testing.T) {
	v := InBack(1.0, 0, 100, 1.0)
	if v != 100 {
		t.Errorf("InBack(1,0,100,1) = %v, want 100", v)
	}
}
