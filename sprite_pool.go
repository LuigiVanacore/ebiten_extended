package ludum

import (
	"sync"

	"github.com/LuigiVanacore/ludum/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// SpritePool reuses Sprite instances to reduce allocations (e.g. for projectiles, particles).
// Sprites are created from a template texture and reset when returned via Put.
type SpritePool struct {
	pool     *utils.Pool[*Sprite]
	template *ebiten.Image
	layerIdx int
	pivot    bool
	mu       sync.Mutex
}

// NewSpritePool creates a pool of sprites sharing the same texture.
// template is the source image; layerIndex and isPivotToCenter match NewSprite.
// initialSize preallocates that many sprites in the pool.
func NewSpritePool(template *ebiten.Image, layerIndex int, isPivotToCenter bool, initialSize int) *SpritePool {
	if template == nil {
		return &SpritePool{template: nil, layerIdx: layerIndex, pivot: isPivotToCenter}
	}
	factory := func() *Sprite {
		return NewSprite("pooled_sprite", template, layerIndex, isPivotToCenter)
	}
	reset := func(s *Sprite) {
		s.SetPosition(0, 0)
		s.SetRotation(0)
		s.SetScale(1, 1)
		s.SetVisible(true)
		s.SetColorScale(ebiten.ColorScale{})
		s.SetFlipX(false)
		s.SetFlipY(false)
		s.SetAlpha(1)
	}
	p := utils.NewPool(factory, reset)
	sp := &SpritePool{pool: p, template: template, layerIdx: layerIndex, pivot: isPivotToCenter}
	for i := 0; i < initialSize; i++ {
		sp.Put(NewSprite("pooled_sprite", template, layerIndex, isPivotToCenter))
	}
	return sp
}

// Get returns a sprite from the pool, or creates a new one if the pool is empty.
// The sprite is reset to default state (position 0,0, visible, no tint, no flip).
func (p *SpritePool) Get() *Sprite {
	if p.template == nil {
		return NewSprite("pooled_sprite", nil, p.layerIdx, p.pivot)
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.pool.Get()
}

// Put returns a sprite to the pool for reuse. The sprite is reset before being stored.
func (p *SpritePool) Put(s *Sprite) {
	if p.template == nil || s == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	p.pool.Put(s)
}

// Clear discards all pooled sprites.
func (p *SpritePool) Clear() {
	if p.pool != nil {
		p.mu.Lock()
		defer p.mu.Unlock()
		p.pool.Clear()
	}
}
