package collision

import (
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
)

type CollisionRect struct {
	rectangle math2d.Rectangle
}

// NewCollisionRect creates a CollisionRect from a math2d.Rectangle.
func NewCollisionRect(rectangle math2d.Rectangle) *CollisionRect {
	return &CollisionRect{rectangle: rectangle}
}

func (c *CollisionRect) shapeKind() shapeKind {
	return kindRect
}

func (c *CollisionRect) IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool {
	return ShapeCollides(c, tSelf, other, tOther)
}
