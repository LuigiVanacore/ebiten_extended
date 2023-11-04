package ebiten_extended

import (
	"image/color"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)


var textManager_instance *textManager

func TextManager() *textManager {
	if instance == nil {
		panic("TextManager is not initialized correctly")
	}

	return textManager_instance
}

type textManager struct {
	dst *ebiten.Image
	defaultFont font.Face

}

func InitTextManager(dst *ebiten.Image)  {
	textManager_instance = &textManager{ dst: dst}
}

func (t *textManager) WriteText(message string, position math2D.Vector2D, color color.Color) {
	text.Draw(t.dst, message, t.defaultFont, int(position.X()), int(position.Y()), color)
}