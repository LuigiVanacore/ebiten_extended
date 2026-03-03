package tween

import (
	"math"
	"testing"
)

func floatEqualLoose(a, b float32) bool {
	return float32(math.Abs(float64(a-b))) <= 1e-1
}

func TestEasingFunctions_EndAtBPlusC(t *testing.T) {
	easings := []struct {
		name  string
		fn    TweenFunc
		needs []float32 // time points to test (0, mid, d)
	}{
		{"Linear", Linear, []float32{0, 0.5, 1}},
		{"InQuad", InQuad, []float32{0, 0.5, 1}},
		{"OutQuad", OutQuad, []float32{0, 0.5, 1}},
		{"InOutQuad", InOutQuad, []float32{0, 0.3, 0.7, 1}},
		{"OutInQuad", OutInQuad, []float32{0, 0.25, 0.75, 1}},
		{"InCubic", InCubic, []float32{0, 0.5, 1}},
		{"OutCubic", OutCubic, []float32{0, 0.5, 1}},
		{"InOutCubic", InOutCubic, []float32{0, 0.5, 1}},
		{"OutInCubic", OutInCubic, []float32{0, 0.25, 0.75, 1}},
		{"InQuart", InQuart, []float32{0, 0.5, 1}},
		{"OutQuart", OutQuart, []float32{0, 0.5, 1}},
		{"InOutQuart", InOutQuart, []float32{0, 0.5, 1}},
		{"OutInQuart", OutInQuart, []float32{0, 0.25, 0.75, 1}},
		{"InQuint", InQuint, []float32{0, 0.5, 1}},
		{"OutQuint", OutQuint, []float32{0, 0.5, 1}},
		{"InOutQuint", InOutQuint, []float32{0, 0.5, 1}},
		{"OutInQuint", OutInQuint, []float32{0, 0.25, 0.75, 1}},
		{"InSine", InSine, []float32{0, 0.5, 1}},
		{"OutSine", OutSine, []float32{0, 0.5, 1}},
		{"InOutSine", InOutSine, []float32{0, 0.5, 1}},
		{"OutInSine", OutInSine, []float32{0, 0.25, 0.75, 1}},
		{"InExpo", InExpo, []float32{0, 0.5, 1}},
		{"OutExpo", OutExpo, []float32{0, 0.5, 1}},
		{"InOutExpo", InOutExpo, []float32{0, 0.5, 1}},
		{"OutInExpo", OutInExpo, []float32{0, 0.25, 0.75, 1}},
		{"InCirc", InCirc, []float32{0, 0.5, 1}},
		{"OutCirc", OutCirc, []float32{0, 0.5, 1}},
		{"InOutCirc", InOutCirc, []float32{0, 0.5, 1}},
		{"OutInCirc", OutInCirc, []float32{0, 0.25, 0.75, 1}},
		{"InElastic", InElastic, []float32{0, 0.3, 0.7, 1}},
		{"OutElastic", OutElastic, []float32{0, 0.3, 0.7, 1}},
		{"InOutElastic", InOutElastic, []float32{0, 0.3, 0.7, 1.5, 2}},
		{"OutInElastic", OutInElastic, []float32{0, 0.25, 0.75, 1}},
		{"InBack", InBack, []float32{0, 0.5, 1}},
		{"OutBack", OutBack, []float32{0, 0.5, 1}},
		{"InOutBack", InOutBack, []float32{0, 0.5, 1}},
		{"OutInBack", OutInBack, []float32{0, 0.25, 0.75, 1}},
		{"OutBounce", OutBounce, []float32{0, 0.2, 0.5, 0.8, 1}},
		{"InBounce", InBounce, []float32{0, 0.5, 1}},
		{"InOutBounce", InOutBounce, []float32{0, 0.25, 0.75, 1}},
		{"OutInBounce", OutInBounce, []float32{0, 0.25, 0.75, 1}},
	}

	b, c, d := float32(10), float32(90), float32(1.0)
	for _, e := range easings {
		t.Run(e.name, func(t *testing.T) {
			for _, tt := range e.needs {
				v := e.fn(tt, b, c, d)
				if tt == 0 && !floatEqual(v, b) {
					t.Errorf("%s(0): got %v, want %v", e.name, v, b)
				}
				if tt >= d {
					if !floatEqual(v, b+c) && !floatEqualLoose(v, b+c) {
						t.Errorf("%s(%v): got %v, want ~%v", e.name, tt, v, b+c)
					}
				}
			}
		})
	}
}

