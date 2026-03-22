package ludum

import (
	"image"
	"math"

	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Node2D
	textureRect math2d.Rectangle
	texture     *ebiten.Image
	subImage    *ebiten.Image // cached sub-region; updated by updateSubImage
	layerIndex  int

	colorScale ebiten.ColorScale
	flipX      bool
	flipY      bool
	visible    bool // when false, Draw is skipped
	blend      ebiten.Blend
	filter     ebiten.Filter
}

func NewSprite(name string, texture *ebiten.Image, layerIndex int, isPivotToCenter bool) *Sprite {
	if texture == nil {
		sprite := &Sprite{Node2D: *NewNode2D(name), layerIndex: layerIndex}
		return sprite
	}

	// Use Dx/Dy so the rect is correct even for sub-images with non-zero Min.
	textureRect := math2d.NewRectangle(
		math2d.ZeroVector2D(),
		math2d.NewVector2D(float64(texture.Bounds().Dx()), float64(texture.Bounds().Dy())),
	)

	sprite := &Sprite{
		Node2D:      *NewNode2D(name),
		textureRect: textureRect,
		texture:     texture,
		layerIndex:  layerIndex,
		visible:     true,
	}
	sprite.updateSubImage()

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

func (s *Sprite) GetTextureRect() math2d.Rectangle {
	return s.textureRect
}

// SetTextureRect sets the source region (offset from texture origin) used for rendering.
// Use this to select a frame from a sprite sheet; (0,0) is the top-left of the texture.
// Values are clamped to texture bounds when texture is set.
func (s *Sprite) SetTextureRect(x, y, width, height float64) {
	if s.texture != nil {
		b := s.texture.Bounds()
		texW, texH := float64(b.Dx()), float64(b.Dy())
		x = math.Max(0, math.Min(x, texW))
		y = math.Max(0, math.Min(y, texH))
		maxW := texW - x
		maxH := texH - y
		width = math.Max(0, math.Min(width, maxW))
		height = math.Max(0, math.Min(height, maxH))
	}
	s.textureRect = math2d.NewRectangle(math2d.NewVector2D(x, y), math2d.NewVector2D(width, height))
	s.updateSubImage()
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}

func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
	if texture != nil {
		s.textureRect = math2d.NewRectangle(
			math2d.ZeroVector2D(),
			math2d.NewVector2D(float64(texture.Bounds().Dx()), float64(texture.Bounds().Dy())),
		)
	} else {
		s.textureRect = math2d.NewRectangle(math2d.ZeroVector2D(), math2d.ZeroVector2D())
	}
	s.updateSubImage()
}

// updateSubImage caches a sub-image view of the texture according to the current textureRect.
// Called whenever texture or textureRect changes; zero-allocation at Draw time.
func (s *Sprite) updateSubImage() {
	if s.texture == nil {
		s.subImage = nil
		return
	}
	sz := s.textureRect.GetSize()
	if sz.X() <= 0 || sz.Y() <= 0 {
		s.subImage = nil
		return
	}
	b := s.texture.Bounds()
	ox := b.Min.X + int(s.textureRect.GetPosition().X())
	oy := b.Min.Y + int(s.textureRect.GetPosition().Y())
	rect := image.Rect(
		ox,
		oy,
		ox+int(s.textureRect.GetSize().X()),
		oy+int(s.textureRect.GetSize().Y()),
	)
	s.subImage = s.texture.SubImage(rect).(*ebiten.Image)
}

func (s *Sprite) GetLayer() int {
	return s.layerIndex
}

func (s *Sprite) SetLayer(layerIndex int) {
	s.layerIndex = layerIndex
}

func (s *Sprite) SetPivotToCenter() {
	tr := s.GetTransform()
	center := s.textureRect.GetCenter()
	tr.SetPivot(center.X(), center.Y())
	s.SetTransform(tr)
}

// SetPivot sets the pivot point (center of scale/rotation) in local texture coordinates.
func (s *Sprite) SetPivot(x, y float64) {
	tr := s.GetTransform()
	tr.SetPivot(x, y)
	s.SetTransform(tr)
}

// SetColorScale sets the color modulation for rendering. Use for tinting, flash on hit, etc.
func (s *Sprite) SetColorScale(cs ebiten.ColorScale) {
	s.colorScale = cs
}

// GetColorScale returns the current color scale.
func (s *Sprite) GetColorScale() ebiten.ColorScale {
	return s.colorScale
}

// SetAlpha sets the alpha to the given value (0–1), using white. Resets any previous SetColorScale.
func (s *Sprite) SetAlpha(alpha float32) {
	s.colorScale = ebiten.ColorScale{}
	s.colorScale.ScaleAlpha(alpha)
}

// SetFlip sets both horizontal and vertical flip.
func (s *Sprite) SetFlip(flipX, flipY bool) {
	s.flipX = flipX
	s.flipY = flipY
}

