package ebiten_extended

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// AnimationPlayer coordinates playback and switching between multiple named AnimationSet sequences for a node.
type AnimationPlayer struct {
	Node2D
	animationMap       map[string]*AnimationSet
	currentAnimationId string
	isPlaying          bool
}

// NewAnimationPlayer initializes an empty AnimationPlayer mapping string identifiers to individual AnimationSets.
func NewAnimationPlayer() *AnimationPlayer {
	return &AnimationPlayer{animationMap: make(map[string]*AnimationSet)}
}

// IsPlaying returns whether the player is currently actively tracking and updating frames.
func (a *AnimationPlayer) IsPlaying() bool {
	return a.isPlaying
}

// Start resumes or initiates the timing update playback of the currently designated animation.
func (a *AnimationPlayer) Start() {
	a.isPlaying = true
}

// Stop pauses the currently playing animation in its current sequential state.
func (a *AnimationPlayer) Stop() {
	a.isPlaying = false
}

// GetCurrentAnimation identifies the string key ID of the actively managed animation sequence.
func (a *AnimationPlayer) GetCurrentAnimation() string {
	return a.currentAnimationId
}

// SetCurrentAnimation transitions the player's active state to a new recognized animation map ID.
func (a *AnimationPlayer) SetCurrentAnimation(animationId string) {
	a.currentAnimationId = animationId
}

// AddAnimation indexes a provided AnimationSet into the player's dictionary under the specified key.
func (a *AnimationPlayer) AddAnimation(animationSet *AnimationSet, animationId string) {
	a.animationMap[animationId] = animationSet
}

// DeleteAnimation eliminates an assigned AnimationSet configuration from the player library.
func (a *AnimationPlayer) DeleteAnimation(animationId string) {
	delete(a.animationMap, animationId)
}

// Update advances the internal clock constraint of the actively executing sequence.
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

// Draw overlays the currently processed frame slice from the prevailing active animation sequence to the target image.
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
