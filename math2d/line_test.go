package math2D

import "testing"

func TestNewLine(t *testing.T) {
	l := NewLine(NewVector2D(0, 0), NewVector2D(1, 0))
	base := l.GetBase()
	dir := l.GetDirection()
	if base.X() != 0 || base.Y() != 0 || dir.X() != 1 || dir.Y() != 0 {
		t.Errorf("NewLine: got base=(%v,%v) dir=(%v,%v)", base.X(), base.Y(), dir.X(), dir.Y())
	}
}

func TestLine_SetBase_SetDirection(t *testing.T) {
	l := NewLine(ZeroVector2D(), ZeroVector2D())
	l.SetBase(NewVector2D(5, 6))
	l.SetDirection(NewVector2D(0, 1))
	base := l.GetBase()
	dir := l.GetDirection()
	if base.X() != 5 || base.Y() != 6 || dir.X() != 0 || dir.Y() != 1 {
		t.Errorf("SetBase/SetDirection: got base=(%v,%v) dir=(%v,%v)", base.X(), base.Y(), dir.X(), dir.Y())
	}
}
