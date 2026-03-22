package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// CollisionOrientedRect is a rotated rectangle (center, half-extents, rotation).
// Use for platforms, sloped surfaces, or any axis-aligned rect that needs rotation.
type CollisionOrientedRect struct {
	rectangle math2D.OrientedRectangle
}

// NewCollisionOrientedRect creates a CollisionOrientedRect from a math2D.OrientedRectangle.
// halfExtended is half-width and half-height; center and rotation come from the node's transform.
func NewCollisionOrientedRect(rect math2D.OrientedRectangle) *CollisionOrientedRect {
	return &CollisionOrientedRect{rectangle: rect}
}

// NewCollisionOrientedRectFromSize creates an oriented rect with the given half-width and half-height.
// Center and rotation are applied from the transform when colliding.
func NewCollisionOrientedRectFromSize(halfWidth, halfHeight float64) *CollisionOrientedRect {
	return &CollisionOrientedRect{
		rectangle: math2D.NewOrientedRectangle(
			math2D.ZeroVector2D(),
			math2D.NewVector2D(halfWidth, halfHeight),
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
func (c *CollisionOrientedRect) GetOrientedRectangle() math2D.OrientedRectangle {
	return c.rectangle
}
