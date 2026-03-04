package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestCollisionManager_AddParticipantNil(t *testing.T) {
	manager := NewCollisionManager()
	manager.AddParticipant(nil)
	manager.CheckCollision() // Should not panic with empty participants
}

func TestCollisionManager_SingleParticipant(t *testing.T) {
	manager := NewCollisionManager()
	shape := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(10, 10)))
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	c, err := NewCollider(shape, mask)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	manager.AddCollider(c)
	manager.CheckCollision()
	// Single participant - no pairs, no panic
}

func TestCollisionManager_CanCollideWithFilter(t *testing.T) {
	manager := NewCollisionManager()
	mask1 := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(2)) // layer 1 collides with 2
	mask2 := NewCollisionMask(utils.ByteSet(2), utils.ByteSet(1)) // layer 2 collides with 1
	mask3 := NewCollisionMask(utils.ByteSet(3), utils.ByteSet(3)) // layer 3 only with 3

	shape := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	c1, err := NewCollider(shape, mask1)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c2, err := NewCollider(shape, mask2)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c3, err := NewCollider(shape, mask3)
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	c1.SetPosition(0, 0)
	c2.SetPosition(5, 5)  // overlaps c1 (masks compatible)
	c3.SetPosition(10, 10) // overlaps c1,c2 but mask3 doesn't collide with 1 or 2

	manager.AddParticipant(c1)
	manager.AddParticipant(c2)
	manager.AddParticipant(c3)

	enter12 := 0
	enter13 := 0
	c1.OnCollisionEnter.Connect(nil, func(other *Collider) {
		if other == c2 {
			enter12++
		}
		if other == c3 {
			enter13++
		}
	})

	manager.CheckCollision()

	if enter12 != 1 {
		t.Errorf("c1-c2 should collide (mask compatible), got enter12=%d", enter12)
	}
	if enter13 != 0 {
		t.Errorf("c1-c3 should not collide (mask incompatible), got enter13=%d", enter13)
	}
}

func TestCollisionManager_Area2DOnBodyStay(t *testing.T) {
	manager := NewCollisionManager()
	areaShape := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(100, 100)))
	area, err := NewArea2D(areaShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	if err != nil {
		t.Fatalf("NewArea2D failed: %v", err)
	}
	area.SetPosition(100, 100)

	collShape := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	coll, err := NewCollider(collShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	coll.SetPosition(120, 120)

	manager.AddParticipant(coll)
	manager.AddParticipant(area)

	stayCount := 0
	area.OnBodyStay.Connect(nil, func(ev Area2DBodyEvent) {
		stayCount++
	})

	manager.CheckCollision() // Enter
	manager.CheckCollision() // Stay
	manager.CheckCollision() // Stay

	if stayCount != 2 {
		t.Errorf("OnBodyStay should fire 2 times (frames 2 and 3), got %d", stayCount)
	}
}

func TestCollisionManager_AABBBroadPhaseLargeBody(t *testing.T) {
	// Floor-like: wide rect at bottom. Ball at left edge. With AABB broad phase both share cells.
	manager := NewCollisionManager()
	manager.CellSize = 100

	floorShape := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(640, 40)))
	floor, err := NewArea2D(floorShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	if err != nil {
		t.Fatalf("NewArea2D failed: %v", err)
	}
	floor.SetPosition(320, 460)

	ballShape := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 20))
	ball, err := NewCollider(ballShape, NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)))
	if err != nil {
		t.Fatalf("NewCollider failed: %v", err)
	}
	ball.SetPosition(20, 450) // left edge of floor

	manager.AddParticipant(floor)
	manager.AddParticipant(ball)

	entered := false
	floor.OnBodyEntered.Connect(nil, func(ev Area2DBodyEvent) {
		entered = true
	})

	manager.CheckCollision()

	if !entered {
		t.Error("Ball at floor left edge should trigger Area2D OnBodyEntered (AABB broad phase)")
	}
}
