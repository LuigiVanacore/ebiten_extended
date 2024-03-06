package manager

import (
	"image/color"
	"log"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)


var textManager_instance *textManager

func TextManager() *textManager {
	if textManager_instance == nil {
		panic("TextManager is not initialized correctly")
	}

	return textManager_instance
}

type textManager struct {
	defaultFont font.Face
	defaultFontSize uint
	defualtFontDPI uint
	textMessageList []*textMessage

}

type textMessage struct {
	message string
	position math2D.Vector2D
	color color.Color
}

func newTextMessage(message string, position math2D.Vector2D, color color.Color) *textMessage {
	return &textMessage{ message: message, position: position, color: color}
}

func InitTextManager()  {
	textManager_instance = &textManager{ defaultFontSize: 24, defualtFontDPI: 72, textMessageList:  make([]*textMessage, 0)}
	textManager_instance.loadDefaultFont()
}

func (t *textManager) loadDefaultFont() {
		tt, err := opentype.Parse(resources.DefaultFont)
		if err != nil {
			log.Fatal(err)
		}
	
		gamefont, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:   float64(t.defaultFontSize) ,
			DPI:    float64(t.defualtFontDPI),
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal(err)
		}
		t.defaultFont = gamefont
}

func (t *textManager) WriteText(message string, position math2D.Vector2D, color color.Color) {
	t.textMessageList = append(t.textMessageList, newTextMessage(message, position, color))
}

func (t *textManager) Draw(target *ebiten.Image) {
	for _, textMessage := range t.textMessageList {
		text.Draw(target, textMessage.message, t.defaultFont, int(textMessage.position.X()), int(textMessage.position.Y()),textMessage.color)
	}
}