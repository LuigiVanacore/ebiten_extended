package collision

import (
	"github.com/LuigiVanacore/ludum"
)

type Collidable interface {
	ludum.Tagable
	ludum.Updatable
	IsCollide(other Collidable) bool
	GetShape() CollisionShape
}
