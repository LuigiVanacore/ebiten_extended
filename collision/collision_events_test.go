package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestCollisionLifecycleEvents(t *testing.T) {
	manager := NewCollisionManager()

	// NewCollisionRect was not exposed, creating direct struct
	shape1 := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(100, 100))}
	mask1 := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	c1, err := NewCollider("c1", shape1, mask1)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c1.SetPosition(0, 0)

	shape2 := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(100, 100))}
	mask2 := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	c2, err := NewCollider("c2", shape2, mask2)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c2.SetPosition(0, 0)

	manager.AddCollider(c1)
	manager.AddCollider(c2)

	enterCalled := 0
	stayCalled := 0
	exitCalled := 0

	c1.CollisionEnter().Connect(nil, func(arg *Collider) {
		enterCalled++
	})
	c1.CollisionStay().Connect(nil, func(arg *Collider) {
		stayCalled++
	})
	c1.CollisionExit().Connect(nil, func(arg *Collider) {
		exitCalled++
	})

	// Tick 1: overlapping exactly at (0,0) -> Enter should trigger
	manager.CheckCollision()

	if enterCalled != 1 {
		t.Errorf("Expected 1 Enter call, got %d", enterCalled)
	}
	if stayCalled != 0 {
		t.Errorf("Expected 0 Stay calls initially, got %d", stayCalled)
	}

	// Tick 2: still overlapping -> Stay should trigger
	manager.CheckCollision()

	if enterCalled != 1 {
		t.Errorf("Expected Enter to remain 1, got %d", enterCalled)
	}
	if stayCalled != 1 {
		t.Errorf("Expected 1 Stay call, got %d", stayCalled)
	}

	// Tick 3: move c2 far away -> Exit should trigger
	c2.SetPosition(5000, 5000)

	manager.CheckCollision()

	if exitCalled != 1 {
		t.Errorf("Expected 1 Exit call, got %d", exitCalled)
	}
	if stayCalled != 1 {
		t.Errorf("Expected Stay to not increment when separated, got %d", stayCalled)
	}
}

func TestCollisionManager_ColliderRetrocompat(t *testing.T) {
	manager := NewCollisionManager()
	shape := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(100, 100))}
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	c1, err := NewCollider("c1", shape, mask)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c2, err := NewCollider("c2", shape, mask)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c1.SetPosition(0, 0)
	c2.SetPosition(0, 0)

	manager.AddCollider(c1)
	manager.AddCollider(c2)

	enterCount := 0
	c1.CollisionEnter().Connect(nil, func(*Collider) { enterCount++ })

	manager.CheckCollision()

	if enterCount != 1 {
		t.Errorf("AddCollider retrocompat: expected 1 Enter, got %d", enterCount)
	}
}

func TestSpatialGridSeparation(t *testing.T) {
	manager := NewCollisionManager()

	shape1 := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(10, 10))}
	mask1 := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	c1, err := NewCollider("c1", shape1, mask1)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c1.SetPosition(0, 0) // Cell x=0, y=0

	shape2 := &CollisionRect{rectangle: math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(10, 10))}
	mask2 := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	c2, err := NewCollider("c2", shape2, mask2)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c2.SetPosition(2000, 2000) // Cell x=20, y=20 (Grid won't match)

	manager.AddCollider(c1)
	manager.AddCollider(c2)

	enterCalled := 0

	c1.CollisionEnter().Connect(nil, func(arg *Collider) {
		enterCalled++
	})

	manager.CheckCollision()

	if enterCalled > 0 {
		t.Errorf("Grid isolation failed, colliders at distance detected false positive: %d", enterCalled)
	}
}
