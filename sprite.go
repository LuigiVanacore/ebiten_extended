package ebiten_extended

import (
	"fmt"

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

func (s *Sprite) SetScale(x, y float64) {
	s.transform.Scale(x, y)
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

func (s *Sprite) SetRotation(rotation int) {
	s.transform.SetRotation(rotation)
}

func (s *Sprite) GetRotation() int {
	return s.transform.GetRotation()
}

func (s *Sprite) updateGeoM(op *ebiten.DrawImageOptions ) {
	op.GeoM.Translate(-s.transform.GetPivot().X(), -s.transform.GetPivot().Y())
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DebugInfo()
	if s.texture != nil {
		s.updateGeoM(op)
		target.DrawImage(s.texture, op)
	}
}


func (s *Sprite) DebugInfo() {
	if GameManager().IsDebug() {
		fmt.Printf("The position is x: %f y: %f, the rotation is %d \n", s.GetPosition().X(), s.GetPosition().Y(), s.GetRotation())
	}
}