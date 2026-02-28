package input

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestGodotLikeInputMapping(t *testing.T) {
	manager := NewInputManager()

	// 1. Add "shoot" Action mapped to both KEYBOARD and MOUSE
	manager.AddAction("shoot", NewKeyRawInputButton(ebiten.KeySpace))
	manager.AddAction("shoot", NewMouseRawInputButton(ebiten.MouseButtonLeft))

	// By default, since we are not actively pressing anything in test context,
	// ebiten.IsKeyPressed will return false. We verify the API works without panicking.
	if manager.IsActionPressed("shoot") {
		t.Error("Did not expect 'shoot' to be pressed initially")
	}

	if manager.IsActionJustPressed("shoot") {
		t.Error("Did not expect 'shoot' to be just pressed initially")
	}

	if !manager.IsActionReleased("shoot") {
		t.Error("Expected 'shoot' to report as released")
	}

	// 2. Add an action then remove it
	manager.AddAction("jump", NewKeyRawInputButton(ebiten.KeyW))
	manager.RemoveAction("jump")

	if manager.IsActionPressed("jump") {
		t.Error("Removed action must not report as pressed")
	}
}
