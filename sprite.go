package ebiten_extended

import (
	"image"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Node2D
	textureRect math2D.Rectangle
	texture     *ebiten.Image
	subImage    *ebiten.Image // cached sub-region; updated by updateSubImage
	layerIndex  int
}

func NewSprite(name string, texture *ebiten.Image, layerIndex int, isPivotToCenter bool) *Sprite {
	if texture == nil {
		sprite := &Sprite{Node2D: *NewNode2D(name), layerIndex: layerIndex}
		return sprite
	}

	// Use Dx/Dy so the rect is correct even for sub-images with non-zero Min.
	textureRect := math2D.NewRectangle(
		math2D.ZeroVector2D(),
		math2D.NewVector2D(float64(texture.Bounds().Dx()), float64(texture.Bounds().Dy())),
	)

	sprite := &Sprite{Node2D: *NewNode2D(name), textureRect: textureRect, texture: texture, layerIndex: layerIndex}
	sprite.updateSubImage()

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

func (s *Sprite) GetTextureRect() math2D.Rectangle {
	return s.textureRect
}

// SetTextureRect sets the source region (offset from texture origin) used for rendering.
// Use this to select a frame from a sprite sheet; (0,0) is the top-left of the texture.
func (s *Sprite) SetTextureRect(x, y, width, height float64) {
	s.textureRect = math2D.NewRectangle(math2D.NewVector2D(x, y), math2D.NewVector2D(width, height))
	s.updateSubImage()
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}

func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
	if texture != nil {
		s.textureRect = math2D.NewRectangle(
			math2D.ZeroVector2D(),
			math2D.NewVector2D(float64(texture.Bounds().Dx()), float64(texture.Bounds().Dy())),
		)
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

func (s *Sprite) setPositionToPivot(op *ebiten.DrawImageOptions) {
	tr := s.GetTransform()
	pivot := tr.GetPivot()
	op.GeoM.Translate(-pivot.X(), -pivot.Y())
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if s.subImage == nil {
		return
	}
	s.setPositionToPivot(op)
	target.DrawImage(s.subImage, op)
}
