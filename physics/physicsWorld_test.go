package physics

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestPhysicsWorld_StepNoOverlap(t *testing.T) {
	world := NewPhysicsWorld()
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))

	b1 := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	b1.SetPosition(0, 0)
	b1.SetVelocity(math2D.NewVector2D(50, 0))
	b1.UsesGravity = false

	b2 := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	b2.SetPosition(100, 0)
	b2.UsesGravity = false

	world.AddRigidBody(b1)
	world.AddRigidBody(b2)
	world.Step(0.016)

	// Both should move; no overlap so no push-out
	if b1.GetPosition().X() < 0.5 {
		t.Error("b1 should have moved")
	}
}

func TestPhysicsWorld_StepPushOut(t *testing.T) {
	world := NewPhysicsWorld()
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))

	b1 := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	b1.SetPosition(0, 0)
	b1.UsesGravity = false

	b2 := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	b2.SetPosition(5, 0) // overlapping (sum radii 20, distance 5)
	b2.UsesGravity = false

	world.AddRigidBody(b1)
	world.AddRigidBody(b2)
	world.Step(0.016)

	// After resolution they should not overlap (sum radii = 20)
	pos1 := b1.GetPosition()
	pos2 := b2.GetPosition()
	dist := math2D.SubtractVectors(pos2, pos1)
	distSq := math2D.DotProduct(dist, dist)
	minDistSq := 19.0 * 19.0
	if distSq < minDistSq {
		t.Errorf("bodies still overlapping: distance^2=%v, want >= %v", distSq, minDistSq)
	}
}

func TestPhysicsWorld_AddRemoveRigidBody(t *testing.T) {
	world := NewPhysicsWorld()
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	body := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)

	world.AddRigidBody(body)
	world.Step(0.016)

	world.RemoveRigidBody(body)
	world.Step(0.016)
	// Should not panic
}

func TestPhysicsWorld_StaticBodyNotMoved(t *testing.T) {
	world := NewPhysicsWorld()
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))

	dynamic := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	dynamic.SetPosition(0, 0)
	dynamic.SetVelocity(math2D.NewVector2D(100, 0))
	dynamic.UsesGravity = false

	static := NewRigidBody2D(collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(50, 50))), mask)
	static.SetPosition(30, 0)
	static.UsesGravity = false
	static.Static = true

	world.AddRigidBody(dynamic)
	world.AddRigidBody(static)

	initialStaticPos := static.GetPosition()
	world.Step(0.016)

	if static.GetPosition().X() != initialStaticPos.X() || static.GetPosition().Y() != initialStaticPos.Y() {
		t.Errorf("Static body should not move, got (%v,%v)", static.GetPosition().X(), static.GetPosition().Y())
	}
}

func TestPhysicsWorld_Gravity(t *testing.T) {
	world := NewPhysicsWorld()
	world.Gravity = math2D.NewVector2D(0, 400)
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))

	body := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	body.SetPosition(100, 100)
	body.UsesGravity = true

	world.AddRigidBody(body)
	world.Step(1.0 / 60.0) // one frame

	// v_y should increase (gravity * dt)
	vy := body.GetVelocity().Y()
	if vy <= 0 {
		t.Errorf("Gravity should increase vy, got %v", vy)
	}
}

func TestPhysicsWorld_VelocityZeroedOnCollision(t *testing.T) {
	world := NewPhysicsWorld()
	world.Gravity = math2D.NewVector2D(0, 400)
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))

	ball := NewRigidBody2D(collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)), mask)
	ball.SetPosition(50, 50)
	ball.SetVelocity(math2D.NewVector2D(0, 100))
	ball.UsesGravity = true

	floor := NewRigidBody2D(collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(200, 20))), mask)
	floor.SetPosition(50, 100)
	floor.UsesGravity = false
	floor.Static = true

	world.AddRigidBody(ball)
	world.AddRigidBody(floor)

	for i := 0; i < 60; i++ {
		world.Step(1.0 / 60.0)
	}

	vy := ball.GetVelocity().Y()
	if vy > 20 {
		t.Errorf("Ball landing on floor should have low vy (velocity zeroed on impact), got %v", vy)
	}
}
