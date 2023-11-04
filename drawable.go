package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	transform.Transformable
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}