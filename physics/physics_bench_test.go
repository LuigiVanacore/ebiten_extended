package physics

import (
	"testing"

	"github.com/LuigiVanacore/ludum/collision"
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/utils"
)

func BenchmarkPhysicsStep_50(b *testing.B) {
	mask := collision.NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	w := NewPhysicsWorld()
	for i := 0; i < 50; i++ {
		s := collision.NewCollisionCircle(math2d.NewCircle(math2d.ZeroVector2D(), 15))
		body, _ := NewRigidBody2D("b", s, mask)
		body.SetPosition(float64(i*25), 100)
		w.AddRigidBody(body)
	}
	dt := 1.0 / 60
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Step(dt)
	}
}
