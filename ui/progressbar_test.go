package ui

import (
	"image/color"
	"testing"
)

func TestNewProgressBarNode(t *testing.T) {
	p := NewProgressBarNode("hp_bar", 200, 20)
	if p == nil {
		t.Fatal("NewProgressBarNode returned nil")
	}
	if p.GetProgress() != 0 {
		t.Errorf("expected initial progress 0, got %f", p.GetProgress())
	}
}

func TestProgressBarNodeSetProgress(t *testing.T) {
	p := NewProgressBarNode("bar", 100, 10)

	p.SetProgress(0.5)
	if p.GetProgress() != 0.5 {
		t.Errorf("expected 0.5, got %f", p.GetProgress())
	}

	p.SetProgress(-1)
	if p.GetProgress() != 0 {
		t.Errorf("expected clamp to 0 for negative, got %f", p.GetProgress())
	}

	p.SetProgress(1.5)
	if p.GetProgress() != 1 {
		t.Errorf("expected clamp to 1 for >1, got %f", p.GetProgress())
	}

	p.SetProgress(1)
	if p.GetProgress() != 1 {
		t.Errorf("expected 1, got %f", p.GetProgress())
	}
}

func TestProgressBarNodeSetFillColor(t *testing.T) {
	p := NewProgressBarNode("bar", 50, 10)
	p.SetFillColor(color.RGBA{0, 255, 0, 255})
	// No getter; just ensure no panic
}
