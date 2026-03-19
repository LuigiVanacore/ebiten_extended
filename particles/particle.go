package particles

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

// Particle holds the state of a single particle.
type Particle struct {
	Position   math2D.Vector2D
	Velocity   math2D.Vector2D
	Lifetime   float64   // total lifetime in seconds
	Age        float64   // elapsed time
	Scale      float64
	ColorStart color.Color
	ColorEnd   color.Color
}

// IsAlive returns true if the particle has not yet reached its lifetime.
func (p *Particle) IsAlive() bool {
	return p.Age < p.Lifetime
}

// Progress returns 0-1 normalized age for interpolation (e.g. color fade).
func (p *Particle) Progress() float64 {
	if p.Lifetime <= 0 {
		return 1
	}
	return p.Age / p.Lifetime
}
