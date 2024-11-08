package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)




type AnimationSet struct {
	spriteSheet []*ebiten.Image
	pivot math2D.Vector2D
	currentFrame uint
	frameCount uint
	timePerFrame float64
	elapsedTime float64
	isLooped bool
}

func NewAnimationSet(spriteSheet []*ebiten.Image, pivot math2D.Vector2D, frameCount uint, duration float64, isLooped bool) *AnimationSet {
	timePerFrame :=  1 / duration
	if ( duration == 0) {
		timePerFrame = 0
	}
	animationSet:=  &AnimationSet{ spriteSheet:  spriteSheet, pivot: pivot, frameCount: frameCount, timePerFrame: timePerFrame, isLooped: isLooped}
	return animationSet
}

func (a *AnimationSet) Update() {

	if a.IsEnded() && !a.isLooped {
		return
	}

	a.elapsedTime += FIXED_DELTA

	if a.elapsedTime >= a.timePerFrame {
		a.currentFrame++
		if a.IsEnded() && a.isLooped {
			a.currentFrame = 0
		}
		a.elapsedTime = 0
	}

}
 
func (a *AnimationSet) updateGeomToPivot(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(-a.pivot.X(), -a.pivot.Y())
}

func (a *AnimationSet) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	a.updateGeomToPivot(op)
	target.DrawImage(a.spriteSheet[a.currentFrame], op)
}

func (a *AnimationSet) IsEnded() bool {
	return a.currentFrame + 1 >= a.frameCount 
}
