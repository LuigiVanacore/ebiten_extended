package math2D

 

type Rectangle struct {
	position Vector2D
	size     Vector2D
}

func NewRectangle(x float64, y float64, width float64, height float64) Rectangle {
	return Rectangle{
		position: Vector2D{x: x, y: y},
		size:     Vector2D{x: x, y: y},
	}
}

func (r *Rectangle) GetPosition() Vector2D {
	return r.position
}

func (r *Rectangle) SetPosition(x, y float64) {
	r.position.x = x
	r.position.y = y
}

func (r *Rectangle) Translate(x float64, y float64) {
	r.position.x += x
	r.position.y += y
}

func (r *Rectangle) GetSize() Vector2D {
	return r.size
}

func (r *Rectangle) SetSize(x, y float64) {
	r.size.x = x
	r.size.y = y
}

func (r Rectangle) GetCenter()  Vector2D {
	return  NewVector2D(r.size.X() / 2, r.size.X() / 2 )
}

func (r Rectangle) Equal(rectangle Rectangle) bool {
	return (r.position.X() == rectangle.position.X()) && (r.size.X() == rectangle.size.X()) && (r.position.Y() == rectangle.position.Y()) && (r.size.Y() == rectangle.size.Y())
}
 