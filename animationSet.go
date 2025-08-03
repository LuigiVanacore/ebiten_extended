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

func (a *AnimationSet) GetTexture() *ebiten.Image {
	if len(a.spriteSheet) == 0  {
		return nil
	}
	return a.spriteSheet[a.currentFrame]
}

func (a *AnimationSet) GetPivot() math2D.Vector2D {
	return a.pivot
}

func (a *AnimationSet) SetPivot(pivot math2D.Vector2D) {
	a.pivot = pivot
}

func (a *AnimationSet) GetFrameCount() uint {
	return a.frameCount
}

func (a *AnimationSet) SetFrameCount(frameCount uint) {
	a.frameCount = frameCount
	if a.currentFrame >= a.frameCount {
		a.currentFrame = 0
	}
}

func (a *AnimationSet) GetCurrentFrame() uint {
	return a.currentFrame
}

func (a *AnimationSet) SetCurrentFrame(frame uint) {
	if frame < a.frameCount {
		a.currentFrame = frame
	} else {
		a.currentFrame = 0
	}
}

func (a *AnimationSet) GetTimePerFrame() float64 {
	return a.timePerFrame
}

func (a *AnimationSet) SetTimePerFrame(timePerFrame float64) {
	if timePerFrame < 0 {
		timePerFrame = 1
	}
	
	a.timePerFrame = timePerFrame
	
}

func (a *AnimationSet) SetLooped(isLooped bool) {
	a.isLooped = isLooped
}

func (a *AnimationSet) IsLooped() bool {
	return a.isLooped
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
	return a.currentFrame >= a.frameCount 
}
