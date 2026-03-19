package ui

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

// Layout arranges child nodes spatially. Apply positions children in sequence.
type Layout interface {
	// Apply positions the children of the given parent node.
	// Children must be transform.Transformable (e.g. Node2D) for positioning to take effect.
	Apply(parent ebiten_extended.SceneNode)
}

// HBoxLayout arranges children in a horizontal row (left to right).
type HBoxLayout struct {
	Spacing float64 // gap between children
	Padding float64 // padding from edges
}

// NewHBoxLayout creates an HBoxLayout with default spacing and padding.
func NewHBoxLayout() *HBoxLayout {
	return &HBoxLayout{Spacing: 4, Padding: 4}
}

// positionable is used by layouts to position children. Node2D, PanelNode, TextNode, etc. implement this.
type positionable interface {
	SetPosition(x, y float64)
	GetPosition() math2D.Vector2D
}

// Apply positions children horizontally.
// Uses SizeProvider for width when available; otherwise advances by Spacing.
func (h *HBoxLayout) Apply(parent ebiten_extended.SceneNode) {
	children := parent.GetChildren()
	if len(children) == 0 {
		return
	}
	parentPos := math2D.ZeroVector2D()
	if tr, ok := parent.(positionable); ok {
		parentPos = tr.GetPosition()
	}
	x := parentPos.X() + h.Padding
	for _, child := range children {
		if pos, ok := child.(positionable); ok {
			p := pos.GetPosition()
			pos.SetPosition(x, p.Y())
			w := h.Spacing
			if sp, ok := child.(SizeProvider); ok {
				w = sp.GetWidth() + h.Spacing
			}
			x += w
		}
	}
}

// VBoxLayout arranges children in a vertical column (top to bottom).
type VBoxLayout struct {
	Spacing float64
	Padding float64
}

// NewVBoxLayout creates a VBoxLayout with default spacing and padding.
func NewVBoxLayout() *VBoxLayout {
	return &VBoxLayout{Spacing: 4, Padding: 4}
}

// Apply positions children vertically.
// Uses SizeProvider for height when available; otherwise advances by Spacing.
func (v *VBoxLayout) Apply(parent ebiten_extended.SceneNode) {
	children := parent.GetChildren()
	if len(children) == 0 {
		return
	}
	parentPos := math2D.ZeroVector2D()
	if tr, ok := parent.(positionable); ok {
		parentPos = tr.GetPosition()
	}
	y := parentPos.Y() + v.Padding
	for _, child := range children {
		if pos, ok := child.(positionable); ok {
			p := pos.GetPosition()
			pos.SetPosition(p.X(), y)
			th := v.Spacing
			if sp, ok := child.(SizeProvider); ok {
				th = sp.GetHeight() + v.Spacing
			}
			y += th
		}
	}
}

// SizeProvider is an optional interface for nodes that report their size for layout.
// Implement this for accurate spacing in HBoxLayout and VBoxLayout.
type SizeProvider interface {
	GetWidth() float64
	GetHeight() float64
}

