package math2d

import "testing"

func TestNewSegment(t *testing.T) {
	s := NewSegment(NewVector2D(0, 0), NewVector2D(10, 0))
	start := s.GetStartPoint()
	end := s.GetEndPoint()
	if start.X() != 0 || start.Y() != 0 || end.X() != 10 || end.Y() != 0 {
		t.Errorf("NewSegment: got (%v,%v)-(%v,%v)", start.X(), start.Y(), end.X(), end.Y())
	}
}

func TestSegment_SetStartPoint_SetEndPoint(t *testing.T) {
	s := NewSegment(ZeroVector2D(), ZeroVector2D())
	s.SetStartPoint(NewVector2D(1, 2))
	s.SetEndPoint(NewVector2D(3, 4))
	if s.startPoint.x != 1 || s.startPoint.y != 2 || s.endPoint.x != 3 || s.endPoint.y != 4 {
		t.Errorf("SetStartPoint/SetEndPoint: got (%v,%v)-(%v,%v)", s.startPoint.x, s.startPoint.y, s.endPoint.x, s.endPoint.y)
	}
}

func TestSegment_ProjectSegment(t *testing.T) {
	s := NewSegment(NewVector2D(0, 0), NewVector2D(10, 0))
	axis := NewVector2D(1, 0)
	r := s.ProjectSegment(axis, true)
	if r.GetMinimum() != 0 || r.GetMaximum() != 10 {
		t.Errorf("ProjectSegment onto (1,0): got (%v,%v), want (0,10)", r.GetMinimum(), r.GetMaximum())
	}
}
