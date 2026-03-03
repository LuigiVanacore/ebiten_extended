package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

// Drawable formalizes an interface for objects possessing both spatial awareness (transform) and the capacity to render themselves.
// GetLayer returns the z-order for sibling drawables within the same layer (lower = drawn first/behind).
type Drawable interface {
	transform.Transformable
	GetLayer() int
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}