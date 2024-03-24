package ebiten_extended

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type LabelText struct {
	BaseNode
	message  string
	color color.Color
	font  font.Face
}

func NewLabelText(name string, message string, position math2D.Vector2D, font font.Face, color color.Color) *LabelText {
	return &LabelText{message: message, BaseNode: *NewBaseNode(name),  font: font, color: color}
}


func (l *LabelText) SetMessage(message string) {
	l.message = message
}

func (l *LabelText) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	text.Draw(target, l.message, l.font, int(l.transform.GetPosition().X()), int(l.transform.GetPosition().Y()), l.color)
}