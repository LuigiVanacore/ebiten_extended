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

func TestDotProduct(t *testing.T) {
    v1 := NewVector2D(3, 4)
    v2 := NewVector2D(2, 0)
    dp := DotProduct(v1, v2)
    if dp != 6 {
        t.Errorf("DotProduct((3,4),(2,0)) = %v, want 6", dp)
    }
}

func TestMultiplyScalar(t *testing.T) {
    v := NewVector2D(2, -3)
    scaled := v.MultiplyScalar(4)
    if scaled.x != 8 || scaled.y != -12 {
        t.Errorf("MultiplyScalar(4) = (%v,%v), want (8,-12)", scaled.x, scaled.y)
    }
}

func TestClone(t *testing.T) {
    v := NewVector2D(7, 11)
    c := v.Clone()
    if c.x != v.x || c.y != v.y {
        t.Errorf("Clone = (%v,%v), want (%v,%v)", c.x, c.y, v.x, v.y)
    }
}

func TestLength(t *testing.T) {
    v := NewVector2D(3, 4)
    if v.Length() != 5 {
        t.Errorf("Length((3,4)) = %v, want 5", v.Length())
    }
}

func TestProjectVector(t *testing.T) {
    v := NewVector2D(4, 0)
    onto := NewVector2D(1, 0)
    proj := v.ProjectVector(onto)
    if proj.x != 4 || proj.y != 0 {
        t.Errorf("ProjectVector((4,0) onto (1,0)) = (%v,%v), want (4,0)", proj.x, proj.y)
    }
}

func TestDivideScalar(t *testing.T) {
    v := NewVector2D(10, 6)
    d := v.DivideScalar(2)
    if d.x != 5 || d.y != 3 {
        t.Errorf("DivideScalar(2) = (%v,%v), want (5,3)", d.x, d.y)
    }
}

func TestNormalize(t *testing.T) {
    v := NewVector2D(3, 4)
    n := v.Normalize()
    if math.Abs(n.Length()-1) > 1e-9 {
        t.Errorf("Normalized length = %v, want 1", n.Length())
    }
}

func TestNormalize_ZeroVector(t *testing.T) {
    v := ZeroVector2D()
    n := v.Normalize()
    if n.x != 0 || n.y != 0 {
        t.Errorf("Normalize of zero vector should return zero, got (%v,%v)", n.x, n.y)
    }
}

func TestRotateVector90(t *testing.T) {
    v := NewVector2D(1, 0)
    r := v.RotateVector90()
    if r.x != 0 || r.y != 1 {
        t.Errorf("RotateVector90(1,0) = (%v,%v), want (0,1)", r.x, r.y)
    }
}

func TestRotateVector180(t *testing.T) {
    v := NewVector2D(1, 2)
    r := v.RotateVector180()
    if r.x != -1 || r.y != -2 {
        t.Errorf("RotateVector180(1,2) = (%v,%v), want (-1,-2)", r.x, r.y)
    }
}

func TestRotateVector270(t *testing.T) {
    v := NewVector2D(1, 0)
    r := v.RotateVector270()
    if r.x != 0 || r.y != -1 {
        t.Errorf("RotateVector270(1,0) = (%v,%v), want (0,-1)", r.x, r.y)
    }
}

func TestSetPosition(t *testing.T) {
    v := NewVector2D(0, 0)
    v.SetPosition(NewVector2D(7, 8))
    if v.x != 7 || v.y != 8 {
        t.Errorf("SetPosition: got (%v,%v), want (7,8)", v.x, v.y)
    }
}

func TestTranslate(t *testing.T) {
    v := NewVector2D(1, 2)
    v.Translate(3, 4)
    if v.x != 4 || v.y != 6 {
        t.Errorf("Translate(3,4): got (%v,%v), want (4,6)", v.x, v.y)
    }
}

func TestSetToZero(t *testing.T) {
    v := NewVector2D(5, 5)
    v.SetToZero()
    if v.x != 0 || v.y != 0 {
        t.Errorf("SetToZero: got (%v,%v)", v.x, v.y)
    }
}

func TestIsEqual_False(t *testing.T) {
    v1 := NewVector2D(1, 2)
    v2 := NewVector2D(1, 3)
    if v1.IsEqual(v2) {
        t.Error("IsEqual((1,2),(1,3)) should be false")
    }
}

func TestDivideVectors(t *testing.T) {
    v1 := NewVector2D(10, 8)
    v2 := NewVector2D(2, 4)
    d := DivideVectors(v1, v2)
    if d.x != 5 || d.y != 2 {
        t.Errorf("DivideVectors((10,8),(2,4)) = (%v,%v), want (5,2)", d.x, d.y)
    }
}

func TestString(t *testing.T) {
    v := NewVector2D(1.5, 2.5)
    s := v.String()
    if s == "" {
        t.Error("String() should not be empty")
    }
}