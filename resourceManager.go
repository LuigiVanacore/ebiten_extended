package ebiten_extended

import (
	"bytes"
	"image"

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

// RemoveImage removes a cached image by its texture ID.
// Returns true if the image existed and was removed.
func (r *ResourceManager) RemoveImage(textureId string) bool {
	if _, ok := r.images[textureId]; !ok {
		return false
	}
	delete(r.images, textureId)
	return true
}

// ClearImages removes all cached images.
// Returns the number of images removed.
func (r *ResourceManager) ClearImages() int {
	n := len(r.images)
	r.images = make(map[string]*ebiten.Image)
	return n
}

// GetFont provides access to a loaded text face via its sequentially assigned integer ID.
func (r *ResourceManager) GetFont(fontId uint) text.Face {
	if fontId >= uint(len(r.fonts)) {
		return nil
	}
	return r.fonts[fontId]
}

// ClearFonts removes all loaded fonts.
// Returns the number of fonts removed.
func (r *ResourceManager) ClearFonts() int {
	n := len(r.fonts)
	r.fonts = nil
	return n
}

// Clear removes all cached assets (images and fonts).
func (r *ResourceManager) Clear() {
	r.ClearImages()
	r.ClearFonts()
}

// AddImage decodes raw image bytes, converts them into an ebiten Image, and binds them to the provided textureId.
// Returns an error if decoding fails.
func (r *ResourceManager) AddImage(textureId string, texture []byte) error {
	img, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		return err
	}
	ebitenImage := ebiten.NewImageFromImage(img)
	r.images[textureId] = ebitenImage
	return nil
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
