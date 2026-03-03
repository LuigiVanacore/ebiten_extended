package physics

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestRigidBody2D_VelocityIntegration(t *testing.T) {
	world := NewPhysicsWorld()
	shape := collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10))
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	body := NewRigidBody2D(shape, mask)
	body.SetPosition(0, 0)
	body.SetVelocity(math2D.NewVector2D(100, 50))
	body.UsesGravity = false

	world.AddRigidBody(body)
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
	shape := collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10))
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	body := NewRigidBody2D(shape, mask)
	body.SetPosition(100, 100)
	body.UsesGravity = true
	body.GravityScale = 1

	world.AddRigidBody(body)
	world.Step(0.016)

	v := body.GetVelocity()
	if v.Y() <= 0 {
		t.Errorf("velocity Y should increase (gravity down), got %v", v.Y())
	}
}

func TestRigidBody2D_ApplyImpulse(t *testing.T) {
	body := NewRigidBody2D(
		collision.NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10)),
		collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1)),
	)
	body.ApplyImpulse(math2D.NewVector2D(10, -5))

	v := body.GetVelocity()
	if v.X() != 10 || v.Y() != -5 {
		t.Errorf("velocity after impulse expected (10,-5), got (%v,%v)", v.X(), v.Y())
	}
}
