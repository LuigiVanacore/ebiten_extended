package math2D

import "testing"

func TestNewRectangle(t *testing.T) {
	r := NewRectangle(NewVector2D(10, 20), NewVector2D(40, 30))
	if r.position.x != 10 || r.position.y != 20 || r.size.x != 40 || r.size.y != 30 {
		t.Errorf("NewRectangle: got pos=(%v,%v) size=(%v,%v)", r.position.x, r.position.y, r.size.x, r.size.y)
	}
}

func TestRectangle_GetPosition_SetPosition(t *testing.T) {
	r := NewRectangle(ZeroVector2D(), NewVector2D(10, 10))
	r.SetPosition(NewVector2D(5, 6))
	pos := r.GetPosition()
	if pos.X() != 5 || pos.Y() != 6 {
		t.Errorf("SetPosition/GetPosition: got (%v,%v), want (5,6)", pos.X(), pos.Y())
	}
}

func TestRectangle_GetSize_SetSize(t *testing.T) {
	r := NewRectangle(ZeroVector2D(), NewVector2D(10, 20))
	size := r.GetSize()
	if size.X() != 10 || size.Y() != 20 {
		t.Errorf("GetSize: got (%v,%v), want (10,20)", size.X(), size.Y())
	}
	r.SetSize(NewVector2D(50, 60))
	size = r.GetSize()
	if size.X() != 50 || size.Y() != 60 {
		t.Errorf("SetSize: got (%v,%v), want (50,60)", size.X(), size.Y())
	}
}

func TestRectangle_GetCenter_SetCenter(t *testing.T) {
	r := NewRectangle(NewVector2D(0, 0), NewVector2D(20, 10))
	center := r.GetCenter()
	if center.X() != 10 || center.Y() != 5 {
		t.Errorf("GetCenter: got (%v,%v), want (10,5)", center.X(), center.Y())
	}
	r.SetCenter(NewVector2D(50, 50))
	pos := r.GetPosition()
	if pos.X() != 40 || pos.Y() != 45 {
		t.Errorf("SetCenter(50,50) size(20,10): top-left = (%v,%v), want (40,45)", pos.X(), pos.Y())
	}
}

func TestRectangle_GetLeft_GetRight_GetTop_GetBottom(t *testing.T) {
	r := NewRectangle(NewVector2D(10, 20), NewVector2D(40, 30))
	if r.GetLeft() != 10 {
		t.Errorf("GetLeft = %v, want 10", r.GetLeft())
	}
	if r.GetRight() != 50 {
		t.Errorf("GetRight = %v, want 50", r.GetRight())
	}
	if r.GetTop() != 20 {
		t.Errorf("GetTop = %v, want 20", r.GetTop())
	}
	if r.GetBottom() != 50 {
		t.Errorf("GetBottom = %v, want 50", r.GetBottom())
	}
}

func TestRectangle_Translate(t *testing.T) {
	r := NewRectangle(NewVector2D(0, 0), NewVector2D(10, 10))
	r.Translate(5, 3)
	pos := r.GetPosition()
	if pos.X() != 5 || pos.Y() != 3 {
		t.Errorf("Translate(5,3): pos = (%v,%v), want (5,3)", pos.X(), pos.Y())
	}
}

func TestRectangle_Equal(t *testing.T) {
	r1 := NewRectangle(NewVector2D(0, 0), NewVector2D(10, 10))
	r2 := NewRectangle(NewVector2D(0, 0), NewVector2D(10, 10))
	if !r1.Equal(&r2) {
		t.Error("Equal rects should return true")
	}
	r3 := NewRectangle(NewVector2D(1, 0), NewVector2D(10, 10))
	if r1.Equal(&r3) {
		t.Error("Different rects should return false")
	}
}

func TestRectangle_Inflate(t *testing.T) {
	r := NewRectangle(NewVector2D(10, 20), NewVector2D(40, 30))
	r.Inflate(10, 6)
	pos := r.GetPosition()
	size := r.GetSize()
	if pos.X() != 5 || pos.Y() != 17 {
		t.Errorf("Inflate: pos = (%v,%v), want (5,17)", pos.X(), pos.Y())
	}
	if size.X() != 50 || size.Y() != 36 {
		t.Errorf("Inflate: size = (%v,%v), want (50,36)", size.X(), size.Y())
	}
}
