package ludum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// ResourceManager organizes and caches loaded game assets such as images and fonts safely.
type ResourceManager struct {
	mu     sync.RWMutex
	images map[string]*ebiten.Image
	fonts  map[string]text.Face
}

// NewResourceManager creates an empty ResourceManager instance ready for asset loading.
func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		images: make(map[string]*ebiten.Image),
		fonts:  make(map[string]text.Face),
	}
}

// GetImages returns a copy of the cached ebiten images mapping.
// Modifying the returned map does not affect the ResourceManager.
func (r *ResourceManager) GetImages() map[string]*ebiten.Image {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[string]*ebiten.Image, len(r.images))
	for k, v := range r.images {
		out[k] = v
	}
	return out
}

// GetImage fetches a specific loaded ebiten image by its arbitrary string identifier.
func (r *ResourceManager) GetImage(textureId string) *ebiten.Image {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.images[textureId]
}

// RemoveImage removes a cached image by its texture ID.
// Returns true if the image existed and was removed.
func (r *ResourceManager) RemoveImage(textureId string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.images[textureId]; !ok {
		return false
	}
	delete(r.images, textureId)
	return true
}

// ClearImages removes all cached images.
// Returns the number of images removed.
func (r *ResourceManager) ClearImages() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	n := len(r.images)
	r.images = make(map[string]*ebiten.Image)
	return n
}

// GetFonts returns a copy of the cached fonts mapping.
// Modifying the returned map does not affect the ResourceManager.
func (r *ResourceManager) GetFonts() map[string]text.Face {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[string]text.Face, len(r.fonts))
	for k, v := range r.fonts {
		out[k] = v
	}
	return out
}

// GetFont provides access to a loaded text face via its string ID.
func (r *ResourceManager) GetFont(fontId string) text.Face {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.fonts[fontId]
}

// RemoveFont removes a cached font by its string ID.
// Returns true if the font existed and was removed.
func (r *ResourceManager) RemoveFont(fontId string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.fonts[fontId]; !ok {
		return false
	}
	delete(r.fonts, fontId)
	return true
}

// ClearFonts removes all loaded fonts.
// Returns the number of fonts removed.
func (r *ResourceManager) ClearFonts() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	n := len(r.fonts)
	r.fonts = make(map[string]text.Face)
	return n
}

// Clear removes all cached assets (images and fonts).
func (r *ResourceManager) Clear() {
	r.ClearImages()
	r.ClearFonts()
}

// AssetType enumerates the types of loadable assets for batching.
type AssetType string

const (
	AssetTypeImage AssetType = "image"
	AssetTypeFont  AssetType = "font"
	AssetTypeAtlas AssetType = "atlas"
)

// PreloadAsset defines a single asset descriptor for batch loading.
type PreloadAsset struct {
	ID   string
	Path string
	Type AssetType
	// Only for AssetTypeFont:
	FontSize float64
	// Only for AssetTypeAtlas:
	JsonPath string
}

// PreloadBatch asynchronously loads a list of assets and triggers progress callbacks.
// The onProgress callback receives the number of loaded assets, total assets, and ID of the last one.
// The onComplete callback triggers when the entire batch is finished. Errors are logged but do not interrupt the batch.
func (r *ResourceManager) PreloadBatch(
	assets []PreloadAsset,
	onProgress func(loaded, total int, lastLoadedID string),
	onComplete func(failedAssets []string),
) {
	go func() {
		total := len(assets)
		var failed []string

		for i, asset := range assets {
			var err error
			switch asset.Type {
			case AssetTypeImage:
				err = r.LoadImageFromFile(asset.ID, asset.Path)
			case AssetTypeFont:
				err = r.LoadFontFromFile(asset.ID, asset.Path, asset.FontSize)
			case AssetTypeAtlas:
				err = r.LoadAtlas(asset.ID, asset.Path, asset.JsonPath)
			default:
				err = fmt.Errorf("unknown asset type: %s", asset.Type)
			}

			if err != nil {
				failed = append(failed, asset.ID)
			}

			if onProgress != nil {
				onProgress(i+1, total, asset.ID)
			}
		}

		if onComplete != nil {
			onComplete(failed)
		}
	}()
}

// LoadImageFromFile loads an image from the given file path and caches it under textureId.
// Supports common formats (PNG, JPG, GIF, etc.). Returns an error if the file cannot be read or decoded.
func (r *ResourceManager) LoadImageFromFile(textureId string, path string) error {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.images[textureId] = img
	return nil
}

// AddImage decodes raw image bytes, converts them into an ebiten Image, and binds them to the provided textureId.
// Returns an error if decoding fails.
func (r *ResourceManager) AddImage(textureId string, texture []byte) error {
	img, _, err := image.Decode(bytes.NewReader(texture))
	if err != nil {
		return err
	}
	ebitenImage := ebiten.NewImageFromImage(img)

	r.mu.Lock()
	defer r.mu.Unlock()
	r.images[textureId] = ebitenImage
	return nil
}

// AtlasRegion defines the bounding box of a sub-image inside an atlas.
type AtlasRegion struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// atlasData is used to decode TexturePacker Hash format JSONs.
type atlasData struct {
	Frames map[string]struct {
		Frame AtlasRegion `json:"frame"`
	} `json:"frames"`
}

// LoadAtlas loads a texture atlas image and its associated JSON descriptor.
// The JSON must match the TexturePacker "JSON (Hash)" format with "frames" containing "frame" definitions.
// Each region will be available via GetAtlasRegion(atlasId, regionId) or GetImage(atlasId + ":" + regionId).
func (r *ResourceManager) LoadAtlas(atlasId string, imgPath string, jsonPath string) error {
	// First load the actual image. We temporarily use atlasId for it.
	err := r.LoadImageFromFile(atlasId, imgPath)
	if err != nil {
		return err
	}

	// Now read the JSON descriptor.
	jsonData, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}

	var data atlasData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	baseImage := r.images[atlasId]
	if baseImage == nil {
		return fmt.Errorf("base image %s not found after loading", atlasId)
	}

	for regionId, frameData := range data.Frames {
		f := frameData.Frame
		rect := image.Rect(f.X, f.Y, f.X+f.W, f.Y+f.H)
		subImg := baseImage.SubImage(rect).(*ebiten.Image)
		r.images[atlasId+":"+regionId] = subImg
	}

	return nil
}

// GetAtlasRegion is a convenience method to fetch a sub-image loaded via LoadAtlas.
// It is equivalent to calling GetImage(atlasId + ":" + regionId).
func (r *ResourceManager) GetAtlasRegion(atlasId, regionId string) *ebiten.Image {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.images[atlasId+":"+regionId]
}

// LoadFontFromFile loads an OpenType font from the given file path and caches it under fontId.
// Returns an error if the file cannot be read or decoded.
func (r *ResourceManager) LoadFontFromFile(fontId string, path string, fontSize float64) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return r.LoadFont(fontId, data, fontSize)
}

// LoadFont loads an OpenType font from a reader and adds it to the manager.
// Returns an error if decoding fails.
func (r *ResourceManager) LoadFont(fontId string, f []byte, fontSize float64) error {
	reader := bytes.NewReader(f)

	source, err := text.NewGoTextFaceSource(reader)
	if err != nil {
		return err
	}

	face := &text.GoTextFace{
		Source: source,
		Size:   fontSize,
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.fonts[fontId] = face
	return nil
}
