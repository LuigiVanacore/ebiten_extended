package math2D

type Shape interface {
	Intersect(shape Shape) bool
}
