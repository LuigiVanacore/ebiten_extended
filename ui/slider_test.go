package ui

import (
	"testing"
)

func TestNewSliderNode(t *testing.T) {
	// SliderNode requires InputManager; pass nil - Update will no-op
	sl := NewSliderNode("slider", 200, 20, nil)
	if sl == nil {
		t.Fatal("NewSliderNode returned nil")
	}
	if sl.GetValue() != 0.5 {
		t.Errorf("expected default value 0.5, got %f", sl.GetValue())
	}
}

func TestSliderNodeSetValue(t *testing.T) {
	sl := NewSliderNode("slider", 100, 15, nil)

	sl.SetValue(0.8)
	if sl.GetValue() != 0.8 {
		t.Errorf("expected 0.8, got %f", sl.GetValue())
	}

	sl.SetValue(-1)
	if sl.GetValue() != 0 {
		t.Errorf("expected clamp to 0, got %f", sl.GetValue())
	}

	sl.SetValue(1.5)
	if sl.GetValue() != 1 {
		t.Errorf("expected clamp to 1, got %f", sl.GetValue())
	}
}

func TestSliderNodeSetRange(t *testing.T) {
	sl := NewSliderNode("slider", 100, 15, nil)
	sl.SetRange(10, 100)
	if sl.GetValue() != 10 {
		t.Errorf("expected clamp to 10 after SetRange, got %f", sl.GetValue())
	}
	sl.SetValue(55)
	if sl.GetValue() != 55 {
		t.Errorf("expected 55, got %f", sl.GetValue())
	}
	sl.SetValue(5)
	if sl.GetValue() != 10 {
		t.Errorf("expected clamp to 10, got %f", sl.GetValue())
	}
}