func TestOutBounce_AllBranches(t *testing.T) {
	// OutBounce has 4 branches based on t/d
	d := float32(1.0)
	tests := []float32{0.1, 0.3, 0.6, 0.9, 1.0}
	for _, tt := range tests {
		v := OutBounce(tt, 0, 100, d)
		if tt == 0 && v != 0 {
			t.Errorf("OutBounce(0): got %v", v)
		}
		if tt >= d && !floatEqual(v, float32(100)) {
			t.Errorf("OutBounce(1): got %v", v)
		}
	}
}

func TestInOutElastic_Branches(t *testing.T) {
	// t/d*2: 0, 1, 2 branches; t-1: t<0 and t>=0
	v0 := InOutElastic(0, 0, 100, 1.0)
	if v0 != 0 {
		t.Errorf("InOutElastic(0): got %v", v0)
	}
	v2 := InOutElastic(1.0, 0, 100, 1.0)
	if !floatEqual(v2, 100) {
		t.Errorf("InOutElastic(1): got %v", v2)
	}
	// Elastic can overshoot/undershoot beyond [0,100]
	vMid := InOutElastic(0.3, 0, 100, 1.0)
	if vMid < -200 || vMid > 200 {
		t.Errorf("InOutElastic(0.3): got %v, expect reasonable value", vMid)
	}
}

func TestInOutQuad_BothBranches(t *testing.T) {
	v1 := InOutQuad(0.3, 0, 100, 1.0) // t < 1 branch
	if v1 < 0 || v1 > 100 {
		t.Errorf("InOutQuad first branch: got %v", v1)
	}
	v2 := InOutQuad(0.8, 0, 100, 1.0) // t >= 1 branch
	if v2 < 0 || v2 > 100 {
		t.Errorf("InOutQuad second branch: got %v", v2)
	}
}

func TestOutInQuad_BothBranches(t *testing.T) {
	v1 := OutInQuad(0.3, 0, 100, 1.0) // t < d/2
	if v1 < 0 || v1 > 100 {
		t.Errorf("OutInQuad first branch: got %v", v1)
	}
	v2 := OutInQuad(0.8, 0, 100, 1.0) // t >= d/2
	if v2 < 0 || v2 > 100 {
		t.Errorf("OutInQuad second branch: got %v", v2)
	}
}

func TestInOutExpo_BothMidBranches(t *testing.T) {
	v1 := InOutExpo(0.3, 0, 100, 1.0) // t < 1
	if v1 < 0 || v1 > 100 {
		t.Errorf("InOutExpo first mid: got %v", v1)
	}
	v2 := InOutExpo(0.8, 0, 100, 1.0) // t >= 1
	if v2 < 0 || v2 > 100 {
		t.Errorf("InOutExpo second mid: got %v", v2)
	}
}

func TestInElastic_AllBranches(t *testing.T) {
	v0 := InElastic(0, 0, 100, 1.0)
	if v0 != 0 {
		t.Errorf("InElastic(0): got %v", v0)
	}
	v1 := InElastic(1.0, 0, 100, 1.0)
	if !floatEqual(v1, 100) {
		t.Errorf("InElastic(1): got %v", v1)
	}
	vMid := InElastic(0.5, 0, 100, 1.0)
	if vMid < -50 || vMid > 150 {
		t.Errorf("InElastic(0.5) can overshoot: got %v", vMid)
	}
}

func TestOutExpo_EdgeCases(t *testing.T) {
	v0 := OutExpo(0, 0, 100, 1.0)
	if v0 != 0 {
		t.Errorf("OutExpo(0): got %v", v0)
	}
	v1 := OutExpo(1.0, 0, 100, 1.0)
	if v1 != 100 {
		t.Errorf("OutExpo(1): got %v", v1)
	}
}
