package ui

import (
	"image/color"
	"testing"
)

func TestNewPanelNode(t *testing.T) {
	p := NewPanelNode("test", 100, 50)
	if p == nil {
		t.Fatal("NewPanelNode returned nil")
	}
	w, h := p.GetSize()
	if w != 100 || h != 50 {
		t.Errorf("expected size 100x50, got %.0fx%.0f", w, h)
	}
}

func TestPanelNodeSetSize(t *testing.T) {
	p := NewPanelNode("test", 10, 10)
	p.SetSize(200, 80)
	w, h := p.GetSize()
	if w != 200 || h != 80 {
		t.Errorf("expected size 200x80 after SetSize, got %.0fx%.0f", w, h)
	}
}

func TestPanelNodeContainsPoint(t *testing.T) {
	p := NewPanelNode("test", 100, 100)
	p.SetPosition(10, 20)

	if !p.ContainsPoint(10, 20) {
		t.Error("expected (10,20) top-left to be inside")
	}
	if !p.ContainsPoint(60, 70) {
		t.Error("expected (60,70) center to be inside")
	}
	if !p.ContainsPoint(109, 119) {
		t.Error("expected (109,119) bottom-right to be inside")
	}
	if p.ContainsPoint(111, 20) {
		t.Error("expected (111,20) outside right to be false")
	}
	if p.ContainsPoint(10, 121) {
		t.Error("expected (10,121) outside bottom to be false")
	}
	if p.ContainsPoint(9, 20) {
		t.Error("expected (9,20) outside left to be false")
	}
}

func TestPanelNodeSetBackgroundColor(t *testing.T) {
	p := NewPanelNode("test", 50, 50)
	p.SetBackgroundColor(color.RGBA{255, 0, 0, 255})
	// Can't easily assert - just ensure no panic
}

func TestPanelNodeSetLayer(t *testing.T) {
	p := NewPanelNode("test", 50, 50)
	p.SetLayer(3)
	if p.GetLayer() != 3 {
		t.Errorf("expected layer 3, got %d", p.GetLayer())
	}
}
