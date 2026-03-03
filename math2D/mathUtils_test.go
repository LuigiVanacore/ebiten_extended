package math2D

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	if Min(3, 5) != 3 {
		t.Errorf("Min(3,5) = %v, want 3", Min(3, 5))
	}
	if Min(5, 3) != 3 {
		t.Errorf("Min(5,3) = %v, want 3", Min(5, 3))
	}
	if Min(0, 5) != 0 {
		t.Errorf("Min(0,5) = %v, want 0", Min(0, 5))
	}
	if Min(-1, 1) != -1 {
		t.Errorf("Min(-1,1) = %v, want -1", Min(-1, 1))
	}
}

func TestMax(t *testing.T) {
	if Max(3, 5) != 5 {
		t.Errorf("Max(3,5) = %v, want 5", Max(3, 5))
	}
	if Max(5, 3) != 5 {
		t.Errorf("Max(5,3) = %v, want 5", Max(5, 3))
	}
	if Max(-2, 0) != 0 {
		t.Errorf("Max(-2,0) = %v, want 0", Max(-2, 0))
	}
}

func TestDistance(t *testing.T) {
	d := Distance(0, 0, 3, 4)
	if math.Abs(d-5) > 1e-9 {
		t.Errorf("Distance(0,0,3,4) = %v, want 5", d)
	}
	d2 := Distance(1, 1, 1, 1)
	if d2 != 0 {
		t.Errorf("Distance(1,1,1,1) = %v, want 0", d2)
	}
}

func TestDegreesToRadian(t *testing.T) {
	rad := DegreesToRadian(180)
	if math.Abs(rad-math.Pi) > 1e-9 {
		t.Errorf("DegreesToRadian(180) = %v, want pi", rad)
	}
	rad90 := DegreesToRadian(90)
	if math.Abs(rad90-math.Pi/2) > 1e-9 {
		t.Errorf("DegreesToRadian(90) = %v, want pi/2", rad90)
	}
}

func TestRadianToDegrees(t *testing.T) {
	deg := RadianToDegrees(math.Pi)
	if math.Abs(deg-180) > 1e-9 {
		t.Errorf("RadianToDegrees(pi) = %v, want 180", deg)
	}
}
