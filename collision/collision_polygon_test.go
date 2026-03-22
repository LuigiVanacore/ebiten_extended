package collision

import (
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
	"github.com/LuigiVanacore/ludum/utils"
)

func TestCollisionPolygon_Triangle(t *testing.T) {
	verts := []math2d.Vector2D{
		math2d.NewVector2D(0, 10),
		math2d.NewVector2D(-10, -10),
		math2d.NewVector2D(10, -10),
	}
	poly := NewCollisionPolygon(verts)
	if poly == nil {
		t.Fatal("NewCollisionPolygon returned nil")
	}
	tA := transform.Transform{}
	tA.SetPosition(math2d.NewVector2D(50, 50))
	tB := transform.Transform{}
	tB.SetPosition(math2d.NewVector2D(55, 50))
	circ := NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 5))
	if !ShapeCollides(poly, tA, circ, tB) {
		t.Error("triangle and circle at center should collide")
	}
}

func TestCollisionPolygon_PolygonPolygon(t *testing.T) {
	v1 := []math2d.Vector2D{
		math2d.NewVector2D(-5, 5),
		math2d.NewVector2D(5, 5),
		math2d.NewVector2D(0, -5),
	}
	v2 := []math2d.Vector2D{
		math2d.NewVector2D(-3, 0),
		math2d.NewVector2D(3, 0),
		math2d.NewVector2D(0, -3),
	}
	p1 := NewCollisionPolygon(v1)
	p2 := NewCollisionPolygon(v2)
	tA := transform.Transform{}
	tA.SetPosition(math2d.NewVector2D(0, 0))
	tB := transform.Transform{}
	tB.SetPosition(math2d.NewVector2D(2, 0))
	if !ShapeCollides(p1, tA, p2, tB) {
		t.Error("overlapping triangles should collide")
	}
	tB.SetPosition(math2d.NewVector2D(50, 0))
	if ShapeCollides(p1, tA, p2, tB) {
		t.Error("separate triangles should not collide")
	}
}

func TestCollisionPolygon_OverlapPoint(t *testing.T) {
	verts := []math2d.Vector2D{
		math2d.NewVector2D(-10, 0),
		math2d.NewVector2D(0, 10),
		math2d.NewVector2D(10, 0),
		math2d.NewVector2D(0, -10),
	}
	poly := NewCollisionPolygon(verts)
	col, _ := NewCollider("p", poly, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	col.SetPosition(100, 100)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	results := mgr.OverlapPoint(math2d.NewVector2D(100, 100))
	if len(results) != 1 {
		t.Errorf("OverlapPoint at center: got %d", len(results))
	}
}
