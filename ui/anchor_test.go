package ui

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

func TestAnchorLayoutTopLeft(t *testing.T) {
	parent := NewPanelNode("parent", 100, 100)
	parent.SetPosition(10, 10)

	child := NewPanelNode("child", 20, 20)
	child.SetAnchor(AnchorTopLeft)
	child.SetAnchorMargin(math2D.NewVector2D(5, 5))
	parent.AddChildren(child)

	layout := NewAnchorLayout()
	layout.Apply(parent)

	pos := child.GetPosition()
	if pos.X() != 15 || pos.Y() != 15 {
		t.Errorf("Expected 15,15, got %v,%v", pos.X(), pos.Y())
	}
}

func TestAnchorLayoutCenter(t *testing.T) {
	parent := NewPanelNode("parent", 100, 100)
	child := NewPanelNode("child", 50, 50)
	child.SetAnchor(AnchorCenter)
	parent.AddChildren(child)

	layout := NewAnchorLayout()
	layout.Apply(parent)

	pos := child.GetPosition()
	if pos.X() != 25 || pos.Y() != 25 {
		t.Errorf("Expected 25,25, got %v,%v", pos.X(), pos.Y())
	}
}

func TestAnchorLayoutStretch(t *testing.T) {
	parent := NewPanelNode("parent", 100, 100)
	child := NewPanelNode("child", 10, 10)
	child.SetAnchor(AnchorStretch)
	child.SetAnchorMargin(math2D.NewVector2D(10, 10))
	parent.AddChildren(child)

	layout := NewAnchorLayout()
	layout.Apply(parent)

	w, h := child.GetSize()
	if w != 80 || h != 80 {
		t.Errorf("Expected size 80,80, got %v,%v", w, h)
	}

	pos := child.GetPosition()
	if pos.X() != 10 || pos.Y() != 10 {
		t.Errorf("Expected pos 10,10, got %v,%v", pos.X(), pos.Y())
	}
}
