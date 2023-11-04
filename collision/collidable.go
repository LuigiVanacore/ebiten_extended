package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

type Collidable interface {
	ebiten_extended.Tagable
	IsCollide(other Collidable) bool
	GetShape() math2D.Shape
	ToString() string
}
