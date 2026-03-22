package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ludum"
	"github.com/LuigiVanacore/ludum/input"
	"github.com/hajimehoshi/ebiten/v2"
)

// ScrollPanelNode is a panel that clips its content and supports vertical scrolling.
// Add a single content node as child; its local position is offset by (-scrollX, -scrollY).
// Mouse wheel scrolls vertically. SetContentSize to the total content height for bounds.
type ScrollPanelNode struct {
	PanelNode
	InputManager *input.InputManager
	contentW     float64
	contentH     float64
	scrollX      float64
	scrollY      float64
	scrollSpeed  float64
}

// NewScrollPanelNode creates a scroll panel with the given viewport size.
func NewScrollPanelNode(name string, viewportW, viewportH float64, im *input.InputManager) *ScrollPanelNode {
	p := NewPanelNode(name, viewportW, viewportH)
	p.SetBackgroundColor(color.RGBA{30, 30, 30, 255})
	return &ScrollPanelNode{
		PanelNode:    *p,
		InputManager: im,
		contentW:     viewportW,
		contentH:     viewportH,
		scrollSpeed:  30,
	}
}

// SetContentSize sets the total scrollable content size. Scroll is clamped to keep content in view.
func (s *ScrollPanelNode) SetContentSize(w, h float64) {
	s.contentW, s.contentH = w, h
	s.clampScroll()
}

// SetScroll sets the scroll offset.
func (s *ScrollPanelNode) SetScroll(x, y float64) {
	s.scrollX, s.scrollY = x, y
	s.clampScroll()
}

// GetScroll returns the current scroll offset.
func (s *ScrollPanelNode) GetScroll() (float64, float64) {
	return s.scrollX, s.scrollY
}

func (s *ScrollPanelNode) clampScroll() {
	if s.contentW < s.width {
		s.scrollX = 0
	} else if s.scrollX < 0 {
		s.scrollX = 0
	} else if s.scrollX > s.contentW-s.width {
		s.scrollX = s.contentW - s.width
	}
	if s.contentH < s.height {
		s.scrollY = 0
	} else if s.scrollY < 0 {
		s.scrollY = 0
	} else if s.scrollY > s.contentH-s.height {
		s.scrollY = s.contentH - s.height
	}
}

// Update handles mouse wheel scrolling and updates content offset.
func (s *ScrollPanelNode) Update() {
	if s.InputManager != nil {
		_, dy := ebiten.Wheel()
		if dy != 0 {
			s.scrollY -= dy * s.scrollSpeed
			s.clampScroll()
		}
	}
	for _, ch := range s.GetChildren() {
		if n, ok := ch.(*ludum.Node2D); ok {
			n.SetPosition(-s.scrollX, -s.scrollY)
			break
		}
	}
}

// Draw renders the panel background, then draws children.
// Content is positioned at (-scrollX, -scrollY) via the child's local position.
// Note: overflow is not clipped; ensure content size matches or use opaque background to hide.
func (s *ScrollPanelNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.PanelNode.Draw(target, op)
	// Children are drawn by World with their transform; we set their position in Update
}
