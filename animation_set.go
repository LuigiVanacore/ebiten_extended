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

// NewAnimationSet creates a playback sequence from a sprite sheet slice.
// fps is the playback speed in frames per second; pass 0 to pause animation.
func NewAnimationSet(spriteSheet []*ebiten.Image, pivot math2D.Vector2D, frameCount uint, fps float64, isLooped bool) *AnimationSet {
	var timePerFrame float64
	if fps > 0 {
		timePerFrame = 1 / fps
	}
	return &AnimationSet{spriteSheet: spriteSheet, pivot: pivot, frameCount: frameCount, timePerFrame: timePerFrame, isLooped: isLooped}
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



// Update advances the animation by one fixed frame. Implements Updatable.
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
	return a.currentFrame >= a.frameCount 
}
