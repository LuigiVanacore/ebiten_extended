package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ProgressBarOrientation defines the fill direction.
type ProgressBarOrientation int

const (
	ProgressBarHorizontal ProgressBarOrientation = iota
	ProgressBarVertical
)

// ProgressBarNode represents a UI bar that fills reflecting a progress value (0.0 to 1.0).
type ProgressBarNode struct {
	PanelNode
	progress    float64 // Range: 0.0 to 1.0
	fillColor   color.Color
	orientation ProgressBarOrientation
}

// NewProgressBarNode creates a new ProgressBarNode with a default background and fill color.
func NewProgressBarNode(name string, width, height float64) *ProgressBarNode {
	p := NewPanelNode(name, width, height)
	p.SetBackgroundColor(color.RGBA{30, 30, 30, 255}) // Dark background by default

	return &ProgressBarNode{
		PanelNode:    *p,
		progress:     0.0,
		fillColor:   color.RGBA{0, 200, 0, 255}, // Green fill by default
		orientation: ProgressBarHorizontal,
	}
}

// SetOrientation sets the fill direction (horizontal or vertical).
func (p *ProgressBarNode) SetOrientation(o ProgressBarOrientation) {
	p.orientation = o
}

// SetProgress updates the progression value clamped between 0.0 and 1.0.
func (p *ProgressBarNode) SetProgress(val float64) {
	if val < 0 {
		val = 0
	}
	if val > 1 {
		val = 1
	}
	p.progress = val
}

// GetProgress returns the current progress value.
func (p *ProgressBarNode) GetProgress() float64 {
	return p.progress
}

// SetFillColor sets the color of the inner filled bar.
func (p *ProgressBarNode) SetFillColor(c color.Color) {
	p.fillColor = c
}

// Draw renders the background panel first, then draws the foreground filled bar.
func (p *ProgressBarNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	// First draw the Background using PanelNode's standard logic (Color or Image)
	p.PanelNode.Draw(target, op)

	// Custom Foreground logic for ProgressBar
	if p.progress > 0 && p.fillColor != nil {
		worldPos := p.GetWorldPosition()
		scale := p.GetWorldScale()
		scaledW := p.width * scale.X()
		scaledH := p.height * scale.Y()

		if p.orientation == ProgressBarHorizontal {
			fillWidth := scaledW * p.progress
			vector.DrawFilledRect(
				target,
				float32(worldPos.X()),
				float32(worldPos.Y()),
				float32(fillWidth),
				float32(scaledH),
				p.fillColor,
				true,
			)
		} else {
			fillHeight := scaledH * p.progress
			vector.DrawFilledRect(
				target,
				float32(worldPos.X()),
				float32(worldPos.Y()+scaledH-fillHeight),
				float32(scaledW),
				float32(fillHeight),
				p.fillColor,
				true,
			)
		}
	}
}
