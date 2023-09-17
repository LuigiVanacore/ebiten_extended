package ebiten_extended

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Transform struct {
	position Vector2D
	pivot    Vector2D
	rotation int
	geoM     ebiten.GeoM
}

func (t *Transform) GetPosition() Vector2D {
	return t.position
}

func (t *Transform) SetPosition(x, y float64) {
	t.position.X = x
	t.position.Y = y
}

func (t *Transform) GetRotation() int {
	return t.rotation
}

func (t *Transform) SetRotation(rotation int) {
	t.rotation = rotation
}

func (t *Transform) SetPivot(x, y float64) {
	t.pivot.X = x
	t.pivot.Y = y
}

func (t *Transform) GetPivot() Vector2D {
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

func (t *Transform) Move(x, y float64) {
	t.position.X += x
	t.position.Y += y
}

func (t *Transform) UpdateGeoM(geom ebiten.GeoM) ebiten.GeoM {
	geom.Translate(-t.pivot.X, -t.pivot.Y)
	geom.Rotate(float64(t.rotation%360) * 2 * math.Pi / 360)
	geom.Translate(t.position.X, t.position.Y)
	return geom
}