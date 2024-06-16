package ebiten_extended

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type LabelText struct {
	Node2D
	message string
	color   color.Color
	font    text.Face
}

func NewLabelText(name string, message string, position math2D.Vector2D, font text.Face, color color.Color) *LabelText {
	label := &LabelText{message: message, Node2D: *NewNode2D(name), font: font, color: color}
	label.SetPosition(position.X(), position.Y())
	return label
}

func (l *LabelText) SetMessage(message string) {
	l.message = message
}

func (l *LabelText) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	text_op := &text.DrawOptions{}
	text_op.GeoM.Translate(l.GetPosition().X(), l.GetPosition().Y())
	//text_op.ColorM = ebiten.ColorM{}
	text.Draw(target, l.message, l.font, text_op)
}
