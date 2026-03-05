// Package ui provides 2D scene graph nodes for user interface elements like Panels and Buttons.
package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// PanelNode represents a generic UI container that occupies a specific screen area (Width x Height).
// It can optionally draw a background color and/or a background image.
type PanelNode struct {
	ebiten_extended.Node2D
	width           float64
	height          float64
	backgroundColor color.Color
	backgroundImage *ebiten.Image
	layer           int
}

// NewPanelNode creates a new geometric UI panel.
func NewPanelNode(name string, width, height float64) *PanelNode {
	return &PanelNode{
		Node2D: *ebiten_extended.NewNode2D(name),
		width:  width,
		height: height,
	}
}

// SetSize updates the panel's dimensions.
func (p *PanelNode) SetSize(w, h float64) {
	p.width = w
	p.height = h
}

// GetSize returns the panel's dimensions.
func (p *PanelNode) GetSize() (float64, float64) {
	return p.width, p.height
}

// SetBackgroundColor sets the solid background color of the panel. Set to nil for transparency.
func (p *PanelNode) SetBackgroundColor(c color.Color) {
	p.backgroundColor = c
}

// SetBackgroundImage sets an image to be drawn stretched over the panel area. Set to nil to remove.
func (p *PanelNode) SetBackgroundImage(img *ebiten.Image) {
	p.backgroundImage = img
}

// GetBackgroundImage returns the current background image, or nil.
func (p *PanelNode) GetBackgroundImage() *ebiten.Image {
	return p.backgroundImage
}

// GetLayer returns the render layer of this panel.
func (p *PanelNode) GetLayer() int {
	return p.layer
}

// SetLayer sets the render layer of this panel.
func (p *PanelNode) SetLayer(layer int) {
	p.layer = layer
}

// Draw renders the panel background at its world-space position.
// UI panels intentionally draw in world/screen space and do not inherit the camera transform
// from the op GeoM. This is the expected behaviour for HUD/overlay elements that should remain
// fixed on screen regardless of camera movement. If you need a panel that moves with the world,
// extract the translation from op.GeoM and apply it manually.
func (p *PanelNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	worldPos := p.GetWorldPosition()
	scale := p.GetWorldScale()

	if p.backgroundColor != nil {
		// Fill a rectangle based on the calculated world coordinates and size (width/height are scaled)
		vector.DrawFilledRect(
			target,
			float32(worldPos.X()),
			float32(worldPos.Y()),
			float32(p.width*scale.X()),
			float32(p.height*scale.Y()),
			p.backgroundColor,
			true,
		)
	}

	if p.backgroundImage != nil {
		bounds := p.backgroundImage.Bounds()
		imgW, imgH := float64(bounds.Dx()), float64(bounds.Dy())
		if imgW > 0 && imgH > 0 {
			drawOp := &ebiten.DrawImageOptions{}
			// Scale the image to fit the panel dimension
			scaleX := (p.width * scale.X()) / imgW
			scaleY := (p.height * scale.Y()) / imgH
			drawOp.GeoM.Scale(scaleX, scaleY)
			drawOp.GeoM.Translate(worldPos.X(), worldPos.Y())

			// If a parent op is provided with color scaling, we could apply it, but UI panels usually draw absolute.
			if op != nil {
				drawOp.ColorScale = op.ColorScale
			}

			target.DrawImage(p.backgroundImage, drawOp)
		}
	}

	// Because it's a Node2D, its children will be automatically drawn by the World graph
}

// ContainsPoint returns true if the specified global (x, y) point falls within the panel's bounding box.
func (p *PanelNode) ContainsPoint(x, y float64) bool {
	pos := p.GetWorldPosition()
	scale := p.GetWorldScale()

	minX := pos.X()
	minY := pos.Y()
	// Note: Origin is assumed to be top-left
	maxX := minX + p.width*scale.X()
	maxY := minY + p.height*scale.Y()

	return x >= minX && x <= maxX && y >= minY && y <= maxY
}
