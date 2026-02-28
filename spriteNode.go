package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteNode struct {
	Node2D
	textureRect math2D.Rectangle
	texture     *ebiten.Image
}

func NewSprite(name string, texture *ebiten.Image, isPivotToCenter bool) *SpriteNode {

	textureRect := math2D.NewRectangle(math2D.NewVector2D(float64(texture.Bounds().Min.X), float64(texture.Bounds().Min.Y)),
		math2D.NewVector2D(float64(texture.Bounds().Max.X), float64(texture.Bounds().Max.Y)))

	sprite := &SpriteNode{Node2D: *NewNode2D(name), textureRect: textureRect, texture: texture}

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

func (s *SpriteNode) GetTextureRect() math2D.Rectangle {
	return s.textureRect
}

func (s *SpriteNode) SetTextureRect(width, height float64) {
	s.textureRect = math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(width, height))
}

func (s *SpriteNode) GetTexture() *ebiten.Image {
	return s.texture
}
func (s *SpriteNode) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

func (s *SpriteNode) SetPivotToCenter() {
	t := s.GetTransform()
	t.SetPivot(s.GetTextureRect().GetCenter().X(), s.GetTextureRect().GetCenter().Y())
	s.SetTransform(t)
}

func (s *SpriteNode) updateGeoM(op *ebiten.DrawImageOptions) {
	t := s.GetTransform()
	op.GeoM.Translate(-t.GetPivot().X(), -t.GetPivot().Y())
}

func (s *SpriteNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DebugInfo()
	if s.texture != nil {
		s.updateGeoM(op)
		target.DrawImage(s.texture, op)
	}
}

func (s *SpriteNode) DebugInfo() {
	// DebugInfo logic has been decoupled from the global engine state.
	// If you need debug, do it from the engine or inject a debug flag.
}
