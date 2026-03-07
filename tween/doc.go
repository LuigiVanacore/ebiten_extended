// Package tween provides simple tweening for animating values over time.
//
// Use NewTween and call StepDelta(ebiten_extended.FIXED_DELTA) for fixed
// independent animations. For integration with the scene graph, add a TweenNode
// to the World; it implements Updatable and runs the tween automatically.
// Sequence chains multiple tweens and also supports StepDelta and Update.
package tween
