package ebiten_extended

type Shape interface {
	Intersect(shape Shape) bool
	SetTransform(transform *Transform)
}
