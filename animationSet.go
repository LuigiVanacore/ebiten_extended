package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)




// AnimationSet represents a singular, self-contained repeating or one-shot frame sequence.
type AnimationSet struct {
	spriteSheet []*ebiten.Image
	pivot math2D.Vector2D
	currentFrame uint
	frameCount uint
	timePerFrame float64
	elapsedTime float64
	isLooped bool
}

// NewAnimationSet establishes a playback sequence containing slice references, a duration constraint, and loop flag.
func NewAnimationSet(spriteSheet []*ebiten.Image, pivot math2D.Vector2D, frameCount uint, duration float64, isLooped bool) *AnimationSet {
	timePerFrame :=  1 / duration
	if ( duration == 0) {
		timePerFrame = 0
	}
	animationSet:=  &AnimationSet{ spriteSheet:  spriteSheet, pivot: pivot, frameCount: frameCount, timePerFrame: timePerFrame, isLooped: isLooped}
	return animationSet
}

// Update steps forward the animation sequence's relative frame pointer respecting its defined time signature boundaries.
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

// Draw displays the presently active chronological frame texture bounded within the assigned configuration sequence onto the target.
func (a *AnimationSet) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if a.frameCount == 0 || len(a.spriteSheet) == 0 {
		return
	}
	frame := a.currentFrame
	if frame >= a.frameCount {
		frame = a.frameCount - 1
	}
	if frame >= uint(len(a.spriteSheet)) {
		return
	}
	a.updateGeomToPivot(op)
	target.DrawImage(a.spriteSheet[frame], op)
}

// IsEnded evaluates whether a non-looping animation set profile has run completely through to its specified finish.
func (a *AnimationSet) IsEnded() bool {
	return a.currentFrame + 1 >= a.frameCount 
}
