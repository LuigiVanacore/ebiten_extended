package transform

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

func TestNewTransform(t *testing.T) {
	position := math2D.NewVector2D(1, 2)
	pivot := math2D.NewVector2D(3, 4)
	rotation := 45

	transform := NewTransform(position, pivot, rotation)

	if !transform.GetPosition().IsEqual(position) {
		t.Errorf("Expected position %v, got %v", position, transform.GetPosition())
	}
	if !transform.GetPivot().IsEqual(pivot) {
		t.Errorf("Expected pivot %v, got %v", pivot, transform.GetPivot())
	}
	if transform.GetRotation() != rotation {
		t.Errorf("Expected rotation %d, got %d", rotation, transform.GetRotation())
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
	transform.SetRotation(90)

	if transform.GetRotation() != 90 {
		t.Errorf("Expected rotation 90, got %d", transform.GetRotation())
	}
}

func TestSetPivot(t *testing.T) {
	transform := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 0)
	transform.SetPivot(math2D.NewVector2D(7, 8))

	expected := math2D.NewVector2D(7, 8)
	if !transform.GetPivot().IsEqual(expected) {
		t.Errorf("Expected pivot %v, got %v", expected, transform.GetPivot())
	}
}

func TestTranslate(t *testing.T) {
	transform := NewTransform(math2D.NewVector2D(1, 1), math2D.ZeroVector2D(), 0)
	transform.Translate(math2D.NewVector2D(3, 4))

	expected := math2D.NewVector2D(4, 5)
	if !transform.GetPosition().IsEqual(expected) {
		t.Errorf("Expected position %v, got %v", expected, transform.GetPosition())
	}
}

func TestRotate(t *testing.T) {
	transform := NewTransform(math2D.ZeroVector2D(), math2D.ZeroVector2D(), 45)
	transform.Rotate(15)

	if transform.GetRotation() != 60 {
		t.Errorf("Expected rotation 60, got %d", transform.GetRotation())
	}
}

func TestConcat(t *testing.T) {
	transform1 := NewTransform(math2D.NewVector2D(1, 1), math2D.ZeroVector2D(), 30)
	transform2 := NewTransform(math2D.NewVector2D(2, 3), math2D.ZeroVector2D(), 15)

	transform1.Concat(transform2)

	expectedPosition := math2D.NewVector2D(3, 4)
	expectedRotation := 45

	if !transform1.GetPosition().IsEqual(expectedPosition) {
		t.Errorf("Expected position %v, got %v", expectedPosition, transform1.GetPosition())
	}
	if transform1.GetRotation() != expectedRotation {
		t.Errorf("Expected rotation %d, got %d", expectedRotation, transform1.GetRotation())
	}
}