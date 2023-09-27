package ebiten_extended

import (
 
	"github.com/hajimehoshi/ebiten/v2"
)


type AnimationPlayer struct {
	transform Transform
	animationMap map[uint]*AnimationSet
	currentAnimationId uint
	isPlaying bool
}

func NewAnimationPlayer() *AnimationPlayer {
	return &AnimationPlayer{ animationMap: make(map[uint]*AnimationSet)}
}

func (a *AnimationPlayer) GetTransform() *Transform {
	return &a.transform
}

func (a *AnimationPlayer) SetTransform(transform Transform) {
	a.transform = transform
}

func (a *AnimationPlayer) IsPlaying() bool {
	return a.isPlaying
}

func (a *AnimationPlayer) Start() {
	a.isPlaying = true
}

func (a *AnimationPlayer) Stop() {
	a.isPlaying = false
}

func (a *AnimationPlayer) GetCurrentAnimation() uint {
	return a.currentAnimationId
}

func (a *AnimationPlayer) SetCurrentAnimation(animationId uint) {
	a.currentAnimationId = animationId
}

func (a *AnimationPlayer) AddAnimation(animationSet *AnimationSet, animationId uint) {
	a.animationMap[animationId] = animationSet
}

func (a *AnimationPlayer) DeleteAnimation(animationId uint) {
	delete(a.animationMap, animationId)
}

func (a *AnimationPlayer) Update(dt float64) {
	if a.isPlaying {
		a.animationMap[a.currentAnimationId].Update(dt)
	}
}

func (a *AnimationPlayer) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
		a.animationMap[a.currentAnimationId].Draw(target, op)
}