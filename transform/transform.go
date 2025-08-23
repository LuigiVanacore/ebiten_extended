package transform

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

// Transform represents a 2D position, pivot, rotation and scale.
type Transform struct {
	position math2D.Vector2D
	pivot    math2D.Vector2D
	rotation int
	scale    math2D.Vector2D
}

// NewTransform creates a new transform with the given position, pivot and rotation.
// The scale is initialized to (1,1).
func NewTransform(position math2D.Vector2D, pivot math2D.Vector2D, rotation int) Transform {
	return Transform{position: position, pivot: pivot, rotation: rotation, scale: math2D.NewVector2D(1, 1)}
}

// GetPosition returns the transform's position.
func (t *Transform) GetPosition() math2D.Vector2D {
	return t.position
}

// SetPosition sets the transform's position.
func (t *Transform) SetPosition(position math2D.Vector2D) {
	t.position.SetPosition(position)
}

// GetRotation returns the current rotation in degrees.
func (t *Transform) GetRotation() int {
	return t.rotation
}

// SetRotation sets the rotation in degrees.
func (t *Transform) SetRotation(rotation int) {
	t.rotation = rotation
}

// SetPivot sets the pivot point for transformations.
func (t *Transform) SetPivot(pivot math2D.Vector2D) {
	t.pivot.SetPosition(pivot)
}

// SetScale sets the scaling factors.
func (t *Transform) SetScale(x, y float64) {
	t.scale = math2D.NewVector2D(x, y)
}

// GetScale returns the current scaling factors.
func (t *Transform) GetScale() math2D.Vector2D {
	return t.scale
}

// GetPivot returns the current pivot point.
func (t *Transform) GetPivot() math2D.Vector2D {
	return t.pivot
}

// Rotate adds rotation in degrees to the transform.
func (t *Transform) Rotate(rotation int) {
	t.rotation += rotation
}

// Translate moves the transform by the provided vector.
func (t *Transform) Translate(vec math2D.Vector2D) {
	t.position.Translate(vec.X(), vec.Y())
}



// Concat applies another transform's translation and rotation to the current one.
func (t *Transform) Concat(transform Transform) {
<<<<<<< HEAD
	position := transform.GetPosition()
	t.Translate(position)
	rotation := transform.GetRotation()
	t.Rotate(rotation)
}
=======
	t.Translate(transform.GetPosition())
	t.Rotate(transform.GetRotation())
}
>>>>>>> 756ed3b79758c5f65aad19fb6a9076530088a44f
