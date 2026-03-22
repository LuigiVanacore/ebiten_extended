package collision

import (
	"testing"

	"github.com/LuigiVanacore/ludum/math2d"
)

func TestShapeAABB_Circle(t *testing.T) {
	c := NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 15))
	tC := transformAt(100, 50)

	minX, minY, maxX, maxY := ShapeAABB(c, tC)

	if minX != 85 || minY != 35 || maxX != 115 || maxY != 65 {
		t.Errorf("Circle AABB at (100,50) r=15: got (%.0f,%.0f,%.0f,%.0f), want (85,35,115,65)",
			minX, minY, maxX, maxY)
	}
}

func TestShapeAABB_Rect(t *testing.T) {
	r := NewCollisionRect(math2d.NewRectangle(math2d.ZeroVector2D(), math2d.NewVector2D(40, 30)))
	tR := transformAt(100, 50) // center -> rect top-left (80,35)

	minX, minY, maxX, maxY := ShapeAABB(r, tR)

	if minX != 80 || minY != 35 || maxX != 120 || maxY != 65 {
		t.Errorf("Rect AABB center (100,50) size (40,30): got (%.0f,%.0f,%.0f,%.0f), want (80,35,120,65)",
			minX, minY, maxX, maxY)
	}
}
