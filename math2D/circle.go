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

func (c *Circle) Intersect(shape Shape) bool {
	// switch s := shape.(type) {
	// case *Circle:
	// 	position1_x := c.center.X()
	// 	position1_y := c.center.Y()
	// 	position2_x := s.center.X()
	// 	position2_y := s.center.Y()
	// 	x := position1_x - position2_x
	// 	y := position1_y - position2_y
	// 	return math.Sqrt(x*x+y*y) <= c.radius+s.radius
	// }
	return false
}
