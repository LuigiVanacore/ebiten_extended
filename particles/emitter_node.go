package particles

import (
	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ParticleEmitterNode is a Node2D that hosts a ParticleEmitter and integrates with the scene graph.
// Implements Drawable and Updatable.
type ParticleEmitterNode struct {
	ludum.Node2D
	emitter *ParticleEmitter
	layer   int
}

// NewParticleEmitterNode creates a ParticleEmitterNode with the given emitter.
func NewParticleEmitterNode(name string, emitter *ParticleEmitter, layer int) *ParticleEmitterNode {
	return &ParticleEmitterNode{
		Node2D:  *ludum.NewNode2D(name),
		emitter: emitter,
		layer:   layer,
	}
}

// GetEmitter returns the particle emitter.
func (n *ParticleEmitterNode) GetEmitter() *ParticleEmitter {
	return n.emitter
}

// GetLayer implements Drawable.
func (n *ParticleEmitterNode) GetLayer() int {
	return n.layer
}

// SetLayer sets the draw layer.
func (n *ParticleEmitterNode) SetLayer(l int) {
	n.layer = l
}

// Update advances the emitter by one frame. Implements Updatable.
func (n *ParticleEmitterNode) Update() {
	if n.emitter != nil {
		n.emitter.Update(ludum.FIXED_DELTA, math2d.ZeroVector2D())
	}
}

// Draw renders all live particles. Implements Drawable.
// Particle positions are in the emitter's local space; they are transformed to world space
// using the node's GeoM (op), so rotation and scale of the node are applied correctly.
func (n *ParticleEmitterNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if n.emitter == nil {
		return
	}
	var geoM ebiten.GeoM
	if op != nil {
		geoM = op.GeoM
	}

	for _, p := range n.emitter.Particles() {
		if !p.IsAlive() {
			continue
		}
		prog := p.Progress()
		clr := lerpColor(p.ColorStart, p.ColorEnd, prog)
		// Transform particle position from local to world space (handles node position, rotation, scale)
		worldX, worldY := geoM.Apply(p.Position.X(), p.Position.Y())
		x, y := float32(worldX), float32(worldY)
		s := float32(p.Scale)

		if n.emitter.Texture != nil {
			bounds := n.emitter.Texture.Bounds()
			w, h := float32(bounds.Dx()), float32(bounds.Dy())
			drawOp := &ebiten.DrawImageOptions{}
			drawOp.GeoM.Translate(-float64(w)/2, -float64(h)/2)
			drawOp.GeoM.Scale(float64(s), float64(s))
			drawOp.GeoM.Translate(float64(x), float64(y))
			cr, cg, cb, ca := clr.RGBA()
			drawOp.ColorScale.SetR(float32(cr>>8) / 255)
			drawOp.ColorScale.SetG(float32(cg>>8) / 255)
			drawOp.ColorScale.SetB(float32(cb>>8) / 255)
			drawOp.ColorScale.SetA(float32(ca>>8) / 255)
			target.DrawImage(n.emitter.Texture, drawOp)
		} else {
			half := s * 4 // 8x8 rect at scale 1
			vector.DrawFilledRect(target, x-half, y-half, half*2, half*2, clr, true)
		}
	}
}
