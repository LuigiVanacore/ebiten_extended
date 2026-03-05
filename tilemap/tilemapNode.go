package tilemap

import (
	"fmt"
	"image"
	"math"
	"path/filepath"

	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type animatedTileState struct {
	currentFrameIndex int
	timerMillis       float64
}

type tileAnimKey struct {
	tilesetName string
	localID     uint32
}

// TileMapNode wraps a Tiled map (.tmx) into an Ebiten Node2D.
// It loads images directly into Ebiten GPU memory and renders tiles efficiently using batching.
type TileMapNode struct {
	*ebiten_extended.Node2D
	MapData         *tiled.Map
	tilesets        map[string]*Tileset
	layerIndex      int
	animationStates map[tileAnimKey]*animatedTileState
}

// NewTileMapNode creates and initializes a TileMapNode, parsing the .tmx file
// and loading related tileset images automatically.
func NewTileMapNode(tmxPath string) (*TileMapNode, error) {
	parsedMap, err := tiled.LoadFile(tmxPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load TMX file %s: %w", tmxPath, err)
	}

	tm := &TileMapNode{
		Node2D:          ebiten_extended.NewNode2D("TileMapNode"),
		MapData:         parsedMap,
		tilesets:        make(map[string]*Tileset),
		animationStates: make(map[tileAnimKey]*animatedTileState),
	}

	baseDir := filepath.Dir(tmxPath)

	// Parse and store Tilesets using our wrapper
	for _, ts := range parsedMap.Tilesets {
		tilesetWrapper, err := NewTileset(baseDir, ts)
		if err != nil {
			fmt.Printf("Warning: failed to load tileset %s: %v\n", ts.Name, err)
			continue
		}
		// We index by name since Source might be empty for embedded tilesets,
		// but typically go-tiled uses unique names per map.
		tm.tilesets[ts.Name] = tilesetWrapper

		// Pre-populate animation states
		for localID := range tilesetWrapper.Animations {
			key := tileAnimKey{tilesetName: ts.Name, localID: localID}
			tm.animationStates[key] = &animatedTileState{}
		}
	}

	return tm, nil
}

// GetMapData returns the loaded tiled.Map data.
func (t *TileMapNode) GetMapData() *tiled.Map {
	return t.MapData
}

// SetLayer explicitly sets the rendering layer for this tilemap.
func (t *TileMapNode) SetLayer(l int) {
	t.layerIndex = l
}

// GetLayer implements the Drawable interface.
func (t *TileMapNode) GetLayer() int {
	return t.layerIndex
}

// Update implements the Updatable interface to process tile animations.
func (t *TileMapNode) Update() {
	deltaMillis := ebiten_extended.FIXED_DELTA * 1000.0

	for key, state := range t.animationStates {
		wrapper, ok := t.tilesets[key.tilesetName]
		if !ok {
			continue
		}

		frames := wrapper.GetTileAnimation(key.localID)
		if len(frames) == 0 {
			continue
		}

		state.timerMillis += deltaMillis
		currentFrame := frames[state.currentFrameIndex]
		if state.timerMillis >= float64(currentFrame.Duration) {
			state.timerMillis -= float64(currentFrame.Duration)
			state.currentFrameIndex = (state.currentFrameIndex + 1) % len(frames)
		}
	}
}

// BuildCollisions parses the ObjectGroups from the TMX tiles and generates Collider nodes
// as children of the TileMapNode. The specified mask is applied to each generated collider.
func (t *TileMapNode) BuildCollisions(mask collision.CollisionMask) error {
	if t.MapData == nil {
		return nil
	}

	for _, layer := range t.MapData.Layers {
		for y := 0; y < t.MapData.Height; y++ {
			for x := 0; x < t.MapData.Width; x++ {
				idx := y*t.MapData.Width + x
				tile := layer.Tiles[idx]

				if tile.IsNil() || tile.Tileset == nil {
					continue
				}

				ts, ok := t.tilesets[tile.Tileset.Name]
				if !ok {
					continue
				}

				objGroups := ts.GetTileCollisionGroups(tile.ID)
				for _, group := range objGroups {
					for _, obj := range group.Objects {
						var shape collision.CollisionShape

						if len(obj.Ellipses) > 0 {
							// Circle
							radius := math.Max(obj.Width, obj.Height) / 2
							circle := math2D.NewCircle(math2D.NewVector2D(radius, radius), radius) // local center
							shape = collision.NewCollisionCircle(circle)
						} else {
							// Rectangle
							rect := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(obj.Width, obj.Height))
							shape = collision.NewCollisionRect(rect)
						}

						col, err := collision.NewCollider(shape, mask)
						if err != nil {
							continue
						}

						// World coordinates calculate from tile position
						realX := float64(x * t.MapData.TileWidth)
						realY := float64(y * t.MapData.TileHeight)
						col.SetPosition(realX+obj.X, realY+obj.Y)
						col.SetName(fmt.Sprintf("TileCollision_%d_%d", x, y))
						t.AddChildren(col)
					}
				}
			}
		}
	}
	return nil
}