// SetFlipX sets horizontal flip of the sprite.
func (s *Sprite) SetFlipX(flip bool) {
	s.flipX = flip
}

// GetFlipX returns whether the sprite is flipped horizontally.
func (s *Sprite) GetFlipX() bool {
	return s.flipX
}

// SetFlipY sets vertical flip of the sprite.
func (s *Sprite) SetFlipY(flip bool) {
	s.flipY = flip
}

// GetFlipY returns whether the sprite is flipped vertically.
func (s *Sprite) GetFlipY() bool {
	return s.flipY
}

// SetTint sets RGB tint (0–1) with full alpha. Convenience over SetColorScale for simple tinting.
func (s *Sprite) SetTint(r, g, b float32) {
	s.SetTintRGBA(r, g, b, 1)
}

// SetTintRGBA sets RGBA tint (0–1). Use for tinting with transparency.
func (s *Sprite) SetTintRGBA(r, g, b, a float32) {
	s.colorScale = ebiten.ColorScale{}
	s.colorScale.SetR(r)
	s.colorScale.SetG(g)
	s.colorScale.SetB(b)
	s.colorScale.SetA(a)
}

// SetVisible sets whether the sprite is drawn. When false, Draw is skipped.
func (s *Sprite) SetVisible(visible bool) {
	s.visible = visible
}

// GetVisible returns whether the sprite is visible.
func (s *Sprite) GetVisible() bool {
	return s.visible
}

// GetSize returns the size of the current texture rect (width, height).
func (s *Sprite) GetSize() math2d.Vector2D {
	return s.textureRect.GetSize()
}

// SetBlendMode sets the blend mode for rendering (e.g. additive, multiply).
func (s *Sprite) SetBlendMode(blend ebiten.Blend) {
	s.blend = blend
}

// GetBlendMode returns the current blend mode.
func (s *Sprite) GetBlendMode() ebiten.Blend {
	return s.blend
}

// SetFilter sets the filter mode for scaling (e.g. FilterLinear, FilterNearest).
func (s *Sprite) SetFilter(filter ebiten.Filter) {
	s.filter = filter
}

// GetFilter returns the current filter mode.
func (s *Sprite) GetFilter() ebiten.Filter {
	return s.filter
}

// Clone returns a new sprite sharing the same texture and texture rect, with copied state.
// Transform, color, flip, visibility, blend, and filter are copied. The clone has no parent.
func (s *Sprite) Clone() *Sprite {
	clone := &Sprite{
		Node2D:      *NewNode2D(s.GetName() + "_clone"),
		textureRect: s.textureRect,
		texture:     s.texture,
		layerIndex:  s.layerIndex,
		colorScale:  s.colorScale,
		flipX:       s.flipX,
		flipY:       s.flipY,
		visible:     s.visible,
		blend:       s.blend,
		filter:      s.filter,
	}
	clone.SetTransform(s.GetTransform())
	clone.updateSubImage()
	return clone
}

// GetWorldBounds returns the axis-aligned bounding box in world space for culling.
func (s *Sprite) GetWorldBounds() math2d.Rectangle {
	pos := s.GetWorldPosition()
	scale := s.GetWorldScale()
	sz := s.textureRect.GetSize()
	pivot := s.GetPivot()
	// Pivot at (px,py); position is where pivot ends up. Local corners: (0,0), (w,0), (0,h), (w,h).
	// In world (ignoring rotation): pos + scale * (corner - pivot)
	x0 := pos.X() + scale.X()*(0-pivot.X())
	y0 := pos.Y() + scale.Y()*(0-pivot.Y())
	x1 := pos.X() + scale.X()*(sz.X()-pivot.X())
	y1 := pos.Y() + scale.Y()*(sz.Y()-pivot.Y())
	minX, maxX := x0, x1
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minY, maxY := y0, y1
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	return math2d.NewRectangle(
		math2d.NewVector2D(minX, minY),
		math2d.NewVector2D(maxX-minX, maxY-minY),
	)
}

// applyFlip appends a flip transform in image space around the pivot (correct for scaled/rotated sprites).
func (s *Sprite) applyFlip(op *ebiten.DrawImageOptions) {
	if !s.flipX && !s.flipY {
		return
	}
	pivot := s.GetPivot()
	flipGeom := ebiten.GeoM{}
	flipGeom.Translate(-pivot.X(), -pivot.Y())
	sx, sy := 1.0, 1.0
	if s.flipX {
		sx = -1
	}
	if s.flipY {
		sy = -1
	}
	flipGeom.Scale(sx, sy)
	flipGeom.Translate(pivot.X(), pivot.Y())
	op.GeoM.Concat(flipGeom)
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if s.subImage == nil || !s.visible {
		return
	}
	s.applyFlip(op)
	op.ColorScale = s.colorScale
	op.Blend = s.blend
	op.Filter = s.filter
	target.DrawImage(s.subImage, op)
}
