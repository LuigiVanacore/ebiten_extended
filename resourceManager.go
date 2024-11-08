package ebiten_extended

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)



var instance *resourceManager

func ResourceManager() *resourceManager {
	if instance == nil {
		instance = newResourceManager()
	}
	return instance
}

type resourceManager struct {
	images map[string]*ebiten.Image
	fonts  []*font.Face
}

func newResourceManager() *resourceManager {
	return &resourceManager{
		images: make(map[string]*ebiten.Image),
	}
}

func (r *resourceManager) GetImages() map[string]*ebiten.Image {
	return r.images
}

func (r *resourceManager) GetImage(textureId string) *ebiten.Image {
	return r.images[textureId]
}

func (r *resourceManager) GetFont(fontId uint) *font.Face {
	return r.fonts[fontId]
}

func (r *resourceManager) AddImage(textureId string, texture []byte) {
	img, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage := ebiten.NewImageFromImage(img)
	r.images[textureId] = ebitenImage
}

func (r *resourceManager) LoadFont(f []byte, fontSize float64, dpi float64) {
	tt, err := opentype.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	gamefont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	r.fonts = append(r.fonts, &gamefont)
}
