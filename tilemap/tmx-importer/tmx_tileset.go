package tmx_importer

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/lafriks/go-tiled"
)

var (
	// ErrUnsupportedOrientation represents an error in the unsupported orientation for rendering.
	ErrUnsupportedOrientation = errors.New("tiled/render: unsupported orientation")
	// ErrUnsupportedRenderOrder represents an error in the unsupported order for rendering.
	ErrUnsupportedRenderOrder = errors.New("tiled/render: unsupported render order")

	// ErrOutOfBounds represents an error that the index is out of bounds
	ErrOutOfBounds = errors.New("tiled/render: index out of bounds")
)

// RendererEngine is the interface implemented by objects that provide rendering engine for Tiled maps.
type RendererEngine interface {
	Init(m *tiled.Map)
	GetFinalImageSize() image.Rectangle
	RotateTileImage(tile *tiled.LayerTile, img image.Image) image.Image
	GetTilePosition(x, y int) image.Rectangle
}

// Renderer represents an rendering engine.
type Renderer struct {
	m         *tiled.Map
	Result    *image.NRGBA // The image result after rendering using the Render functions.
	tileCache map[uint32]image.Image
	engine    RendererEngine
	fs        fs.FS
}

// NewRenderer creates new rendering engine instance.
func NewRenderer(m *tiled.Map) (*Renderer, error) {
	return NewRendererWithFileSystem(m, nil)
}

// NewRendererWithFileSystem creates new rendering engine instance with a custom file system.
func NewRendererWithFileSystem(m *tiled.Map, fs fs.FS) (*Renderer, error) {
	r := &Renderer{m: m, tileCache: make(map[uint32]image.Image), fs: fs}
	if r.m.Orientation == "orthogonal" {
		r.engine = &OrthogonalRendererEngine{}
	} else {
		return nil, ErrUnsupportedOrientation
	}

	r.engine.Init(r.m)
	r.Clear()

	return r, nil
}

func (r *Renderer) open(f string) (io.ReadCloser, error) {
	if r.fs == nil {
		return os.Open(filepath.FromSlash(f))
	}
	return r.fs.Open(filepath.ToSlash(f))
}

func (r *Renderer) getTileImage(tile *tiled.LayerTile) (image.Image, error) {
	timg, ok := r.tileCache[tile.Tileset.FirstGID+tile.ID]
	if ok {
		return r.engine.RotateTileImage(tile, timg), nil
	}
	// Precache all tiles in tileset
	if tile.Tileset.Image == nil {
		for i := 0; i < len(tile.Tileset.Tiles); i++ {
			if tile.Tileset.Tiles[i].ID == tile.ID {
				sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Tiles[i].Image.Source))
				if err != nil {
					return nil, err
				}
				defer sf.Close()
				timg, _, err = image.Decode(sf)
				if err != nil {
					return nil, err
				}
				r.tileCache[tile.Tileset.FirstGID+tile.ID] = timg
			}
		}
	} else {
		sf, err := r.open(tile.Tileset.GetFileFullPath(tile.Tileset.Image.Source))
		if err != nil {
			return nil, err
		}
		defer sf.Close()

		img, _, err := image.Decode(sf)
		if err != nil {
			return nil, err
		}

		for i := uint32(0); i < uint32(tile.Tileset.TileCount); i++ {
			rect := tile.Tileset.GetTileRect(i)
			r.tileCache[i+tile.Tileset.FirstGID] = imaging.Crop(img, rect)
			if tile.ID == i {
				timg = r.tileCache[i+tile.Tileset.FirstGID]
			}
		}
	}

	return r.engine.RotateTileImage(tile, timg), nil
}

func (r *Renderer) _renderLayer(layer *tiled.Layer) error {
	var xs, xe, xi, ys, ye, yi int
	if r.m.RenderOrder == "" || r.m.RenderOrder == "right-down" {
		xs = 0
		xe = r.m.Width
		xi = 1
		ys = 0
		ye = r.m.Height
		yi = 1
	} else {
		return ErrUnsupportedRenderOrder
	}

	i := 0
	for y := ys; y*yi < ye; y = y + yi {
		for x := xs; x*xi < xe; x = x + xi {
			if layer.Tiles[i].IsNil() {
				i++
				continue
			}

			img, err := r.getTileImage(layer.Tiles[i])
			if err != nil {
				return err
			}

			pos := r.engine.GetTilePosition(x, y)

			if layer.Opacity < 1 {
				mask := image.NewUniform(color.Alpha{uint8(layer.Opacity * 255)})

				draw.DrawMask(r.Result, pos, img, img.Bounds().Min, mask, mask.Bounds().Min, draw.Over)
			} else {
				draw.Draw(r.Result, pos, img, img.Bounds().Min, draw.Over)
			}

			i++
		}
	}

	return nil
}

// RenderGroupLayer renders single map layer in a certain group.
func (r *Renderer) RenderGroupLayer(groupID, layerID int) error {
	if groupID >= len(r.m.Groups) {
		return ErrOutOfBounds
	}
	group := r.m.Groups[groupID]

	if layerID >= len(group.Layers) {
		return ErrOutOfBounds
	}
	return r._renderLayer(group.Layers[layerID])
}

// RenderLayer renders single map layer.
func (r *Renderer) RenderLayer(id int) error {
	if id >= len(r.m.Layers) {
		return ErrOutOfBounds
	}
	return r._renderLayer(r.m.Layers[id])
}

// RenderVisibleLayers renders all visible map layers.
func (r *Renderer) RenderVisibleLayers() error {
	for i := range r.m.Layers {
		if !r.m.Layers[i].Visible {
			continue
		}

		if err := r.RenderLayer(i); err != nil {
			return err
		}
	}

	return nil
}

// Clear clears the render result to allow for separation of layers. For example, you can
// render a layer, make a copy of the render, clear the renderer, and repeat for each
// layer in the Map.
func (r *Renderer) Clear() {
	r.Result = image.NewNRGBA(r.engine.GetFinalImageSize())
}



// OrthogonalRendererEngine represents orthogonal rendering engine.
type OrthogonalRendererEngine struct {
	m *tiled.Map
}

// Init initializes rendering engine with provided map options.
func (e *OrthogonalRendererEngine) Init(m *tiled.Map) {
	e.m = m
}

// GetFinalImageSize returns final image size based on map data.
func (e *OrthogonalRendererEngine) GetFinalImageSize() image.Rectangle {
	return image.Rect(0, 0, e.m.Width*e.m.TileWidth, e.m.Height*e.m.TileHeight)
}

// RotateTileImage rotates provided tile layer.
func (e *OrthogonalRendererEngine) RotateTileImage(tile *tiled.LayerTile, img image.Image) image.Image {
	timg := img
	if tile.DiagonalFlip {
		timg = imaging.FlipH(imaging.Rotate270(timg))
	}
	if tile.HorizontalFlip {
		timg = imaging.FlipH(timg)
	}
	if tile.VerticalFlip {
		timg = imaging.FlipV(timg)
	}

	return timg
}

// GetTilePosition returns tile position in image.
func (e *OrthogonalRendererEngine) GetTilePosition(x, y int) image.Rectangle {
	return image.Rect(x*e.m.TileWidth,
		y*e.m.TileHeight,
		(x+1)*e.m.TileWidth,
		(y+1)*e.m.TileHeight)
}