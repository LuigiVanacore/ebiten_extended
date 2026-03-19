package particles

import (
	"image/color"
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestNewParticleEmitter(t *testing.T) {
	e := NewParticleEmitter(100)
	if e == nil {
		t.Fatal("NewParticleEmitter returned nil")
	}
	if e.maxParticles != 100 {
		t.Errorf("maxParticles: got %d, want 100", e.maxParticles)
	}
	if e.LifetimeMin < e.LifetimeMax {
		// default: 0.5 and 1.5
	}
	if len(e.Particles()) != 0 {
		t.Errorf("new emitter should have 0 particles, got %d", len(e.Particles()))
	}
}

func TestParticleEmitterBurst(t *testing.T) {
	e := NewParticleEmitter(10)
	e.Burst(5)
	if len(e.Particles()) != 5 {
		t.Errorf("Burst(5): got %d particles, want 5", len(e.Particles()))
	}

	e.Burst(10) // try to burst 10 more, but max is 10 total
	if len(e.Particles()) != 10 {
		t.Errorf("Burst with max cap: got %d particles, want 10", len(e.Particles()))
	}
}

func TestParticleEmitterBurstRespectsMax(t *testing.T) {
	e := NewParticleEmitter(3)
	e.Burst(100)
	if len(e.Particles()) != 3 {
		t.Errorf("Burst(100) with maxParticles=3: got %d, want 3", len(e.Particles()))
	}
}

func TestParticleEmitterUpdateEmission(t *testing.T) {
	e := NewParticleEmitter(50)
	e.SetEmissionRate(60) // 60 particles per second = 1 per frame at 60fps
	e.SetLifetimeRange(1, 1) // deterministic lifetime
	e.SetVelocityRange(math2D.ZeroVector2D(), math2D.ZeroVector2D())

	// Update for 1 second
	for i := 0; i < 60; i++ {
		e.Update(1.0/60.0, math2D.ZeroVector2D())
	}
	count := len(e.Particles())
	if count < 40 || count > 70 {
		t.Errorf("After 1s at 60/s: got %d particles, expected ~60 (40-70)", count)
	}
}

func TestParticleEmitterUpdateAging(t *testing.T) {
	e := NewParticleEmitter(20)
	e.SetEmissionRate(0)
	e.SetLifetimeRange(0.1, 0.1) // 0.1s lifetime
	e.SetVelocityRange(math2D.ZeroVector2D(), math2D.ZeroVector2D())
	e.Burst(5)

	e.Update(0.05, math2D.ZeroVector2D())
	if len(e.Particles()) != 5 {
		t.Errorf("after 0.05s: want 5 particles, got %d", len(e.Particles()))
	}

	e.Update(0.1, math2D.ZeroVector2D()) // total 0.15s, all should be dead
	if len(e.Particles()) != 0 {
		t.Errorf("after 0.15s (lifetime 0.1): want 0 particles, got %d", len(e.Particles()))
	}
}

func TestParticleEmitterSetMethods(t *testing.T) {
	e := NewParticleEmitter(10)
	e.SetEmissionRate(100)
	e.SetLifetimeRange(0.5, 2.0)
	e.SetVelocityRange(math2D.NewVector2D(-1, -1), math2D.NewVector2D(1, 1))
	e.SetScaleRange(0.5, 2.0)
	e.SetColorRange(color.White, color.Black)

	if e.EmissionRate != 100 {
		t.Errorf("SetEmissionRate: got %v", e.EmissionRate)
	}
	if e.LifetimeMin != 0.5 || e.LifetimeMax != 2.0 {
		t.Errorf("SetLifetimeRange: got %v-%v", e.LifetimeMin, e.LifetimeMax)
	}
}

func TestParticleIsAliveAndProgress(t *testing.T) {
	p := &Particle{
		Lifetime: 1.0,
		Age:      0,
	}
	if !p.IsAlive() {
		t.Error("particle with Age=0 should be alive")
	}
	if p.Progress() != 0 {
		t.Errorf("Progress at start: got %v, want 0", p.Progress())
	}

	p.Age = 0.5
	if !p.IsAlive() {
		t.Error("particle at half lifetime should be alive")
	}
	if p.Progress() != 0.5 {
		t.Errorf("Progress at half: got %v, want 0.5", p.Progress())
	}

	p.Age = 1.0
	if p.IsAlive() {
		t.Error("particle at full lifetime should be dead")
	}
	if p.Progress() != 1.0 {
		t.Errorf("Progress at end: got %v, want 1.0", p.Progress())
	}
}

func TestParticleProgressZeroLifetime(t *testing.T) {
	p := &Particle{Lifetime: 0, Age: 0}
	if p.Progress() != 1 {
		t.Errorf("zero lifetime: Progress should be 1, got %v", p.Progress())
	}
}

func TestParticleEmitterNodeCreation(t *testing.T) {
	emitter := NewParticleEmitter(50)
	node := NewParticleEmitterNode("test_emitter", emitter, 1)
	if node == nil {
		t.Fatal("NewParticleEmitterNode returned nil")
	}
	if node.GetEmitter() != emitter {
		t.Error("GetEmitter should return the same emitter")
	}
	if node.GetLayer() != 1 {
		t.Errorf("GetLayer: got %d, want 1", node.GetLayer())
	}

	node.SetLayer(2)
	if node.GetLayer() != 2 {
		t.Errorf("SetLayer: got %d, want 2", node.GetLayer())
	}
}

func TestParticleEmitterNodeDrawNoPanic(t *testing.T) {
	// Draw with rect path (no texture)
	emitter := NewParticleEmitter(10)
	emitter.SetEmissionRate(0)
	emitter.SetLifetimeRange(0.5, 0.5)
	emitter.Burst(3)
	node := NewParticleEmitterNode("draw_test", emitter, 0)

	target := ebiten.NewImage(64, 64)
	defer target.Deallocate()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(32, 32)
	node.Draw(target, op)

	// Draw with nil op (defensive)
	node.Draw(target, nil)
}

func TestParticleEmitterNodeDrawWithTextureNoPanic(t *testing.T) {
	tex := ebiten.NewImage(8, 8)
	defer tex.Deallocate()

	emitter := NewParticleEmitter(10)
	emitter.SetEmissionRate(0)
	emitter.SetLifetimeRange(1, 1)
	emitter.SetTexture(tex)
	emitter.Burst(2)
	node := NewParticleEmitterNode("draw_tex_test", emitter, 0)

	target := ebiten.NewImage(64, 64)
	defer target.Deallocate()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(16, 16)
	node.Draw(target, op)
}
