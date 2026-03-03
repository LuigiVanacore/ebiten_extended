package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

func transformAt(x, y float64) transform.Transform {
	return transform.NewTransform(math2D.NewVector2D(x, y), math2D.ZeroVector2D(), 0)
}

func TestShapeCollides_CircleCircle(t *testing.T) {
	c1 := NewCollisionCircle(math2D.NewCircle(math2D.NewVector2D(0, 0), 10))
	c2 := NewCollisionCircle(math2D.NewCircle(math2D.NewVector2D(0, 0), 10))
	c1.UpdateTransform(transformAt(0, 0))
	c2.UpdateTransform(transformAt(15, 0)) // distance 15, sum radii 20 -> overlap

	if !ShapeCollides(c1, c2) {
		t.Error("circles at distance 15 (r=10 each, sum=20) should collide")
	}

	c2.UpdateTransform(transformAt(25, 0)) // distance 25 > 20 -> no overlap
	if ShapeCollides(c1, c2) {
		t.Error("circles at distance 25 (r=10 each) should not collide")
	}
}

func TestShapeCollides_CircleRect(t *testing.T) {
	c := NewCollisionCircle(math2D.NewCircle(math2D.NewVector2D(0, 0), 10))
	r := &CollisionRect{rectangle: math2D.NewRectangle(math2D.NewVector2D(50, 0), math2D.NewVector2D(20, 20))}
	c.UpdateTransform(transformAt(0, 0))
	r.UpdateTransform(transformAt(50, 0))

	if ShapeCollides(c, r) {
		t.Error("circle and rect far apart should not collide")
	}

	r.UpdateTransform(transformAt(5, 0)) // rect now overlaps with circle
	if !ShapeCollides(c, r) {
		t.Error("overlapping circle and rect should collide")
	}
}

func TestShapeCollides_RectRect(t *testing.T) {
	r1 := &CollisionRect{rectangle: math2D.NewRectangle(math2D.NewVector2D(0, 0), math2D.NewVector2D(10, 10))}
	r2 := &CollisionRect{rectangle: math2D.NewRectangle(math2D.NewVector2D(50, 0), math2D.NewVector2D(10, 10))}
	r1.UpdateTransform(transformAt(0, 0))
	r2.UpdateTransform(transformAt(50, 0))

	if ShapeCollides(r1, r2) {
		t.Error("rects far apart should not collide")
	}

	r2.UpdateTransform(transformAt(5, 5))
	if !ShapeCollides(r1, r2) {
		t.Error("overlapping rects should collide")
	}
}

func TestShapeCollides_RectCircle(t *testing.T) {
	r := &CollisionRect{rectangle: math2D.NewRectangle(math2D.NewVector2D(0, 0), math2D.NewVector2D(20, 20))}
	c := NewCollisionCircle(math2D.NewCircle(math2D.NewVector2D(0, 0), 5))
	r.UpdateTransform(transformAt(0, 0))
	c.UpdateTransform(transformAt(10, 10))

	if !ShapeCollides(r, c) {
		t.Error("circle inside rect should collide")
	}
}

