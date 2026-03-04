package ebiten_extended

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// TextNode represents a visual 2D scene graph node dedicated to drawing scalable geometry-based text phrases.
type TextNode struct {
	Node2D
	message string
	color   color.Color
	font    text.Face
	layer   int
}

// NewTextNode instantiates a display entity resolving specific text output using the assigned typography face format.
func NewTextNode(name string, message string, font text.Face, c color.Color) *TextNode {
	label := &TextNode{message: message, Node2D: *NewNode2D(name), font: font, color: c}
	return label
}

// SetMessage dynamically overrides the actively drawn text string maintained by this graph element.
func (l *TextNode) SetMessage(message string) {
	l.message = message
}

// GetMessage returns the current text.
func (l *TextNode) GetMessage() string {
	return l.message
}

// SetFont sets the font face used for drawing.
func (l *TextNode) SetFont(face text.Face) {
	l.font = face
}

// GetFont returns the font face used for drawing.
func (l *TextNode) GetFont() text.Face {
	return l.font
}

// SetColor sets the text color.
func (l *TextNode) SetColor(c color.Color) {
	l.color = c
}

// GetColor returns the text color.
func (l *TextNode) GetColor() color.Color {
	return l.color
}

// GetLayer returns the render layer of this text node.
func (l *TextNode) GetLayer() int {
	return l.layer
}

// SetLayer sets the render layer of this text node.
func (l *TextNode) SetLayer(layer int) {
	l.layer = layer
}

// Draw translates the text formatting onto the assigned canvas plane mapping it through inherent node translation attributes.
func (l *TextNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if l.font == nil || l.message == "" {
		return
	}

	textOp := &text.DrawOptions{}
	if op != nil {
		textOp.GeoM = op.GeoM
	} else {
		worldPos := l.GetWorldPosition()
		textOp.GeoM.Translate(worldPos.X(), worldPos.Y())
	}
	textOp.ColorScale.ScaleWithColor(l.color)

	text.Draw(target, l.message, l.font, textOp)
}
