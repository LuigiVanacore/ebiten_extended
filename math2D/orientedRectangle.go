package math2D
 

type OrientedRectangle struct {
	center       Vector2D
	halfExtended Vector2D
	rotation     float64
}

func (orientedrectangle *OrientedRectangle) GetCenter() Vector2D {
	return orientedrectangle.center
}

func (orientedrectangle *OrientedRectangle) GetHalfExtended() Vector2D {
	return orientedrectangle.halfExtended
}

func (orientedrectangle *OrientedRectangle) GetRotation() float64 {
	return orientedrectangle.rotation
}

func (orientedrectangle *OrientedRectangle) SetCenter(center Vector2D) *OrientedRectangle {
	orientedrectangle.center = center
	return orientedrectangle
}

func (orientedrectangle *OrientedRectangle) SetHalfExtended(halfExtended  Vector2D) *OrientedRectangle {
	orientedrectangle.halfExtended = halfExtended
	return orientedrectangle
}

func (orientedrectangle *OrientedRectangle) SetRotation(rotation float64) *OrientedRectangle {
	orientedrectangle.rotation = rotation
	return orientedrectangle
}