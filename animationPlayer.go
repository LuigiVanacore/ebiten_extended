package ebiten_extended

import (
 
	"github.com/hajimehoshi/ebiten/v2"
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
	if a.isPlaying {
		animationSet := a.animationMap[a.currentAnimationId]
		if animationSet != nil {
			animationSet.Update()
		}
	}
}

func (a *AnimationPlayer) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
		animationSet := a.animationMap[a.currentAnimationId]
		if animationSet != nil {
			animationSet.Draw(target, op)
		}
}