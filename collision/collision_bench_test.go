package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func BenchmarkCheckCollision_100(b *testing.B) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	mgr := NewCollisionManager()
	for i := 0; i < 50; i++ {
		s := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10))
		c, _ := NewCollider("c", s, mask)
		c.SetPosition(float64(i*30), float64(i*30))
		mgr.AddParticipant(c)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mgr.CheckCollision()
	}
}

func BenchmarkOverlapPoint_100(b *testing.B) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	mgr := NewCollisionManager()
	for i := 0; i < 100; i++ {
		s := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
		c, _ := NewCollider("c", s, mask)
		c.SetPosition(float64(i*25), 50)
		mgr.AddParticipant(c)
	}
	pt := math2D.NewVector2D(100, 50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mgr.OverlapPoint(pt)
	}
}

func BenchmarkOverlapCircle_100(b *testing.B) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	mgr := NewCollisionManager()
	for i := 0; i < 100; i++ {
		s := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 10))
		c, _ := NewCollider("c", s, mask)
		c.SetPosition(float64(i*25), 50)
		mgr.AddParticipant(c)
	}
	center := math2D.NewVector2D(100, 50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mgr.OverlapCircle(center, 15)
	}
}

func BenchmarkOverlapRect_100(b *testing.B) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	mgr := NewCollisionManager()
	for i := 0; i < 100; i++ {
		s := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
		c, _ := NewCollider("c", s, mask)
		c.SetPosition(float64(i*25), 50)
		mgr.AddParticipant(c)
	}
	rect := math2D.NewRectangle(math2D.NewVector2D(80, 30), math2D.NewVector2D(50, 50))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = mgr.OverlapRect(rect)
	}
}
