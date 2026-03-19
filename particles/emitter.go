package particles

import (
	"image/color"
	"math/rand"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

// ParticleEmitter emits particles with configurable lifetime, velocity, scale, and color.
type ParticleEmitter struct {
	EmissionRate   float64       // particles per second; 0 = burst only
	LifetimeMin    float64       // min lifetime in seconds
	LifetimeMax    float64       // max lifetime in seconds
	VelocityMin    math2D.Vector2D
	VelocityMax    math2D.Vector2D
	ScaleMin       float64
	ScaleMax       float64
	ColorStart     color.Color
	ColorEnd       color.Color
	BurstCount     int           // particles to emit on Burst(); 0 = no burst
	Texture        *ebiten.Image // nil = draw as colored rect
	accumulator    float64
	particles      []*Particle
	maxParticles   int
}

// NewParticleEmitter creates an emitter with default values.
func NewParticleEmitter(maxParticles int) *ParticleEmitter {
	return &ParticleEmitter{
		LifetimeMin:  0.5,
		LifetimeMax:  1.5,
		VelocityMin: math2D.NewVector2D(-50, -100),
		VelocityMax: math2D.NewVector2D(50, -50),
		ScaleMin:    0.5,
		ScaleMax:    1.0,
		ColorStart:  color.White,
		ColorEnd:    color.Transparent,
		particles:   make([]*Particle, 0, maxParticles),
		maxParticles: maxParticles,
	}
}

// SetEmissionRate sets particles per second.
func (e *ParticleEmitter) SetEmissionRate(rate float64) {
	e.EmissionRate = rate
}

// SetLifetimeRange sets min and max lifetime in seconds.
func (e *ParticleEmitter) SetLifetimeRange(min, max float64) {
	e.LifetimeMin, e.LifetimeMax = min, max
}

// SetVelocityRange sets velocity bounds.
func (e *ParticleEmitter) SetVelocityRange(min, max math2D.Vector2D) {
	e.VelocityMin, e.VelocityMax = min, max
}

// SetScaleRange sets scale bounds.
func (e *ParticleEmitter) SetScaleRange(min, max float64) {
	e.ScaleMin, e.ScaleMax = min, max
}

// SetColorRange sets start and end colors for interpolation.
func (e *ParticleEmitter) SetColorRange(start, end color.Color) {
	e.ColorStart, e.ColorEnd = start, end
}

// SetTexture sets the particle texture. Nil = draw as colored rect.
func (e *ParticleEmitter) SetTexture(tex *ebiten.Image) {
	e.Texture = tex
}

// Burst emits BurstCount particles immediately.
func (e *ParticleEmitter) Burst(count int) {
	for i := 0; i < count && len(e.particles) < e.maxParticles; i++ {
		e.emitOne(math2D.ZeroVector2D())
	}
}

// emitOne creates and adds one particle at the given local offset.
func (e *ParticleEmitter) emitOne(offset math2D.Vector2D) {
	if len(e.particles) >= e.maxParticles {
		return
	}
	life := e.LifetimeMin
	if e.LifetimeMax > e.LifetimeMin {
		life += rand.Float64() * (e.LifetimeMax - e.LifetimeMin)
	}
	vx := e.VelocityMin.X() + rand.Float64()*(e.VelocityMax.X()-e.VelocityMin.X())
	vy := e.VelocityMin.Y() + rand.Float64()*(e.VelocityMax.Y()-e.VelocityMin.Y())
	scale := e.ScaleMin + rand.Float64()*(e.ScaleMax-e.ScaleMin)
	p := &Particle{
		Position:   offset,
		Velocity:   math2D.NewVector2D(vx, vy),
		Lifetime:   life,
		Age:        0,
		Scale:      scale,
		ColorStart: e.ColorStart,
		ColorEnd:   e.ColorEnd,
	}
	e.particles = append(e.particles, p)
}

// Update advances all particles by delta seconds and emits new ones based on EmissionRate.
func (e *ParticleEmitter) Update(delta float64, emitOffset math2D.Vector2D) {
	// Emit new particles
	if e.EmissionRate > 0 {
		e.accumulator += delta * e.EmissionRate
		for e.accumulator >= 1 && len(e.particles) < e.maxParticles {
			e.accumulator -= 1
			e.emitOne(emitOffset)
		}
	}

	// Update existing particles
	alive := 0
	for _, p := range e.particles {
		p.Age += delta
		if p.Age < p.Lifetime {
			p.Position.SetX(p.Position.X() + p.Velocity.X()*delta)
			p.Position.SetY(p.Position.Y() + p.Velocity.Y()*delta)
			e.particles[alive] = p
			alive++
		}
	}
	e.particles = e.particles[:alive]
}

// Particles returns the current live particles (read-only).
func (e *ParticleEmitter) Particles() []*Particle {
	return e.particles
}

// lerpColor interpolates between two colors by t (0-1).
func lerpColor(c1, c2 color.Color, t float64) color.Color {
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	tf := float64(t)
	return color.RGBA{
		R: uint8((float64(r1>>8)*(1-tf) + float64(r2>>8)*tf)),
		G: uint8((float64(g1>>8)*(1-tf) + float64(g2>>8)*tf)),
		B: uint8((float64(b1>>8)*(1-tf) + float64(b2>>8)*tf)),
		A: uint8((float64(a1>>8)*(1-tf) + float64(a2>>8)*tf)),
	}
}
