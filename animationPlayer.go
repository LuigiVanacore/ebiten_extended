package ebiten_extended

import (
 
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)


type AnimationPlayer struct {
	Node2D
	animationMap map[string]*AnimationSet
	currentAnimationId string
	isPlaying bool
}

func NewAnimationPlayer() *AnimationPlayer {
	return &AnimationPlayer{ animationMap: make(map[string]*AnimationSet)}
}

func (a *AnimationPlayer) GetTransform() transform.Transform {
	return a.transform
}

func (a *AnimationPlayer) SetTransform(transform transform.Transform) {
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

func (a *AnimationPlayer) GetCurrentAnimation() string {
	return a.currentAnimationId
}

func (a *AnimationPlayer) SetCurrentAnimation(animationId string) {
	a.currentAnimationId = animationId
}

func (a *AnimationPlayer) AddAnimation(animationSet *AnimationSet, animationId string) {
	a.animationMap[animationId] = animationSet
}

func (a *AnimationPlayer) DeleteAnimation(animationId string) {
	delete(a.animationMap, animationId)
}

func (a *AnimationPlayer) Update() {
	if !a.isPlaying || a.currentAnimationId == "" {
		return
	}
	set, ok := a.animationMap[a.currentAnimationId]
	if !ok || set == nil {
		return
	}
	set.Update()
}

func (a *AnimationPlayer) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if a.currentAnimationId == "" {
		return
	}
	set, ok := a.animationMap[a.currentAnimationId]
	if !ok || set == nil {
		return
	}
	set.Draw(target, op)
}