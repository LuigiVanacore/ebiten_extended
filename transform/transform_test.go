package transform

import (
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
	transform1 := NewTransform(math2D.NewVector2D(1, 1), math2D.ZeroVector2D(), 30.0)
	transform2 := NewTransform(math2D.NewVector2D(2, 3), math2D.ZeroVector2D(), 15.0)

	transform1.Concat(transform2)

	expectedPosition := math2D.NewVector2D(3, 4)
	expectedRotation := 45.0

	if !transform1.GetPosition().IsEqual(expectedPosition) {
		t.Errorf("Expected position %v, got %v", expectedPosition, transform1.GetPosition())
	}
	if transform1.GetRotation() != expectedRotation {
		t.Errorf("Expected rotation %f, got %f", expectedRotation, transform1.GetRotation())
	}
}