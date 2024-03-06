package math2D

type Circle struct {
	center Vector2D
	radius float64
}

func NewCircle(center Vector2D, radius float64) Circle {
	return Circle{center: center, radius: radius}
}

func (c *Circle) SetCenter(center Vector2D) {
	c.center = center
}

func (c *Circle) GetCenterPosition() Vector2D {
	return c.center
}

func (c *Circle) GetRadius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}
