package ebiten_extended

import (
	"testing"
)

func TestNewEngine(t *testing.T) {
	// Test that multiple engines can be instantiated
	// without cross-contamination.
	engine1 := NewEngine()
	engine2 := NewEngine()

	if engine1 == nil {
		t.Fatalf("Engine 1 is nil")
	}

	if engine2 == nil {
		t.Fatalf("Engine 2 is nil")
	}

	if engine1 == engine2 {
		t.Errorf("NewEngine() should return a new instance, but got the same one")
	}

	// Internal managers shouldn't be the same pointers
	if engine1.World() == engine2.World() {
		t.Errorf("Engine instances share the same World pointer")
	}
	if engine1.Input() == engine2.Input() {
		t.Errorf("Engine instances share the same InputManager pointer")
	}
	if engine1.Resources() == engine2.Resources() {
		t.Errorf("Engine instances share the same ResourceManager pointer")
	}
}

func TestEngineWorldInitialization(t *testing.T) {
	engine := NewEngine()
	if engine.World() == nil {
		t.Fatalf("Engine world is not initialized")
	}
}
