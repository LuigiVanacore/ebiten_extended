package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

// SpriteNode represents a visual 2D node capable of drawing a texture or portions of it onto the screen.
type SpriteNode struct {
	Node2D
	textureRect math2D.Rectangle
	texture     *ebiten.Image
}

// NewSprite creates a new SpriteNode with a given name and image, optionally setting its rotation pivot to the image center.
func NewSprite(name string, texture *ebiten.Image, isPivotToCenter bool) *SpriteNode {

	textureRect := math2D.NewRectangle(math2D.NewVector2D(float64(texture.Bounds().Min.X), float64(texture.Bounds().Min.Y)),
		math2D.NewVector2D(float64(texture.Bounds().Max.X), float64(texture.Bounds().Max.Y)))

	sprite := &SpriteNode{Node2D: *NewNode2D(name), textureRect: textureRect, texture: texture}

	if isPivotToCenter {
		sprite.SetPivotToCenter()
	}

	return sprite
}

// GetTextureRect returns the source rectangle defining which section of the texture is drawn.
func (s *SpriteNode) GetTextureRect() math2D.Rectangle {
	return s.textureRect
}

// SetTextureRect updates the visible source area mapping from the top-left of the texture.
func (s *SpriteNode) SetTextureRect(width, height float64) {
	s.textureRect = math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(width, height))
}

// GetTexture returns the underlying ebiten.Image being used by this sprite.
func (s *SpriteNode) GetTexture() *ebiten.Image {
	return s.texture
}

// SetTexture assigns a new ebiten.Image to be rendered by this sprite.
func (s *SpriteNode) SetTexture(texture *ebiten.Image) {
	s.texture = texture
}

// SetPivotToCenter adjusts the sprite's rotation and scaling origin point to the geometric center of its texture rectangle.
func (s *SpriteNode) SetPivotToCenter() {
	t := s.GetTransform()
	t.SetPivot(s.GetTextureRect().GetCenter().X(), s.GetTextureRect().GetCenter().Y())
	s.SetTransform(t)
}

func (s *SpriteNode) updateGeoM(op *ebiten.DrawImageOptions) {
	t := s.GetTransform()
	op.GeoM.Translate(-t.GetPivot().X(), -t.GetPivot().Y())
}

// Draw renders the sprite's image onto the provided target applying the current transformation matrix.
func (s *SpriteNode) Draw(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	s.DebugInfo()
	if s.texture != nil {
		s.updateGeoM(op)
		target.DrawImage(s.texture, op)
	}
}

// DebugInfo outputs debug visualization or state data for this specific sprite.
func (s *SpriteNode) DebugInfo() {
	// DebugInfo logic has been decoupled from the global engine state.
	// If you need debug, do it from the engine or inject a debug flag.
}
