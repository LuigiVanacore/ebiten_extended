package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// CollisionShape is implemented by CollisionCircle, CollisionRect, etc.
// UpdateTransform sets the shape in world space; IsColliding tests overlap with another shape.
type CollisionShape interface {
	IsColliding(other CollisionShape) bool
	UpdateTransform(transform transform.Transform)
}