package math2D

import "testing"

func TestNewCircle(t *testing.T) {
	c := NewCircle(NewVector2D(10, 20), 5)
	if c.center.x != 10 || c.center.y != 20 || c.radius != 5 {
		t.Errorf("NewCircle: got center=(%v,%v) r=%v, want (10,20) 5", c.center.x, c.center.y, c.radius)
	}
}

func TestCircle_SetCenter_GetCenterPosition(t *testing.T) {
	c := NewCircle(ZeroVector2D(), 1)
	c.SetCenter(NewVector2D(7, 11))
	center := c.GetCenterPosition()
	if center.X() != 7 || center.Y() != 11 {
		t.Errorf("SetCenter/GetCenterPosition: got (%v,%v), want (7,11)", center.X(), center.Y())
	}
}

func TestCircle_GetRadius_SetRadius(t *testing.T) {
	c := NewCircle(ZeroVector2D(), 3)
	if c.GetRadius() != 3 {
		t.Errorf("GetRadius = %v, want 3", c.GetRadius())
	}
	c.SetRadius(10)
	if c.GetRadius() != 10 {
		t.Errorf("SetRadius: GetRadius = %v, want 10", c.GetRadius())
	}
}
