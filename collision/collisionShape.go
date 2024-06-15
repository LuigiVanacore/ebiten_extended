package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionShape interface {
	IsColliding(collisionShape CollisionShape) bool
	UpdateTransform(transform transform.Transform)
}