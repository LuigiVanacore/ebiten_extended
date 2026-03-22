package collision

import (
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
)

func transformAt(x, y float64) transform.Transform {
	return transform.NewTransform(math2d.NewVector2D(x, y), math2d.ZeroVector2D(), 0)
}

func TestShapeCollides_CircleCircle(t *testing.T) {
	c1 := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 10))
	c2 := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 10))
	t1 := transformAt(0, 0)
	t2 := transformAt(15, 0) // distance 15, sum radii 20 -> overlap

	if !ShapeCollides(c1, t1, c2, t2) {
		t.Error("circles at distance 15 (r=10 each, sum=20) should collide")
	}

	t2 = transformAt(25, 0) // distance 25 > 20 -> no overlap
	if ShapeCollides(c1, t1, c2, t2) {
		t.Error("circles at distance 25 (r=10 each) should not collide")
	}
}

func TestShapeCollides_CircleRect(t *testing.T) {
	c := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 10))
	r := &CollisionRect{rectangle: math2d.NewRectangle(math2d.NewVector2D(50, 0), math2d.NewVector2D(20, 20))}
	tC := transformAt(0, 0)
	tR := transformAt(50, 0)

	if ShapeCollides(c, tC, r, tR) {
		t.Error("circle and rect far apart should not collide")
	}

	tR = transformAt(5, 0) // rect now overlaps with circle
	if !ShapeCollides(c, tC, r, tR) {
		t.Error("overlapping circle and rect should collide")
	}
}

func TestShapeCollides_RectRect(t *testing.T) {
	r1 := &CollisionRect{rectangle: math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(10, 10))}
	r2 := &CollisionRect{rectangle: math2d.NewRectangle(math2d.NewVector2D(50, 0), math2d.NewVector2D(10, 10))}
	t1 := transformAt(0, 0)
	t2 := transformAt(50, 0)

	if ShapeCollides(r1, t1, r2, t2) {
		t.Error("rects far apart should not collide")
	}

	t2 = transformAt(5, 5)
	if !ShapeCollides(r1, t1, r2, t2) {
		t.Error("overlapping rects should collide")
	}
}

func TestShapeCollides_RectCircle(t *testing.T) {
	r := &CollisionRect{rectangle: math2d.NewRectangle(math2d.NewVector2D(0, 0), math2d.NewVector2D(20, 20))}
	c := NewCollisionCircle(math2d.NewCircle(math2d.NewVector2D(0, 0), 5))
	tR := transformAt(0, 0)
	tC := transformAt(10, 10)

	if !ShapeCollides(r, tR, c, tC) {
		t.Error("circle inside rect should collide")
	}
}
