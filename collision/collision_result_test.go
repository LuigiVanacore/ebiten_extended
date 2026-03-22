package collision

import (
	"math"
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
)

func TestCirclesCollideResult(t *testing.T) {
	tests := []struct {
		name      string
		a         math2d.Circle
		b         math2d.Circle
		wantOver  bool
		wantDepth float64
	}{
		{
			name:      "overlapping",
			a:         math2d.NewCircle(math2d.NewVector2D(0, 0), 10),
			b:         math2d.NewCircle(math2d.NewVector2D(15, 0), 10),
			wantOver:  true,
			wantDepth: 5,
		},
		{
			name:     "not overlapping",
			a:        math2d.NewCircle(math2d.NewVector2D(0, 0), 10),
			b:        math2d.NewCircle(math2d.NewVector2D(25, 0), 10),
			wantOver: false,
		},
		{
			name:     "tangent",
			a:        math2d.NewCircle(math2d.NewVector2D(0, 0), 10),
			b:        math2d.NewCircle(math2d.NewVector2D(20, 0), 10),
			wantOver: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CirclesCollideResult(tt.a, tt.b)
			if res.Overlapping != tt.wantOver {
				t.Errorf("Overlapping = %v, want %v", res.Overlapping, tt.wantOver)
			}
			if tt.wantOver && math.Abs(res.Depth-tt.wantDepth) > 0.001 {
				t.Errorf("Depth = %v, want %v", res.Depth, tt.wantDepth)
			}
			if tt.wantOver && res.Normal.Length() < 0.99 {
				t.Errorf("Normal should be unit length, got %v", res.Normal.Length())
			}
		})
	}
}

func TestCircleRectangleCollideResult(t *testing.T) {
	tests := []struct {
		name     string
		c        math2d.Circle
		r        math2d.Rectangle
		wantOver bool
	}{
		{
			name:     "overlapping",
			c:        math2d.NewCircle(math2d.NewVector2D(15, 15), 10),
			r:        math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(30, 30)),
			wantOver: true,
		},
		{
			name:     "not overlapping",
			c:        math2d.NewCircle(math2d.NewVector2D(50, 50), 10),
			r:        math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(20, 20)),
			wantOver: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := CircleRectangleCollideResult(tt.c, tt.r)
			if res.Overlapping != tt.wantOver {
				t.Errorf("Overlapping = %v, want %v", res.Overlapping, tt.wantOver)
			}
		})
	}
}

func TestRectanglesCollideResult(t *testing.T) {
	tests := []struct {
		name     string
		a        math2d.Rectangle
		b        math2d.Rectangle
		wantOver bool
	}{
		{
			name:     "overlapping",
			a:        math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(10, 10)),
			b:        math2d.NewRectangle(math2d.NewVector2D(5, 5), math2d.NewVector2D(10, 10)),
			wantOver: true,
		},
		{
			name:     "not overlapping",
			a:        math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(10, 10)),
			b:        math2d.NewRectangle(math2d.NewVector2D(20, 20), math2d.NewVector2D(10, 10)),
			wantOver: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := RectanglesCollideResult(tt.a, tt.b)
			if res.Overlapping != tt.wantOver {
				t.Errorf("Overlapping = %v, want %v", res.Overlapping, tt.wantOver)
			}
		})
	}
}

func TestShapeCollisionResult(t *testing.T) {
	c1 := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 10))
	c2 := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 10))
	t1 := transformAt(0, 0)
	t2 := transformAt(12, 0)

	res, ok := ShapeCollisionResult(c1, t1, c2, t2)
	if !ok {
		t.Fatal("ShapeCollisionResult should succeed for circle-circle")
	}
	if !res.Overlapping {
		t.Error("circles should overlap")
	}
	if res.Depth < 7 || res.Depth > 9 {
		t.Errorf("Depth ~8 expected, got %v", res.Depth)
	}
}

func TestShapeCollisionResult_CircleRect(t *testing.T) {
	c := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 10))
	r := NewCollisionRect(math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(30, 30)))
	tC := transformAt(15, 15)
	tR := transformAt(15, 15)

	res, ok := ShapeCollisionResult(c, tC, r, tR)
	if !ok {
		t.Fatal("ShapeCollisionResult should succeed for circle-rect")
	}
	if !res.Overlapping {
		t.Error("circle inside rect should overlap")
	}
	if res.Depth <= 0 {
		t.Errorf("Depth should be positive, got %v", res.Depth)
	}
}

func TestShapeCollisionResult_RectRect(t *testing.T) {
	r1 := NewCollisionRect(math2d.NewRectangle(math2d.ZeroVector2D(), math2d.NewVector2D(10, 10)))
	r2 := NewCollisionRect(math2d.NewRectangle(math2d.ZeroVector2D(), math2d.NewVector2D(10, 10)))
	t1 := transformAt(0, 0)
	t2 := transformAt(5, 5)

	res, ok := ShapeCollisionResult(r1, t1, r2, t2)
	if !ok {
		t.Fatal("ShapeCollisionResult should succeed for rect-rect")
	}
	if !res.Overlapping {
		t.Error("overlapping rects should overlap")
	}
	if res.Depth <= 0 {
		t.Errorf("Depth should be positive, got %v", res.Depth)
	}
}

func TestCirclesCollideResult_Coincident(t *testing.T) {
	a := math2d.NewCircle(math2d.NewVector2D(0, 0), 5)
	b := math2d.NewCircle(math2d.NewVector2D(0, 0), 5)

	res := CirclesCollideResult(a, b)
	if !res.Overlapping {
		t.Error("coincident circles should overlap")
	}
	if res.Depth != 10 {
		t.Errorf("Depth = %v, want 10", res.Depth)
	}
	if res.Normal.Length() < 0.99 {
		t.Errorf("fallback normal should be unit length, got %v", res.Normal.Length())
	}
}
