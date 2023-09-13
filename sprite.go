package ebiten_extended

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	transform Transform
	textureRect Rect
	texture     *ebiten.Image
}

func NewSprite(texture *ebiten.Image, isPivotToCenter bool) *Sprite {

	textureRect := NewRect(float64(texture.Bounds().Min.X),
		float64(texture.Bounds().Min.Y),
		float64(texture.Bounds().Max.X),
		float64(texture.Bounds().Max.Y))

	sprite := &Sprite{textureRect: *textureRect, texture: texture}

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

func (s *Sprite) GetTextureRect() Rect {
	return s.textureRect
}

func (s *Sprite) SetTextureRect(width, height float64) {
	s.textureRect = Rect{Width: width, Height: height}
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}
func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

func (s *Sprite) SetPivotToCenter() {
	rect := s.GetTextureRect()
	x, y := rect.GetCenter()
	s.transform.SetPivot(x, y)
}

func (s *Sprite) GetTransform() *Transform {
	return &s.transform
}

func (s *Sprite) SetTransform(transform Transform) {
	s.transform = transform
}

func (s *Sprite) SetPosition(x, y float64) {
	s.transform.SetPosition(x, y)
}

func (s *Sprite) GetPosition() Vector2D {
	return s.transform.GetPosition()
}

func (s *Sprite) updateGeoM(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(-s.transform.pivot.X, -s.transform.pivot.Y)
	geom.Rotate(float64(s.transform.rotation%360) * 2 * math.Pi / 360)
	geom.Translate(s.transform.position.X, s.transform.position.Y)
	return geom
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	local_op := &ebiten.DrawImageOptions{}
	local_op.GeoM = op.GeoM
	local_op.GeoM = s.updateGeoM(op.GeoM)
	s.transform.SetGeoM(local_op.GeoM)
	if s.texture != nil {
		target.DrawImage(s.texture, local_op)
	}
}
