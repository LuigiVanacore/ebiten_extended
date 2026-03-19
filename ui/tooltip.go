package ui

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Hoverable is implemented by nodes that can report whether the cursor is over them.
type Hoverable interface {
	ContainsPoint(x, y float64) bool
}

// TooltipNode shows a tooltip when the user hovers over its target for a short delay.
// Add a single child that implements Hoverable (e.g. ButtonNode, PanelNode).
type TooltipNode struct {
	ebiten_extended.Node2D
	InputManager *input.InputManager
	target       Hoverable
	text         string
	delay        float64
	hoverTime    float64
	visible      bool

	font       text.Face
	textColor  color.Color
	bgColor    color.Color
	padding    float64
	offsetX    float64
	offsetY    float64
}

// NewTooltipNode creates a tooltip that activates when hovering over the target.
// The target is typically the first child added with AddChildren.
func NewTooltipNode(name string, target Hoverable, tooltipText string, font text.Face, im *input.InputManager) *TooltipNode {
	n := &TooltipNode{
		Node2D:      *ebiten_extended.NewNode2D(name),
		InputManager: im,
		target:     target,
		text:       tooltipText,
		delay:      0.5,
		font:       font,
		textColor:  color.White,
		bgColor:    color.RGBA{30, 30, 30, 230},
		padding:    6,
		offsetX:    12,
		offsetY:    8,
	}
	return n
}

// SetText sets the tooltip text.
func (t *TooltipNode) SetText(s string) {
	t.text = s
}

// SetDelay sets the hover delay in seconds before the tooltip appears.
func (t *TooltipNode) SetDelay(sec float64) {
	t.delay = sec
}

// SetTarget sets the hoverable target (optional if set via constructor).
func (t *TooltipNode) SetTarget(target Hoverable) {
	t.target = target
}

// SetStyle sets text color, background color, and padding.
func (t *TooltipNode) SetStyle(textColor, bgColor color.Color, padding float64) {
	t.textColor = textColor
	t.bgColor = bgColor
	t.padding = padding
}

// SetOffset sets the offset from the cursor when displaying the tooltip.
func (t *TooltipNode) SetOffset(x, y float64) {
	t.offsetX, t.offsetY = x, y
}

// GetWidth returns the tooltip box width (text width + padding). Implements SizeProvider for layout.
// Returns 0 if font is nil or text is empty.
func (t *TooltipNode) GetWidth() float64 {
	if t.font == nil || t.text == "" {
		return 0
	}
	w, _ := text.Measure(t.text, t.font, 0)
	return w + t.padding*2
}

// GetHeight returns the tooltip box height (text height + padding). Implements SizeProvider for layout.
// Returns 0 if font is nil or text is empty.
func (t *TooltipNode) GetHeight() float64 {
	if t.font == nil || t.text == "" {
		return 0
	}
	_, h := text.Measure(t.text, t.font, 0)
	return h + t.padding*2
}

// Update advances the hover timer and toggles visibility.
func (t *TooltipNode) Update() {
	if t.InputManager == nil || t.target == nil || t.text == "" {
		t.visible = false
		t.hoverTime = 0
		return
	}
	cx := t.InputManager.GetCursorPos().X()
	cy := t.InputManager.GetCursorPos().Y()
	if t.target.ContainsPoint(cx, cy) {
		t.hoverTime += ebiten_extended.FIXED_DELTA
		if t.hoverTime >= t.delay {
			t.visible = true
		}
	} else {
		t.visible = false
		t.hoverTime = 0
	}
}

// Draw renders the tooltip when visible.
func (t *TooltipNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if !t.visible || t.font == nil {
		return
	}
	cx := t.InputManager.GetCursorPos().X()
	cy := t.InputManager.GetCursorPos().Y()
	w, h := text.Measure(t.text, t.font, 0)
	boxW := w + t.padding*2
	boxH := h + t.padding*2
	boxX := cx + t.offsetX
	boxY := cy + t.offsetY
	screenW := float64(target.Bounds().Dx())
	screenH := float64(target.Bounds().Dy())
	if boxX+boxW > screenW {
		boxX = cx - t.offsetX - boxW
	}
	if boxY+boxH > screenH {
		boxY = cy - t.offsetY - boxH
	}
	if boxX < 0 {
		boxX = 0
	}
	if boxY < 0 {
		boxY = 0
	}
	if t.bgColor != nil {
		vector.DrawFilledRect(target, float32(boxX), float32(boxY), float32(boxW), float32(boxH), t.bgColor, true)
	}
	textOpts := &text.DrawOptions{}
	textOpts.GeoM.Translate(boxX+t.padding, boxY+t.padding)
	textOpts.ColorScale.ScaleWithColor(t.textColor)
	text.Draw(target, t.text, t.font, textOpts)
}

// Ensure TooltipNode implements Updatable.
var _ ebiten_extended.Updatable = (*TooltipNode)(nil)
