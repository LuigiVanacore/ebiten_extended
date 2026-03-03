package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

func TestShapeAABB_Circle(t *testing.T) {
	c := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 15))
	c.UpdateTransform(transformAt(100, 50))

	minX, minY, maxX, maxY := ShapeAABB(c)

	if minX != 85 || minY != 35 || maxX != 115 || maxY != 65 {
		t.Errorf("Circle AABB at (100,50) r=15: got (%.0f,%.0f,%.0f,%.0f), want (85,35,115,65)",
			minX, minY, maxX, maxY)
	}
}

func TestShapeAABB_Rect(t *testing.T) {
	r := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(40, 30)))
	r.UpdateTransform(transformAt(100, 50)) // center -> rect top-left (80,35)

	minX, minY, maxX, maxY := ShapeAABB(r)

	if minX != 80 || minY != 35 || maxX != 120 || maxY != 65 {
		t.Errorf("Rect AABB center (100,50) size (40,30): got (%.0f,%.0f,%.0f,%.0f), want (80,35,120,65)",
			minX, minY, maxX, maxY)
	}
}
