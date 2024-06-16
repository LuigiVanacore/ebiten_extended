package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 320
	screenHeight = 240
	defaultFontSize = 24
	defualtFontDPI = 72
)

type Game struct {
	counter        int
	textLabel *ebiten_extended.LabelText
}

func NewGame() *Game {
	gameFont := loadDefaultFont()
	textLabel := ebiten_extended.NewLabelText("labelTest", "test label text", math2D.NewVector2D(0,90), gameFont, color.White)
	ebiten_extended.SceneManager().AddSceneNodeToDefaultLayer(textLabel)
	return &Game{ textLabel: textLabel}
}

func (g *Game) Update() error {
	if g.counter%ebiten.TPS() == 0 {
		g.counter=0
	}
	message := "test label text " + strconv.Itoa(g.counter)
	g.textLabel.SetMessage(message)
	g.counter++
	ebiten_extended.GameManager().Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebiten_extended.GameManager().Draw(screen, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenHeight, screenHeight
}

func loadDefaultFont() font.Face {
	tt, err := opentype.Parse(resources.DefaultFont)
	if err != nil {
		log.Fatal(err)
	}

	gamefont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:   float64(defaultFontSize) ,
		DPI:    float64(defualtFontDPI),
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return gamefont
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Animation (Ebitengine Demo)")

	
	
	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}