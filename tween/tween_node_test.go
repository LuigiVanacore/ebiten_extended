package tween

import (
	"testing"

	"github.com/LuigiVanacore/ludum"
)

func TestTweenNode_Update(t *testing.T) {
	tw := NewTween("t", 0, 100, 1.0, Linear)
	node := NewTweenNode("tn", tw)
	var lastVal float32
	node.OnUpdate = func(v float32) { lastVal = v }
	// Advance ~0.5s: 30 frames at 1/60
	for i := 0; i < 30; i++ {
		node.Update()
	}
	if lastVal < 45 || lastVal > 55 {
		t.Errorf("OnUpdate: got %v want ~50", lastVal)
	}
	// Advance remaining ~0.5s to finish
	for i := 0; i < 35; i++ {
		node.Update()
	}
	if !node.IsFinished() {
		t.Error("expected finished")
	}
}

func TestTweenNode_OnComplete(t *testing.T) {
	tw := NewTween("t", 0, 1, 0.1, Linear)
	node := NewTweenNode("tn", tw)
	called := false
	node.OnComplete = func() { called = true }
	frames := int(0.2 / ludum.FIXED_DELTA)
	for i := 0; i < frames; i++ {
		node.Update()
	}
	if !called {
		t.Error("OnComplete should be called")
	}
}

func TestTweenNode_Restart(t *testing.T) {
	tw := NewTween("t", 0, 10, 1.0, Linear)
	node := NewTweenNode("tn", tw)
	frames := int(1.5 / ludum.FIXED_DELTA)
	for i := 0; i < frames; i++ {
		node.Update()
	}
	if !node.IsFinished() {
		t.Error("expected finished")
	}
	node.Restart()
	if node.IsFinished() {
		t.Error("restart should clear finished")
	}
}
