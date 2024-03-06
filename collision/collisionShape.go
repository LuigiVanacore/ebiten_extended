package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionShape interface {
	ebiten_extended.Updatable
	IsColliding(collisionShape CollisionShape) bool
	SetTransform(transfom *transform.Transform)
}