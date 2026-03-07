package tween

import "github.com/LuigiVanacore/ebiten_extended"

// TweenNode is a Node2D that runs a Tween each frame. Add to the World; it implements Updatable.
// Use OnUpdate to apply the interpolated value (e.g. set position, scale, alpha).
type TweenNode struct {
	ebiten_extended.Node2D
	*Tween
	OnUpdate  func(value float32) // called each frame with the current interpolated value
	OnComplete func()             // called once when the tween finishes
	finished   bool
}

// NewTweenNode creates a TweenNode that runs the given Tween with delta from the game loop.
func NewTweenNode(name string, tween *Tween) *TweenNode {
	return &TweenNode{
		Node2D: *ebiten_extended.NewNode2D(name),
		Tween:  tween,
	}
}

// Update advances the tween and invokes callbacks. Implements Updatable.
func (n *TweenNode) Update() {
	if n.Tween == nil || n.finished {
		return
	}
	val, done := n.Tween.StepDelta(ebiten_extended.FIXED_DELTA)
	if n.OnUpdate != nil {
		n.OnUpdate(val)
	}
	if done && !n.finished {
		n.finished = true
		if n.OnComplete != nil {
			n.OnComplete()
		}
	}
}

// IsFinished returns true once the tween has completed.
func (n *TweenNode) IsFinished() bool {
	return n.finished
}

// Restart resets the tween and clears the finished state.
func (n *TweenNode) Restart() {
	n.finished = false
	n.Tween.Reset()
}
