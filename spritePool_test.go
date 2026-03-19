package ebiten_extended

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestNewSpritePool(t *testing.T) {
	img := ebiten.NewImage(32, 32)
	defer img.Deallocate()

	pool := NewSpritePool(img, 0, false, 2)
	if pool == nil {
		t.Fatal("NewSpritePool returned nil")
	}
	// initialSize=2 means 2 sprites pre-allocated in the pool
	if pool.pool.Size() != 2 {
		t.Errorf("initial pool size: got %d, want 2", pool.pool.Size())
	}
}

func TestNewSpritePoolNilTemplate(t *testing.T) {
	pool := NewSpritePool(nil, 0, false, 0)
	if pool == nil {
		t.Fatal("NewSpritePool with nil template returned nil")
	}
	if pool.pool != nil {
		t.Error("pool with nil template should have nil internal pool")
	}
}

func TestSpritePoolGetPut(t *testing.T) {
	img := ebiten.NewImage(16, 16)
	defer img.Deallocate()

	pool := NewSpritePool(img, 0, true, 0)

	s1 := pool.Get()
	if s1 == nil {
		t.Fatal("Get returned nil")
	}
	if s1.GetTexture() != img {
		t.Error("Get sprite should have template texture")
	}

	// Modify sprite
	s1.SetPosition(100, 200)
	s1.SetRotation(1.5)
	s1.SetScale(2, 2)

	pool.Put(s1)

	s2 := pool.Get()
	if s2 == nil {
		t.Fatal("Get after Put returned nil")
	}
	// Should be reset
	pos := s2.GetPosition()
	if pos.X() != 0 || pos.Y() != 0 {
		t.Errorf("Put sprite should be reset: position got (%v, %v), want (0, 0)", pos.X(), pos.Y())
	}
	if s2.GetRotation() != 0 {
		t.Errorf("Put sprite rotation should be 0, got %v", s2.GetRotation())
	}
	scale := s2.GetScale()
	if scale.X() != 1 || scale.Y() != 1 {
		t.Errorf("Put sprite scale should be (1,1), got (%v, %v)", scale.X(), scale.Y())
	}
}

func TestSpritePoolGetNilTemplateReturnsNewSprite(t *testing.T) {
	pool := NewSpritePool(nil, 0, false, 0)
	s := pool.Get()
	if s == nil {
		t.Fatal("Get with nil template returned nil")
	}
	if s.GetTexture() != nil {
		t.Error("sprite from nil-template pool should have nil texture")
	}
}

func TestSpritePoolPutNilIgnored(t *testing.T) {
	img := ebiten.NewImage(8, 8)
	defer img.Deallocate()

	pool := NewSpritePool(img, 0, false, 0)
	pool.Put(nil) // should not panic
}

func TestSpritePoolPutWithNilTemplateIgnored(t *testing.T) {
	pool := NewSpritePool(nil, 0, false, 0)
	s := pool.Get()
	pool.Put(s) // Put on nil-template pool should be no-op, no panic
}

func TestSpritePoolClear(t *testing.T) {
	img := ebiten.NewImage(16, 16)
	defer img.Deallocate()

	pool := NewSpritePool(img, 0, false, 3)
	if pool.pool.Size() != 3 {
		t.Errorf("initial size: got %d, want 3", pool.pool.Size())
	}

	pool.Clear()
	if pool.pool.Size() != 0 {
		t.Errorf("after Clear: got size %d, want 0", pool.pool.Size())
	}

	// Get should still work (creates new)
	s := pool.Get()
	if s == nil {
		t.Fatal("Get after Clear returned nil")
	}
}
