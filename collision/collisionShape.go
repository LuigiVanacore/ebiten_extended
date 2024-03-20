package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
)

type CollisionShape interface {
	ebiten_extended.Updatable
	IsColliding(collisionShape CollisionShape) bool
}