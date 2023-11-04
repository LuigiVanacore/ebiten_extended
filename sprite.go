package ebiten_extended

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	transform transform.Transform
	textureRect *math2D.Rectangle
	texture     *ebiten.Image
}

func NewSprite(texture *ebiten.Image, isPivotToCenter bool) *Sprite {
	
	textureRect := math2D.NewRectangle(math2D.NewVector2D(float64(texture.Bounds().Min.X),float64(texture.Bounds().Min.Y)),
	math2D.NewVector2D(float64(texture.Bounds().Max.X),float64(texture.Bounds().Max.Y)))

	sprite := &Sprite{textureRect: &textureRect, texture: texture}

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

func (s *Sprite) GetTextureRect() *math2D.Rectangle {
	return s.textureRect
}

func (s *Sprite) SetTextureRect(width, height float64) {
	textureRect := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(width, height))
	s.textureRect = &textureRect
}

func (s *Sprite) GetTexture() *ebiten.Image {
	return s.texture
}
func (s *Sprite) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

func (s *Sprite) SetPivotToCenter() {
	s.transform.SetPivot(s.GetTextureRect().GetCenter().X(), s.GetTextureRect().GetCenter().Y())
}

func (s *Sprite) GetTransform() *transform.Transform {
	return &s.transform
}

func (s *Sprite) SetTransform(transform transform.Transform) {
	s.transform = transform
}

func (s *Sprite) SetPosition(x, y float64) {
	s.transform.SetPosition(x, y)
}

func (s *Sprite) GetPosition() math2D.Vector2D {
	return s.transform.GetPosition()
}

func (s *Sprite) updateGeoM(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(-s.transform.GetPivot().X(), -s.transform.GetPivot().Y())
	geom.Rotate(float64(s.transform.GetRotation()%360) * 2 * math.Pi / 360)
	geom.Translate(s.transform.GetPivot().X(), s.transform.GetPivot().Y())
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
