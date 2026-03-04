package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// ProgressBarNode represents a UI bar that fills horizontally reflecting a progress value.
type ProgressBarNode struct {
	PanelNode
	progress  float64 // Range: 0.0 to 1.0
	fillColor color.Color
}

// NewProgressBarNode creates a new ProgressBarNode with a default background and fill color.
func NewProgressBarNode(name string, width, height float64) *ProgressBarNode {
	p := NewPanelNode(name, width, height)
	p.SetBackgroundColor(color.RGBA{30, 30, 30, 255}) // Dark background by default

	return &ProgressBarNode{
		PanelNode: *p,
		progress:  0.0,
		fillColor: color.RGBA{0, 200, 0, 255}, // Green fill by default
	}
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

		// Calculate fill width based on progress
		fillWidth := p.width * scale.X() * p.progress

		var fillHeight float32
		if p.backgroundImage != nil {
			// Calculate if scaled img height exists
			fillHeight = float32(p.height * scale.Y())
		} else {
			fillHeight = float32(p.height * scale.Y())
		}

		vector.DrawFilledRect(
			target,
			float32(worldPos.X()),
			float32(worldPos.Y()),
			float32(fillWidth),
			fillHeight,
			p.fillColor,
			true,
		)
	}
}
