package physics

import (
	"testing"

	"github.com/LuigiVanacore/ludum/collision"
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/utils"
)

func mustRigidBody(t *testing.T, shape collision.CollisionShape, mask collision.CollisionMask) *RigidBody2D {
	t.Helper()
	body, err := NewRigidBody2D("body", shape, mask)
	if err != nil {
		t.Fatalf("NewRigidBody2D failed: %v", err)
	}
	return body
}

func TestRigidBody2D_VelocityIntegration(t *testing.T) {
	world := NewPhysicsWorld()
	shape := collision.NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 10))
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	body := mustRigidBody(t, shape, mask)
	body.SetPosition(0, 0)
	body.SetVelocity(math2d.NewVector2D(100, 50))
	body.UsesGravity = false

	if err := world.AddRigidBody(body); err != nil {
		t.Fatalf("AddRigidBody failed: %v", err)
	}
	world.Step(0.016) // ~1/60 sec

	pos := body.GetPosition()
	if pos.X() < 1.5 || pos.X() > 2.0 {
		t.Errorf("expected X ~1.6, got %v", pos.X())
	}
	if pos.Y() < 0.7 || pos.Y() > 1.0 {
		t.Errorf("expected Y ~0.8, got %v", pos.Y())
	}
}

func TestRigidBody2D_Gravity(t *testing.T) {
	world := NewPhysicsWorld()
	shape := collision.NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 10))
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	body := mustRigidBody(t, shape, mask)
	body.SetPosition(100, 100)
	body.UsesGravity = true
	body.GravityScale = 1

	if err := world.AddRigidBody(body); err != nil {
		t.Fatalf("AddRigidBody failed: %v", err)
	}
	world.Step(0.016)

	v := body.GetVelocity()
	if v.Y() <= 0 {
		t.Errorf("velocity Y should increase (gravity down), got %v", v.Y())
	}
}

func TestRigidBody2D_ApplyImpulse(t *testing.T) {
	body := mustRigidBody(t,
		collision.NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 10)),
		collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)),
	)
	body.ApplyImpulse(math2d.NewVector2D(10, -5))

	v := body.GetVelocity()
	if v.X() != 10 || v.Y() != -5 {
		t.Errorf("velocity after impulse expected (10,-5), got (%v,%v)", v.X(), v.Y())
	}
}
