package tilemap

import (
	"fmt"
	"image"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

// Tileset wraps a parsed Tiled tileset, holding its images, tile properties, and animations.
type Tileset struct {
	tiledData   *tiled.Tileset
	Texture     *ebiten.Image
	TileImages  map[uint32]*ebiten.Image // Used for image collection tilesets
	Properties  map[uint32]tiled.Properties
	Animations  map[uint32][]*tiled.AnimationFrame
	ObjectGroup map[uint32][]*tiled.ObjectGroup // For tile collisions
}

// NewTileset creates and initializes a Tileset from go-tiled data and loads its images.
func NewTileset(baseDir string, ts *tiled.Tileset) (*Tileset, error) {
	parsed := &Tileset{
		tiledData:   ts,
		TileImages:  make(map[uint32]*ebiten.Image),
		Properties:  make(map[uint32]tiled.Properties),
		Animations:  make(map[uint32][]*tiled.AnimationFrame),
		ObjectGroup: make(map[uint32][]*tiled.ObjectGroup),
	}

	// 1. Load single texture tileset if present
	if ts.Image != nil && ts.Image.Source != "" {
		imgPath := filepath.Join(baseDir, ts.Image.Source)
		ebImg, _, err := ebitenutil.NewImageFromFile(imgPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load single tileset image %s: %w", imgPath, err)
		}
		parsed.Texture = ebImg
	}

	// 2. Load individual tiles & process extended tile properties
	for _, t := range ts.Tiles {
		// Individual Tile Images (Image Collection tileset)
		if t.Image != nil && t.Image.Source != "" {
			imgPath := filepath.Join(baseDir, t.Image.Source)
			ebImg, _, err := ebitenutil.NewImageFromFile(imgPath)
			if err != nil {
				fmt.Printf("Warning: failed to load individual tile image %s\n", imgPath)
			} else {
				parsed.TileImages[t.ID] = ebImg
			}
		}

		// Save custom properties for later querying
		if len(t.Properties) > 0 {
			parsed.Properties[t.ID] = t.Properties
		}

		// Animations
		if len(t.Animation) > 0 {
			parsed.Animations[t.ID] = t.Animation
		}

		// Collision objects (Object Group)
		if len(t.ObjectGroups) > 0 {
			parsed.ObjectGroup[t.ID] = t.ObjectGroups
		}
	}

	return parsed, nil
}

// GetTileSourceRect calculates the source image.Rectangle for a given tile ID within this tileset.
// For image collections, it simply returns the bounds of the individual image.
func (ts *Tileset) GetTileSourceRect(localID uint32) (image.Rectangle, *ebiten.Image) {
	if ts.Texture != nil {
		cols := ts.tiledData.Columns
		if cols == 0 {
			// Fallback if cols is zero
			cols = ts.tiledData.Image.Width / ts.tiledData.TileWidth
		}

		row := int(localID) / cols
		col := int(localID) % cols

		sx := ts.tiledData.Margin + col*(ts.tiledData.TileWidth+ts.tiledData.Spacing)
		sy := ts.tiledData.Margin + row*(ts.tiledData.TileHeight+ts.tiledData.Spacing)

		return image.Rect(sx, sy, sx+ts.tiledData.TileWidth, sy+ts.tiledData.TileHeight), ts.Texture
	}

	// It's an image collection tileset
	if img, ok := ts.TileImages[localID]; ok {
		return img.Bounds(), img
	}

	return image.Rectangle{}, nil
}

// GetTileProperties returns the Tiled custom properties assigned to the given local tile ID.
func (ts *Tileset) GetTileProperties(localID uint32) tiled.Properties {
	return ts.Properties[localID]
}

// GetTileAnimation returns the animation sequence, if any, for the given local tile ID.
func (ts *Tileset) GetTileAnimation(localID uint32) []*tiled.AnimationFrame {
	return ts.Animations[localID]
}

// GetTileCollisionGroups returns the ObjectGroups containing collision shapes for the given local tile ID.
func (ts *Tileset) GetTileCollisionGroups(localID uint32) []*tiled.ObjectGroup {
	return ts.ObjectGroup[localID]
}
