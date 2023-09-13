package ebiten_extended

import (
	"math"
)

type Circle struct {
	transform *Transform
	radius    float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) SetTransform(transform *Transform) {
	c.transform = transform
}

func (c *Circle) GetPositionX() float64 {
	return c.transform.geoM.Element(0, 2)
}

func (c *Circle) GetPositionY() float64 {
	return c.transform.geoM.Element(1, 2)
}

func (c *Circle) Intersect(shape Shape) bool {
	switch s := shape.(type) {
	case *Circle:
		position1_x := c.transform.position.X
		position1_y := c.transform.position.Y
		position2_x := s.transform.position.X
		position2_y := s.transform.position.Y
		x := position1_x - position2_x
		y := position1_y - position2_y
		return math.Sqrt(x*x+y*y) <= c.radius+s.radius
	}
	return false
}
