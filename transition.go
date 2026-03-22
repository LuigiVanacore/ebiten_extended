package ludum

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// TransitionState represents the phase of a transition.
type TransitionState int

const (
	TransitionIdle TransitionState = iota
	TransitionIn                   // fading/sliding in
	TransitionOut                  // fading/sliding out
	TransitionDone
)

// Transition provides fade/slide effects between scenes using a 0-1 progress value.
// Use with Tween or manual progress updates.
type Transition struct {
	State    TransitionState
	Progress float32 // 0 = start, 1 = end
	Duration float32
	elapsed  float32
}

// NewTransition creates a transition with the given duration.
func NewTransition(duration float32) *Transition {
	return &Transition{Duration: duration}
}

// StartIn begins an "in" transition (e.g. fade from black to scene).
func (t *Transition) StartIn() {
	t.State = TransitionIn
	t.Progress = 0
	t.elapsed = 0
}

// StartOut begins an "out" transition (e.g. fade from scene to black).
func (t *Transition) StartOut() {
	t.State = TransitionOut
	t.Progress = 0
	t.elapsed = 0
}

// Update advances the transition by delta seconds. Returns true when complete.
func (t *Transition) Update(delta float64) bool {
	if t.State != TransitionIn && t.State != TransitionOut {
		return true
	}
	t.elapsed += float32(delta)
	if t.Duration > 0 {
		t.Progress = t.elapsed / t.Duration
		if t.Progress >= 1 {
			t.Progress = 1
			t.State = TransitionDone
			return true
		}
	}
	return false
}

// Alpha returns the overlay alpha for a fade effect. 0 = transparent (scene visible), 1 = opaque (black).
func (t *Transition) Alpha() float32 {
	switch t.State {
	case TransitionIn:
		return 1 - t.Progress // fade from black to scene
	case TransitionOut:
		return t.Progress // fade from scene to black
	default:
		return 0
	}
}

// DrawFadeOverlay draws a fullscreen fade overlay onto the target.
// Call after drawing the scene. Color is typically black; alpha comes from Alpha().
func (t *Transition) DrawFadeOverlay(target *ebiten.Image, c color.Color) {
	a := t.Alpha()
	if a <= 0 {
		return
	}
	r, g, b, _ := c.RGBA()
	alpha := uint8(float64(a) * 255)
	clr := color.RGBA{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(b >> 8), A: alpha}
	bounds := target.Bounds()
	vector.DrawFilledRect(target, 0, 0, float32(bounds.Dx()), float32(bounds.Dy()), clr, true)
}