// Draw handles the efficient batch rendering of the tile map.
// Currently renders the whole map using Ebiten's internal batching.
func (t *TileMapNode) Draw(target *ebiten.Image, baseOp *ebiten.DrawImageOptions) {
	if t.MapData == nil {
		return
	}

	geom := ebiten.GeoM{}
	if baseOp != nil {
		geom = baseOp.GeoM
	} else {
		worldTransform := t.GetWorldTransform()
		geom = worldTransform.UpdateGeoM(ebiten.GeoM{})
	}

	for _, layer := range t.MapData.Layers {
		if !layer.Visible {
			continue
		}

		// Calcolo il bounding box visibile sullo schermo invertendo il geom matrix del nodo
		startX, startY := 0, 0
		endX, endY := t.MapData.Width, t.MapData.Height

		if geom.IsInvertible() {
			invGeom := geom
			invGeom.Invert()

			w, h := float64(target.Bounds().Dx()), float64(target.Bounds().Dy())
			x0, y0 := invGeom.Apply(0, 0)
			x1, y1 := invGeom.Apply(w, 0)
			x2, y2 := invGeom.Apply(0, h)
			x3, y3 := invGeom.Apply(w, h)

			minLocalX := math.Min(math.Min(x0, x1), math.Min(x2, x3))
			maxLocalX := math.Max(math.Max(x0, x1), math.Max(x2, x3))
			minLocalY := math.Min(math.Min(y0, y1), math.Min(y2, y3))
			maxLocalY := math.Max(math.Max(y0, y1), math.Max(y2, y3))

			sX := int(math.Floor(minLocalX / float64(t.MapData.TileWidth)))
			eX := int(math.Ceil(maxLocalX / float64(t.MapData.TileWidth)))
			sY := int(math.Floor(minLocalY / float64(t.MapData.TileHeight)))
			eY := int(math.Ceil(maxLocalY / float64(t.MapData.TileHeight)))

			if sX > 0 {
				startX = sX
			}
			if sY > 0 {
				startY = sY
			}
			if eX < t.MapData.Width {
				endX = eX
			}
			if eY < t.MapData.Height {
				endY = eY
			}
		}

		for y := startY; y < endY; y++ {
			for x := startX; x < endX; x++ {
				idx := y*t.MapData.Width + x
				tile := layer.Tiles[idx]

				if tile.IsNil() {
					continue
				}

				realX := float64(x * t.MapData.TileWidth)
				realY := float64(y * t.MapData.TileHeight)

				localID := tile.ID
				ts := tile.Tileset
				if ts == nil {
					continue
				}

				var img *ebiten.Image
				var srcRect image.Rectangle

				// Determine if this tile is animated and fetch current frame localID
				animKey := tileAnimKey{tilesetName: ts.Name, localID: localID}
				if state, ok := t.animationStates[animKey]; ok {
					if wrapper, ok := t.tilesets[ts.Name]; ok {
						frames := wrapper.GetTileAnimation(localID)
						if len(frames) > 0 && state.currentFrameIndex < len(frames) {
							localID = frames[state.currentFrameIndex].TileID
						}
					}
				}

				// Fetch from wrapper Tileset
				if wrapper, ok := t.tilesets[ts.Name]; ok {
					srcRect, img = wrapper.GetTileSourceRect(localID)
				}

				if img == nil {
					continue
				}

				tileOp := &ebiten.DrawImageOptions{}
				if baseOp != nil {
					tileOp.ColorScale = baseOp.ColorScale
					tileOp.Blend = baseOp.Blend
					tileOp.Filter = baseOp.Filter
				}

				// Handle transformations centered around the tile
				subGeo := ebiten.GeoM{}

				// Flips
				if tile.DiagonalFlip {
					// Diagonal flip means transpose (swap x and y)
					// Equivalent to rotation by 90 degrees and purely vertical flip?
					// Actually Tiled diagonal flip maps to Rotate90 + FlipH or similar
					subGeo.Scale(-1, 1)
					subGeo.Rotate(-math.Pi / 2)
				}
				if tile.HorizontalFlip {
					subGeo.Scale(-1, 1)
					subGeo.Translate(float64(ts.TileWidth), 0)
				}
				if tile.VerticalFlip {
					subGeo.Scale(1, -1)
					subGeo.Translate(0, float64(ts.TileHeight))
				}

				// Adjusting anchor if bottom-left oriented?
				// Wait, Tiled draws tiles anchored at bottom-left in some cases, but generally top-left for ortho.
				subGeo.Translate(realX, realY)

				// Combine with node world transform
				subGeo.Concat(geom)

				tileOp.GeoM = subGeo

				// Opacity
				if layer.Opacity < 1.0 {
					tileOp.ColorScale.ScaleAlpha(float32(layer.Opacity))
				}

				// Perform sub-draw using SubImage
				subImg := img.SubImage(srcRect).(*ebiten.Image)
				target.DrawImage(subImg, tileOp)
			}
		}
	}
}
