package ebiten_extended

type Rect struct {
	transform Transform
	Position  Vector2D
	Width     float64
	Height    float64
}

func NewRect(x float64, y float64, width float64, height float64) *Rect {
	return &Rect{
		Position: Vector2D{X: x, Y: y},
		Width:    width,
		Height:   height,
	}
}

func (r *Rect) SetTransform(transform Transform) {
	r.transform = transform
}

func (r *Rect) GetPosition() (float64, float64) {
	return r.Position.X, r.Position.Y
}

func (r *Rect) SetPosition(x float64, y float64) {
	r.Position.X = x
	r.Position.Y = y
}

func (r *Rect) Translate(x float64, y float64) {
	r.Position.X += x
	r.Position.Y += y
}

func (r *Rect) GetSize() float64 {
	return r.Width + r.Height
}

func (r *Rect) GetCenter() (float64, float64) {
	return r.Width / 2, r.Height / 2
}
func (r *Rect) Equal(rectangle Rect) bool {
	return (r.Position.X == rectangle.Position.X) && (r.Width == rectangle.Width) && (r.Position.Y == rectangle.Position.Y) && (r.Height == rectangle.Height)
}

//
//func (r *Rect) Contains(x float64, y float64) bool {
//	minX := Min(r.Position.X, r.Position.X+r.Width)
//	maxX := Max(r.Position.X, r.Position.X+r.Width)
//	minY := Min(r.Position.Y, r.Position.Y+r.Height)
//	maxY := Max(r.Position.Y, r.Position.Y+r.Height)
//	return (x >= minX) && (x < maxX) && (y >= minY) && (y < maxY)
//}
//
//func (r *Rect) Intersect(shape Shape) bool {
//	switch shape.(type) {
//	case *Rect:
//		return r.Intersect(shape)
//		break
//	}
//	return false
//}
//
//func (r *Rect) IntersectRect(rectangle *Rect) bool {
//
//	r1MinX := Min(r.Position.X, r.Position.X+r.Width)
//	r1MaxX := Max(r.Position.X, r.Position.X+r.Width)
//	r1MinY := Min(r.Position.Y, r.Position.Y+r.Height)
//	r1MaxY := Max(r.Position.Y, r.Position.Y+r.Height)
//
//	r2MinX := Min(rectangle.Position.X, rectangle.Position.X+rectangle.Width)
//	r2MaxX := Max(rectangle.Position.X, rectangle.Position.X+rectangle.Width)
//	r2MinY := Min(rectangle.Position.Y, rectangle.Position.X+rectangle.Width)
//	r2MaxY := Min(rectangle.Position.X, rectangle.Position.X+rectangle.Width)
//
//	interLeft := Min(r1MinX, r2MinX)
//	interTop := Max(r1MinY, r2MinY)
//	interRight := Min(r1MaxX, r2MaxX)
//	interBottom := Max(r1MaxY, r2MaxY)
//
//	if (interLeft < interRight) && (interTop < interBottom) {
//		return true
//	}
//
//	return false
//
//}
