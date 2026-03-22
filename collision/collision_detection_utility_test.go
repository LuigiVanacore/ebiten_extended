package collision

import (
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
)

func TestClampOnRange(t *testing.T) {
	tests := []struct {
		x, min, max, want float64
	}{
		{5, 0, 10, 5},
		{-1, 0, 10, 0},
		{15, 0, 10, 10},
	}
	for _, tt := range tests {
		got := ClampOnRange(tt.x, tt.min, tt.max)
		if got != tt.want {
			t.Errorf("ClampOnRange(%v,%v,%v)=%v, want %v", tt.x, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestClampOnRectangle(t *testing.T) {
	r := math2d.NewRectangle(math2d.NewVector2D(10, 20), math2d.NewVector2D(30, 40))
	// point inside
	p := math2d.NewVector2D(25, 35)
	clamped := ClampOnRectangle(p, r)
	if clamped.X() != 25 || clamped.Y() != 35 {
		t.Errorf("Point inside: clamped = (%v,%v), want (25,35)", clamped.X(), clamped.Y())
	}
	// point outside (right)
	p2 := math2d.NewVector2D(50, 35)
	clamped2 := ClampOnRectangle(p2, r)
	if clamped2.X() != 40 || clamped2.Y() != 35 {
		t.Errorf("Point right: clamped = (%v,%v), want (40,35)", clamped2.X(), clamped2.Y())
	}
}

func TestOverlapping(t *testing.T) {
	if !Overlapping(0, 10, 5, 15) {
		t.Error("Overlapping(0,10,5,15) should be true")
	}
	if Overlapping(0, 5, 10, 15) {
		t.Error("Overlapping(0,5,10,15) should be false")
	}
}
