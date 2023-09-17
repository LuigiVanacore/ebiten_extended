package ebiten_extended

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Transformable
	Draw(target *ebiten.Image, op *ebiten.DrawImageOptions)
}