package math2d

type Shape interface {
	Intersect(shape Shape) bool
}
