package ebiten_extended

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)




type AnimationSet struct {
	transform Transform
	spriteSheet *ebiten.Image
	frameImage  *ebiten.Image
	frameSize Vector2D
	currentFrame uint
	frameCount uint
	timePerFrame float64
	elapsedTime float64
	isLooping bool
}

func NewAnimationSet(spriteSheet *ebiten.Image, frameCount uint, duration float64, isLooping bool) *AnimationSet {
	timePerFrame :=  1 / duration
	frameSize := Vector2D{ X: float64(spriteSheet.Bounds().Dx()) / float64(frameCount), Y: float64(spriteSheet.Bounds().Dy())}
	animationSet:=  &AnimationSet{ spriteSheet:  spriteSheet, frameSize: frameSize, frameCount: frameCount, timePerFrame: timePerFrame, isLooping: isLooping}
	animationSet.updateFrameImage()
	animationSet.SetPivotToCenter()
	animationSet.transform.Move(40,40)
	return animationSet
}

func (a *AnimationSet) GetTransform() *Transform {
	return &a.transform
}

func (a *AnimationSet) SetTransform(transform Transform) {
	a.transform = transform
}

func (a *AnimationSet)  SetPivotToCenter() {
	x, y := a.frameSize.X / 2, a.frameSize.Y / 2
	a.transform.SetPivot(x, y)
}

func (a *AnimationSet) updateFrameImage() {
	sx := int(a.currentFrame)*int(a.frameSize.X)
	a.frameImage =  a.spriteSheet.SubImage(image.Rect(sx,0,sx + int(a.frameSize.X), int(a.frameSize.Y))).(*ebiten.Image)
}

func (a *AnimationSet) Update(dt float64) {

	if a.IsEnded() {
		return
	}

	a.elapsedTime += dt

	if a.elapsedTime > a.timePerFrame {
		a.currentFrame++
		a.updateFrameImage()
		if a.IsEnded() && a.isLooping {
			a.currentFrame = 0
		}
		a.elapsedTime = 0
	}

}

func (a *AnimationSet) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM = a.transform.UpdateGeoM(op.GeoM)
	target.DrawImage(a.frameImage, op)
}

func (a *AnimationSet) IsEnded() bool {
	return a.currentFrame + 1 >= a.frameCount
}
