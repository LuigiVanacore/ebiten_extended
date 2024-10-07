package ebiten_extended

import (
	"fmt"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Node2D
	textureRect math2D.Rectangle
	texture     *ebiten.Image
}

func NewSprite(name string, texture *ebiten.Image, isPivotToCenter bool) *Sprite {

	textureRect := math2D.NewRectangle(math2D.NewVector2D(float64(texture.Bounds().Min.X), float64(texture.Bounds().Min.Y)),
		math2D.NewVector2D(float64(texture.Bounds().Max.X), float64(texture.Bounds().Max.Y)))

	sprite := &Sprite{Node2D: *NewNode2D(name), textureRect: textureRect, texture: texture}

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

func (s *Sprite) SetPivotToCenter() {
	s.transform.SetPivot(s.GetTextureRect().GetCenter().X(), s.GetTextureRect().GetCenter().Y())
}

func (s *Sprite) updateGeoM(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(-s.transform.GetPivot().X(), -s.transform.GetPivot().Y())
}

func (s *Sprite) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DebugInfo()
	if s.texture != nil {
		target.DrawImage(s.texture, op)
	}
}

func (s *Sprite) DebugInfo() {
	if GameManager().IsDebug() {
		fmt.Printf("The position is x: %f y: %f, the rotation is %d \n", s.GetPosition().X(), s.GetPosition().Y(), s.GetRotation())
	}
}
