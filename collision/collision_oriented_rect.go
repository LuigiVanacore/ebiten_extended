package collision

import (
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
)

// CollisionOrientedRect is a rotated rectangle (center, half-extents, rotation).
// Use for platforms, sloped surfaces, or any axis-aligned rect that needs rotation.
type CollisionOrientedRect struct {
	rectangle math2d.OrientedRectangle
}

// NewCollisionOrientedRect creates a CollisionOrientedRect from a math2d.OrientedRectangle.
// halfExtended is half-width and half-height; center and rotation come from the node's transform.
func NewCollisionOrientedRect(rect math2d.OrientedRectangle) *CollisionOrientedRect {
	return &CollisionOrientedRect{rectangle: rect}
}

// NewCollisionOrientedRectFromSize creates an oriented rect with the given half-width and half-height.
// Center and rotation are applied from the transform when colliding.
func NewCollisionOrientedRectFromSize(halfWidth, halfHeight float64) *CollisionOrientedRect {
	return &CollisionOrientedRect{
		rectangle: math2d.NewOrientedRectangle(
			math2d.ZeroVector2D(),
			math2d.NewVector2D(halfWidth, halfHeight),
			0,
		),
	}
}

func (c *CollisionOrientedRect) shapeKind() shapeKind {
	return kindOrientedRect
}

func (c *CollisionOrientedRect) IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool {
	return ShapeCollides(c, tSelf, other, tOther)
}

// GetOrientedRectangle returns the underlying oriented rectangle (for reads only).
func (c *CollisionOrientedRect) GetOrientedRectangle() math2d.OrientedRectangle {
	return c.rectangle
}
