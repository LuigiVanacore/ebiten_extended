package physics

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestKinematic_PushesDynamic(t *testing.T) {
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	dynShape := collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	dyn, _ := NewRigidBody2D("dyn", dynShape, mask)
	dyn.SetPosition(50, 50)

	kinShape := collision.NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	kin, _ := NewRigidBody2D("kin", kinShape, mask)
	kin.Kinematic = true
	kin.SetPosition(55, 50) // overlapping dyn

	w := NewPhysicsWorld()
	w.AddRigidBody(dyn)
	w.AddRigidBody(kin)
	w.Step(1.0 / 60)

	// Dynamic should be pushed left, kinematic unchanged
	if kin.GetPosition().X() != 55 {
		t.Errorf("kinematic moved: x=%v", kin.GetPosition().X())
	}
	// Dyn x should be less than 50 (pushed left)
	if dyn.GetPosition().X() >= 50 {
		t.Errorf("dynamic not pushed: x=%v", dyn.GetPosition().X())
	}
}
