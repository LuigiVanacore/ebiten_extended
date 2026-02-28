package ebiten_extended

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextNode struct {
	Node2D
	message string
	color   color.Color
	font    text.Face
}

func NewTextNode(name string, message string, font text.Face, c color.Color) *TextNode {
	label := &TextNode{message: message, Node2D: *NewNode2D(name), font: font, color: c}
	return label
}

func (l *TextNode) SetMessage(message string) {
	l.message = message
}

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
