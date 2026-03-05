package transform

import (
	"math"
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

func TestNewTransform(t *testing.T) {
	position := math2D.NewVector2D(1, 2)
	pivot := math2D.NewVector2D(3, 4)
	rotation := 45.0

	transform := NewTransform(position, pivot, rotation)

	if !transform.GetPosition().IsEqual(position) {
		t.Errorf("Expected position %v, got %v", position, transform.GetPosition())
	}
	if !transform.GetPivot().IsEqual(pivot) {
		t.Errorf("Expected pivot %v, got %v", pivot, transform.GetPivot())
	}
	if transform.GetRotation() != rotation {
		t.Errorf("Expected rotation %f, got %f", rotation, transform.GetRotation())
	}
}

func TestSetPosition(t *testing.T) {
	transform := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 0)
	transform.SetPosition(math2D.NewVector2D(5, 6))

	expected := math2D.NewVector2D(5, 6)
	if !transform.GetPosition().IsEqual(expected) {
		t.Errorf("Expected position %v, got %v", expected, transform.GetPosition())
	}
}

func TestSetRotation(t *testing.T) {
	transform := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 0)
	transform.SetRotation(90.0)

	if transform.GetRotation() != 90.0 {
		t.Errorf("Expected rotation 90, got %f", transform.GetRotation())
	}
}

func TestSetPivot(t *testing.T) {
	transform := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 0)
	pivot := math2D.NewVector2D(7, 8)
	transform.SetPivot(pivot.X(), pivot.Y())

	if !transform.GetPivot().IsEqual(pivot) {
		t.Errorf("Expected pivot %v, got %v", pivot, transform.GetPivot())
	}
}

func TestTranslate(t *testing.T) {
	transform := NewTransform(math2D.NewVector2D(1, 1), math2D.ZeroVector2D(), 0)
	transform.Translate(3, 4)

	expected := math2D.NewVector2D(4, 5)
	if !transform.GetPosition().IsEqual(expected) {
		t.Errorf("Expected position %v, got %v", expected, transform.GetPosition())
	}
}

func TestRotate(t *testing.T) {
	transform := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 45.0)
	transform.Rotate(15.0)

	if transform.GetRotation() != 60.0 {
		t.Errorf("Expected rotation 60, got %f", transform.GetRotation())
	}
}

func TestConcat(t *testing.T) {
	// Zero rotation: child position is simply added to parent position.
	t1 := NewTransform(math2D.NewVector2D(1, 1), math2D.ZeroVector2D(), 0)
	t2 := NewTransform(math2D.NewVector2D(2, 3), math2D.ZeroVector2D(), 0)
	t1.Concat(t2)
	if !t1.GetPosition().IsEqual(math2D.NewVector2D(3, 4)) {
		t.Errorf("zero-rot: expected (3,4), got %v", t1.GetPosition())
	}
	if t1.GetRotation() != 0 {
		t.Errorf("zero-rot: expected rotation 0, got %f", t1.GetRotation())
	}

	// 90° rotation: child at (1,0) local should appear at parent_pos + (0,1) world.
	// cos(π/2)≈0, sin(π/2)=1 → rotated (1,0) = (0,1).
	tol := 1e-9
	t3 := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), math.Pi/2)
	t4 := NewTransform(math2D.NewVector2D(1, 0), math2D.ZeroVector2D(), 0)
	t3.Concat(t4)
	gotX, gotY := t3.GetPosition().X(), t3.GetPosition().Y()
	if gotX > tol || gotX < -tol || gotY < 1-tol || gotY > 1+tol {
		t.Errorf("90°-rot: expected (0,1), got (%f,%f)", gotX, gotY)
	}

	// Rotations accumulate correctly.
	t5 := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 30.0)
	t6 := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 15.0)
	t5.Concat(t6)
	if t5.GetRotation() != 45.0 {
		t.Errorf("rotation accumulation: expected 45, got %f", t5.GetRotation())
	}
}