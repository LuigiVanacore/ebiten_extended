package ludum

import (
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

// Drawable formalizes an interface for objects possessing both spatial awareness (transform) and the capacity to render themselves.
// GetLayer returns the z-order for sibling drawables within the same layer (lower = drawn first/behind).
type Drawable interface {
	transform.Transformable
	GetLayer() int
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}

// Cullable is an optional interface for Drawables that report their world-space AABB.
// When implemented, the World may skip drawing if the bounds are outside the camera view.
type Cullable interface {
	GetWorldBounds() math2d.Rectangle
}
