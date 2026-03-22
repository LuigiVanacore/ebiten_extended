package ebiten_extended

import "github.com/hajimehoshi/ebiten/v2"

type AnimationSprite struct {
	Node2D
	animationSet *AnimationSet
	layer        int
	isPlaying    bool
}

func NewAnimationSprite(name string, animationSet *AnimationSet, layer int, isPlaying bool) *AnimationSprite {
	return &AnimationSprite{
		Node2D:       *NewNode2D(name),
		animationSet: animationSet,
		layer:        layer,
		isPlaying:    isPlaying,
	}
}

func (a *AnimationSprite) Start() {
	a.isPlaying = true
}

func (a *AnimationSprite) Stop() {
	a.isPlaying = false
}

func (a *AnimationSprite) GetLayer() int {
	return a.layer
}

// Update advances the animation. Implements Updatable.
func (a *AnimationSprite) Update() {
	if a.isPlaying {
		if a.animationSet != nil {
			a.animationSet.Update()
		}
	}
}

func (a *AnimationSprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if a.animationSet != nil {
		a.animationSet.Draw(target, op)
	}
}
