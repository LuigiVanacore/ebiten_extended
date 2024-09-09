package math2D

import (
    "math"
    "testing"
)

func TestNewVector2D(t *testing.T) {
    v := NewVector2D(3, 4)
    if v.x != 3 || v.y != 4 {
        t.Errorf("Expected (3, 4), got (%v, %v)", v.x, v.y)
    }
}

func TestZeroVector2D(t *testing.T) {
    v := ZeroVector2D()
    if v.x != 0 || v.y != 0 {
        t.Errorf("Expected (0, 0), got (%v, %v)", v.x, v.y)
    }
}

func TestOneVector2D(t *testing.T) {
    v := OneVector2D()
    if v.x != 1 || v.y != 1 {
        t.Errorf("Expected (1, 1), got (%v, %v)", v.x, v.y)
    }
}

func TestMagnitude(t *testing.T) {
    v := NewVector2D(3, 4)
    expected := 5.0
    if v.Magnitude() != expected {
        t.Errorf("Expected %v, got %v", expected, v.Magnitude())
    }
}

func TestAddVectors(t *testing.T) {
    v1 := NewVector2D(1, 2)
    v2 := NewVector2D(3, 4)
    v := AddVectors(v1, v2)
    if v.x != 4 || v.y != 6 {
        t.Errorf("Expected (4, 6), got (%v, %v)", v.x, v.y)
    }
}

func TestSubtractVectors(t *testing.T) {
    v1 := NewVector2D(5, 6)
    v2 := NewVector2D(3, 4)
    v := SubtractVectors(v1, v2)
    if v.x != 2 || v.y != 2 {
        t.Errorf("Expected (2, 2), got (%v, %v)", v.x, v.y)
    }
}

func TestRotateVector(t *testing.T) {
    v := NewVector2D(1, 0)
    rotated := v.RotateVector(90)
    if math.Abs(rotated.x) > 1e-9 || math.Abs(rotated.y-1) > 1e-9 {
        t.Errorf("Expected (0, 1), got (%v, %v)", rotated.x, rotated.y)
    }
}

func TestIsZero(t *testing.T) {
    v := ZeroVector2D()
    if !v.IsZero() {
        t.Errorf("Expected true, got false")
    }
}

func TestNegate(t *testing.T) {
    v := NewVector2D(1, -2)
    negated := v.Negate()
    if negated.x != -1 || negated.y != 2 {
        t.Errorf("Expected (-1, 2), got (%v, %v)", negated.x, negated.y)
    }
}

func TestIsEqual(t *testing.T) {
    v1 := NewVector2D(1, 2)
    v2 := NewVector2D(1, 2)
    if !v1.IsEqual(v2) {
        t.Errorf("Expected true, got false")
    }
}