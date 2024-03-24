package transform

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	position math2D.Vector2D
	pivot    math2D.Vector2D
	rotation int
	geoM     ebiten.GeoM
}

func NewTransform(position math2D.Vector2D, pivot math2D.Vector2D, rotation int) Transform{
	return Transform{ position: position, pivot: pivot, rotation: rotation}
}


func (t *Transform) GetPosition() math2D.Vector2D {
	return t.position
}

func (t *Transform) SetPosition(x, y float64) {
	t.position.SetPosition(x, y)
}

func (t *Transform) GetRotation() int {
	return t.rotation
}

func (t *Transform) SetRotation(rotation int) {
	t.rotation = rotation
}

func (t *Transform) SetPivot(x, y float64) {
	t.pivot.SetPosition(x , y)
}

func (t *Transform) Scale(x, y float64) {
	t.geoM.Scale(x, y)
}

func (t *Transform) GetPivot() math2D.Vector2D {
	return t.pivot
}

func (t *Transform) GetGeoM() ebiten.GeoM {
	return t.geoM
}

func (t *Transform) SetGeoM(geoM ebiten.GeoM) {
	t.geoM = geoM
}

func (t *Transform) Rotate(rotation int) {
	t.rotation += rotation
}

func (t *Transform) Translate(x, y float64) {
	t.position.Translate(x, y)
}

func (t *Transform) UpdateGeoM(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(-t.pivot.X(), -t.pivot.Y())
	geom.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)
	geom.Translate(t.position.X(), t.position.Y())
	return geom
}

func (t *Transform) Concat(transform Transform) {
	position := transform.GetPosition()
	t.Translate(position.X(), position.Y())
	rotation := transform.GetRotation()
	t.Rotate(rotation)
	geoM := transform.GetGeoM()
	t.geoM.Concat(geoM)
}