package transform

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

type Transform struct {
	position math2D.Vector2D
	pivot    math2D.Vector2D
	rotation int
	scale math2D.Vector2D
}

func NewTransform(position math2D.Vector2D, pivot math2D.Vector2D, rotation int) Transform{
	return Transform{ position: position, pivot: pivot, rotation: rotation}
}


func (t *Transform) GetPosition() math2D.Vector2D {
	return t.position
}

func (t *Transform) SetPosition(position math2D.Vector2D) {
	t.position.SetPosition(position)
}

func (t *Transform) GetRotation() int {
	return t.rotation
}

func (t *Transform) SetRotation(rotation int) {
	t.rotation = rotation
}

func (t *Transform) SetPivot(pivot math2D.Vector2D) {
	t.pivot.SetPosition(pivot)
}

func (t *Transform) SetScale(x, y float64) {
	t.scale = math2D.NewVector2D(x, y)
}

func (t *Transform) GetScale() math2D.Vector2D {
	return t.scale
}

func (t *Transform) GetPivot() math2D.Vector2D {
	return t.pivot
}

func (t *Transform) Rotate(rotation int) {
	t.rotation += rotation
}

func (t *Transform) Translate(vec math2D.Vector2D) {
	t.position.Translate(vec.X(), vec.Y())
}

// func (t *Transform) UpdateGeoM(geom ebiten.GeoM) ebiten.GeoM {
// 	geom.Translate(-t.pivot.X(), -t.pivot.Y())
// 	geom.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)
// 	geom.Translate(t.position.X(), t.position.Y())
// 	return geom
// }

func (t *Transform) Concat(transform Transform) {
	position := transform.GetPosition()
	t.Translate(position)
	rotation := transform.GetRotation()
	t.Rotate(rotation + t.rotation)
}