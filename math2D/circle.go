package math2D

// Circle represents a 2D circle by its center and radius.
type Circle struct {
	center Vector2D
	radius float64
}

// NewCircle returns a circle with the given center and radius.
func NewCircle(center Vector2D, radius float64) Circle {
	return Circle{center: center, radius: radius}
}

// SetCenter sets the center of the circle.
func (c *Circle) SetCenter(center Vector2D) {
	c.center = center
}

// GetCenterPosition returns the center of the circle.
func (c *Circle) GetCenterPosition() Vector2D {
	return c.center
}

// GetRadius returns the radius of the circle.
func (c *Circle) GetRadius() float64 {
	return c.radius
}

// SetRadius sets the radius of the circle.
func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}
