package math2D

 

type Rectangle struct {
	position Vector2D
	size     Vector2D
}

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
	r.position.SetX(r.position.X()+ x)
	r.position.SetY(r.position.Y() + y)
}

func (r *Rectangle) GetSize() Vector2D {
	return r.size
}

func (r *Rectangle) SetSize(size Vector2D) {
	r.size = size
}

func (r Rectangle) GetCenter() Vector2D {
	return NewVector2D(r.size.X()/2, r.size.X()/2)
}

func (r Rectangle) Equal(rectangle Rectangle) bool {
	return (r.position.X() == rectangle.position.X()) && (r.size.X() == rectangle.size.X()) && (r.position.Y() == rectangle.position.Y()) && (r.size.Y() == rectangle.size.Y())
}