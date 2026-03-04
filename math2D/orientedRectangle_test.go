package math2D

import "testing"

func TestNewOrientedRectangle(t *testing.T) {
	center := NewVector2D(10, 20)
	half := NewVector2D(5, 7)
	rot := 1.25

	r := NewOrientedRectangle(center, half, rot)

	if !r.GetCenter().IsEqual(center) {
		t.Fatalf("center mismatch: got %v, want %v", r.GetCenter(), center)
	}
	if !r.GetHalfExtended().IsEqual(half) {
		t.Fatalf("halfExtended mismatch: got %v, want %v", r.GetHalfExtended(), half)
	}
	if r.GetRotation() != rot {
		t.Fatalf("rotation mismatch: got %v, want %v", r.GetRotation(), rot)
	}
}
