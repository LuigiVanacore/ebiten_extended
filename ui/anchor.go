package ui

import (
	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/math2d"
)

// AnchorType defines how a UI element is positioned relative to its parent's bounds.
type AnchorType int

const (
	AnchorNone AnchorType = iota
	AnchorTopLeft
	AnchorTopCenter
	AnchorTopRight
	AnchorCenterLeft
	AnchorCenter
	AnchorCenterRight
	AnchorBottomLeft
	AnchorBottomCenter
	AnchorBottomRight
	AnchorStretch
)

// Anchorable is an interface for UI elements that can be positioned by an AnchorLayout.
type Anchorable interface {
	GetAnchor() AnchorType
	SetAnchor(a AnchorType)
	GetAnchorMargin() math2d.Vector2D
	SetAnchorMargin(m math2d.Vector2D)
}

// AnchorLayout positions children based on their Anchor properties.
// The parent node must implement SizeProvider for this to work correctly.
type AnchorLayout struct{}

// NewAnchorLayout creates a new layout that adheres to Anchor properties.
func NewAnchorLayout() *AnchorLayout {
	return &AnchorLayout{}
}

// Apply positions children based on their Anchor property.
func (a *AnchorLayout) Apply(parent ludum.SceneNode) {
	parentSP, ok := parent.(SizeProvider)
	if !ok {
		return // Parent needs to have a size to anchor inside it
	}
	pw, ph := parentSP.GetWidth(), parentSP.GetHeight()

	parentPos := math2d.ZeroVector2D()
	if tr, ok := parent.(positionable); ok {
		parentPos = tr.GetPosition()
	}

	for _, child := range parent.GetChildren() {
		sp, isSp := child.(SizeProvider)
		anchorable, isAnchorable := child.(Anchorable)
		pos, isPos := child.(positionable)

		if !isAnchorable || !isPos {
			continue
		}

		anchor := anchorable.GetAnchor()
		if anchor == AnchorNone {
			continue
		}

		margin := anchorable.GetAnchorMargin()
		cw, ch := 0.0, 0.0
		if isSp {
			cw, ch = sp.GetWidth(), sp.GetHeight()
		}

		var x, y float64

		switch anchor {
		case AnchorTopLeft:
			x = margin.X()
			y = margin.Y()
		case AnchorTopCenter:
			x = (pw - cw) / 2
			y = margin.Y()
		case AnchorTopRight:
			x = pw - cw - margin.X()
			y = margin.Y()
		case AnchorCenterLeft:
			x = margin.X()
			y = (ph - ch) / 2
		case AnchorCenter:
			x = (pw - cw) / 2
			y = (ph - ch) / 2
		case AnchorCenterRight:
			x = pw - cw - margin.X()
			y = (ph - ch) / 2
		case AnchorBottomLeft:
			x = margin.X()
			y = ph - ch - margin.Y()
		case AnchorBottomCenter:
			x = (pw - cw) / 2
			y = ph - ch - margin.Y()
		case AnchorBottomRight:
			x = pw - cw - margin.X()
			y = ph - ch - margin.Y()
		case AnchorStretch:
			x = margin.X()
			y = margin.Y()
			if stretchable, ok := child.(interface{ SetSize(w, h float64) }); ok {
				stretchable.SetSize(pw-margin.X()*2, ph-margin.Y()*2)
				if isSp {
					cw, ch = sp.GetWidth(), sp.GetHeight()
				}
			}
		}

		// Convert local to parent position
		pos.SetPosition(parentPos.X()+x, parentPos.Y()+y)
	}
}
