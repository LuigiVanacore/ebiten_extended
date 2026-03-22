package transform

import (
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/hajimehoshi/ebiten/v2"
)

// Transform holds 2D position, pivot, rotation (radians), scale, and an optional GeoM for scale/skew.
type Transform struct {
	position math2d.Vector2D
	pivot    math2d.Vector2D
	rotation float64
	scale    math2d.Vector2D
	geoM     ebiten.GeoM
}

// NewTransform returns a transform with the given position, pivot, and rotation (radians).
// The initial scale is initialized to (1,1).
func NewTransform(position math2d.Vector2D, pivot math2d.Vector2D, rotation float64) Transform {
	return Transform{position: position, pivot: pivot, rotation: rotation, scale: math2d.NewVector2D(1, 1)}
}

func (t *Transform) GetPosition() math2d.Vector2D {
	return t.position
}

// SetPosition sets the transform's position.
func (t *Transform) SetPosition(position math2d.Vector2D) {
	t.position.SetPosition(position)
}

func (t *Transform) GetRotation() float64 {
	return t.rotation
}

func (t *Transform) SetRotation(rotation float64) {
	t.rotation = rotation
}

func (t *Transform) SetPivot(x, y float64) {
	t.pivot.SetPosition(math2d.NewVector2D(x, y))
}

// SetScale sets the scaling factors.
func (t *Transform) SetScale(x, y float64) {
	t.scale = math2d.NewVector2D(x, y)
}

// Scale is an alias for SetScale and exists to match Node2D's usage.
func (t *Transform) Scale(x, y float64) {
	t.SetScale(x, y)
}

// GetScale returns the current scaling factors.
func (t *Transform) GetScale() math2d.Vector2D {
	return t.scale
}

// GetPivot returns the current pivot point.
func (t *Transform) GetPivot() math2d.Vector2D {
	return t.pivot
}

func (t *Transform) GetGeoM() ebiten.GeoM {
	return t.geoM
}

func (t *Transform) SetGeoM(geoM ebiten.GeoM) {
	t.geoM = geoM
}

func (t *Transform) Rotate(rotation float64) {
	t.rotation += rotation
}

func (t *Transform) Translate(x, y float64) {
	t.position.Translate(x, y)
}

func (t *Transform) UpdateGeoM(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(-t.pivot.X(), -t.pivot.Y())
	geom.Scale(t.scale.X(), t.scale.Y())
	geom.Rotate(t.rotation)
	geom.Translate(t.position.X(), t.position.Y())
	return geom
}

// Concat applies other (child/local) on top of t (parent/world).
// The child position is first scaled by the parent scale, then rotated by the parent rotation,
// and finally added to the parent position — matching the standard 2D transform hierarchy.
// geoM from other is not applied.
func (t *Transform) Concat(other Transform) {
	// Scale local position by parent scale, then rotate by parent rotation
	scaledX := other.GetPosition().X() * t.scale.X()
	scaledY := other.GetPosition().Y() * t.scale.Y()
	rotated := math2d.NewVector2D(scaledX, scaledY).RotateVector(t.rotation)
	t.Translate(rotated.X(), rotated.Y())
	t.scale.SetPosition(math2d.NewVector2D(t.scale.X()*other.GetScale().X(), t.scale.Y()*other.GetScale().Y()))
	t.Rotate(other.GetRotation())
}
