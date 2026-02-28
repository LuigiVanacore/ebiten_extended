package ebiten_extended

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)



type ResourceManager struct {
	images map[string]*ebiten.Image
	fonts  []*font.Face
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		images: make(map[string]*ebiten.Image),
	}
}

func (r *ResourceManager) GetImages() map[string]*ebiten.Image {
	return r.images
}

func (r *ResourceManager) GetImage(textureId string) *ebiten.Image {
	return r.images[textureId]
}

func (r *ResourceManager) GetFont(fontId uint) *font.Face {
	if fontId >= uint(len(r.fonts)) {
		return nil
	}
	return r.fonts[fontId]
}

func (r *ResourceManager) AddImage(textureId string, texture []byte) {
	img, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage := ebiten.NewImageFromImage(img)
	r.images[textureId] = ebitenImage
}

// LoadFont loads a font from bytes and adds it to the manager.
// Returns the font index to use with GetFont.
func (r *ResourceManager) LoadFont(f []byte, fontSize float64, dpi float64) (uint, error) {
	tt, err := opentype.Parse(f)
	if err != nil {
		return 0, err
	}
	gamefont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return 0, err
	}
	r.fonts = append(r.fonts, &gamefont)
	return uint(len(r.fonts) - 1), nil
}
