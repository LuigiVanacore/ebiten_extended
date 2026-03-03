package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Node2D
	textureRect math2D.Rectangle
	texture     *ebiten.Image
	layerIndex int
}

func NewSprite(name string, texture *ebiten.Image, layerIndex int, isPivotToCenter bool) *Sprite {

	textureRect := math2D.NewRectangle(math2D.NewVector2D(float64(texture.Bounds().Min.X), float64(texture.Bounds().Min.Y)),
		math2D.NewVector2D(float64(texture.Bounds().Max.X), float64(texture.Bounds().Max.Y)))

	sprite := &Sprite{Node2D: *NewNode2D(name), textureRect: textureRect, texture: texture, layerIndex: layerIndex}

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

func (s *Sprite) GetTextureRect() math2D.Rectangle {
	return s.textureRect
}

func (s *Sprite) SetTextureRect(width, height float64) {
	s.textureRect = math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(width, height))
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}
func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
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
	if s.texture != nil {
		s.setPositionToPivot(op)
		target.DrawImage(s.texture, op)
	}
}
