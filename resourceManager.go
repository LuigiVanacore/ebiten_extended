package ebiten_extended

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ResourceManager organizes and caches loaded game assets such as images and fonts.
type ResourceManager struct {
	images map[string]*ebiten.Image
	fonts  []text.Face
}

// NewResourceManager creates an empty ResourceManager instance ready for asset loading.
func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		images: make(map[string]*ebiten.Image),
		animations: make(map[string]*AnimationSet),
	}
}

// GetImages retrieves the complete underlying dictionary mapping of cached ebiten images.
func (r *ResourceManager) GetImages() map[string]*ebiten.Image {
	return r.images
}

// GetImage fetches a specific loaded ebiten image by its arbitrary string identifier.
func (r *ResourceManager) GetImage(textureId string) *ebiten.Image {
	return r.images[textureId]
}

// GetFont provides access to a loaded text face via its sequentially assigned integer ID.
func (r *ResourceManager) GetFont(fontId uint) text.Face {
	if fontId >= uint(len(r.fonts)) {
		return nil
	}
	return r.fonts[fontId]
}

// AddImage decodes raw image bytes, converts them into an ebiten Image, and binds them to the provided textureId.
func (r *ResourceManager) AddImage(textureId string, texture []byte) {
	img, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		log.Fatal(err)
	}
	ebitenImage := ebiten.NewImageFromImage(img)
	r.images[textureId] = ebitenImage
}

// LoadFont loads an OpenType font from a reader and adds it to the manager.
// Returns the font index to use with GetFont.
func (r *ResourceManager) LoadFont(f []byte, fontSize float64) (uint, error) {
	reader := bytes.NewReader(f)

	source, err := text.NewGoTextFaceSource(reader)
	if err != nil {
		return 0, err
	}

	face := &text.GoTextFace{
		Source: source,
		Size:   fontSize,
	}

	r.fonts = append(r.fonts, face)
	return uint(len(r.fonts) - 1), nil
}
