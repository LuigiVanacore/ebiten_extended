package math2d

// Rectangle represents an axis-aligned rectangle by its top-left position and size (width, height).
type Rectangle struct {
	position Vector2D
	size     Vector2D
}

// NewRectangle returns a rectangle with the given position and size.
func NewRectangle(position Vector2D, size Vector2D) Rectangle {
	return Rectangle{
		position: position,
		size:     size,
	}
}

func (r *Rectangle) GetPosition() Vector2D {
	return r.position
}

func (r *Rectangle) SetPosition(position Vector2D) {
	r.position = position
}

func (r *Rectangle) Translate(x float64, y float64) {
	r.position.SetX(r.position.X() + x)
	r.position.SetY(r.position.Y() + y)
}

func (r *Rectangle) GetSize() Vector2D {
	return r.size
}

func (r *Rectangle) SetSize(size Vector2D) {
	r.size = size
}

func (r *Rectangle) GetCenter() Vector2D {
	return NewVector2D(r.position.x+r.size.X()/2, r.position.y+r.size.Y()/2)
}

func (r *Rectangle) SetCenter(center Vector2D) {
	r.position.SetX(center.X() - r.size.X()/2)
	r.position.SetY(center.Y() - r.size.Y()/2)
}

func (r *Rectangle) SetCenterX(x float64) {
	r.position.SetX(x - r.size.X()/2)
}

func (r *Rectangle) SetCenterY(y float64) {
	r.position.SetY(y - r.size.Y()/2)
}

func (r *Rectangle) GetLeft() float64 {
	return r.position.X()
}

func (r *Rectangle) SetLeft(value float64) {
	r.position.SetX(value)
}

func (r *Rectangle) GetTop() float64 {
	return r.position.Y()
}

func (r *Rectangle) SetTop(value float64) {
	r.position.SetY(value)
}

func (r *Rectangle) SetRight(value float64) {
	r.position.SetX(value - r.size.X())
}

func (r Rectangle) GetRight() float64 {
	return r.position.X() + r.size.X()
}

func (r *Rectangle) SetBottom(value float64) {
	r.position.SetY(value - r.size.Y())
}

func (r Rectangle) GetBottom() float64 {
	return r.position.Y() + r.size.Y()
}

func (r *Rectangle) Equal(rectangle *Rectangle) bool {
	return (r.position.X() == rectangle.position.X()) && (r.size.X() == rectangle.size.X()) && (r.position.Y() == rectangle.position.Y()) && (r.size.Y() == rectangle.size.Y())
}

func (r *Rectangle) Inflate(x, y float64) {
	r.position = Vector2D{x: r.position.X() - x/2, y: r.position.Y() - y/2}
	r.size = Vector2D{x: r.size.X() + x, y: r.size.Y() + y}
}

// IntersectsRect returns true if this rectangle overlaps other (axis-aligned).
func (r Rectangle) IntersectsRect(other Rectangle) bool {
	return r.GetLeft() < other.GetRight() && r.GetRight() > other.GetLeft() &&
		r.GetTop() < other.GetBottom() && r.GetBottom() > other.GetTop()
}
