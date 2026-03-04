package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// CollisionShape is implemented by CollisionCircle, CollisionRect, etc.
// IsColliding tests overlap with another shape, given their respective world transforms,
// without mutating the underlying geometric data.
type CollisionShape interface {
	IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool
}
