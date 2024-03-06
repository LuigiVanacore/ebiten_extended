package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
)

type Collidable interface {
	ebiten_extended.Tagable
	ebiten_extended.Updatable
	IsCollide(other Collidable) bool
	GetShape() CollisionShape
}
