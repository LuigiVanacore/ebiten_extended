package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

// Drawable formalizes an interface for objects possessing both spatial awareness (transform) and the capacity to render themselves.
type Drawable interface {
	transform.Transformable
	// Draw applies visual properties onto the supplied frame surface according to inherent or inherited attributes.
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}