package ebiten_extended

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type LabelText struct {
	Node2D
	message string
	color   color.Color
	font    font.Face
}

func NewLabelText(name string, message string, position math2D.Vector2D, font font.Face, color color.Color) *LabelText {
	label := &LabelText{message: message, Node2D: *NewNode2D(name), font: font, color: color}
	label.SetPosition(position.X(), position.Y())
	return label
}

func (l *LabelText) SetMessage(message string) {
	l.message = message
}

func (l *LabelText) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	text.Draw(target, l.message, l.font, int(l.transform.GetPosition().X()), int(l.transform.GetPosition().Y()), l.color)
}
