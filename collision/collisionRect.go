package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionRect struct {
	rectangle math2D.Rectangle
}

// NewCollisionRect creates a CollisionRect from a math2D.Rectangle.
func NewCollisionRect(rectangle math2D.Rectangle) *CollisionRect {
	return &CollisionRect{rectangle: rectangle}
}

func (c *CollisionRect) shapeKind() shapeKind {
	return kindRect
}

func (c *CollisionRect) IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool {
	return ShapeCollides(c, tSelf, other, tOther)
}
