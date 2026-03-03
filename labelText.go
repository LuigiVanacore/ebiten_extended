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

// Draw translates the text formatting onto the assigned canvas plane mapping it through inherent node translation attributes.
func (l *TextNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if l.font == nil || l.message == "" {
		return
	}

	worldPos := l.GetWorldPosition()

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(worldPos.X(), worldPos.Y())
	textOp.ColorScale.ScaleWithColor(l.color)

	text.Draw(target, l.message, l.font, textOp)
}
