package math2d

import "testing"

func TestNewRange(t *testing.T) {
	r := NewRange(1, 10)
	if r.minimum != 1 || r.maximum != 10 {
		t.Errorf("NewRange(1,10): got (%v,%v)", r.minimum, r.maximum)
	}
}

func TestRange_GetMinimum_GetMaximum_SetMinimum_SetMaximum(t *testing.T) {
	r := NewRange(0, 5)
	if r.GetMinimum() != 0 || r.GetMaximum() != 5 {
		t.Errorf("GetMinimum/Maximum: got (%v,%v)", r.GetMinimum(), r.GetMaximum())
	}
	r.SetMinimum(2)
	r.SetMaximum(8)
	if r.GetMinimum() != 2 || r.GetMaximum() != 8 {
		t.Errorf("SetMinimum/Maximum: got (%v,%v)", r.GetMinimum(), r.GetMaximum())
	}
}

func TestOverlappingRanges(t *testing.T) {
	if !OverlappingRanges(NewRange(0, 10), NewRange(5, 15)) {
		t.Error("OverlappingRanges(0,10) and (5,15) should overlap")
	}
	if OverlappingRanges(NewRange(0, 5), NewRange(10, 20)) {
		t.Error("OverlappingRanges(0,5) and (10,20) should not overlap")
	}
	if !OverlappingRanges(NewRange(0, 10), NewRange(10, 20)) {
		t.Error("OverlappingRanges touching at 10 should overlap")
	}
}

func TestRangeHull(t *testing.T) {
	a := NewRange(0, 5)
	b := NewRange(3, 8)
	hull := RangeHull(a, b)
	if hull.GetMinimum() != 0 || hull.GetMaximum() != 8 {
		t.Errorf("RangeHull: got (%v,%v), want (0,8)", hull.GetMinimum(), hull.GetMaximum())
	}
}

func TestRange_SortRange(t *testing.T) {
	r := NewRange(10, 0)
	sorted := r.SortRange()
	if sorted.GetMinimum() != 0 || sorted.GetMaximum() != 10 {
		t.Errorf("SortRange(10,0): got (%v,%v), want (0,10)", sorted.GetMinimum(), sorted.GetMaximum())
	}
	r2 := NewRange(0, 5)
	sorted2 := r2.SortRange()
	if sorted2.GetMinimum() != 0 || sorted2.GetMaximum() != 5 {
		t.Errorf("SortRange(0,5) unchanged: got (%v,%v)", sorted2.GetMinimum(), sorted2.GetMaximum())
	}
}
