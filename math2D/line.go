package math2D
 

type Line struct {
	base      Vector2D
	direction Vector2D
}

func NewLine(base, direction Vector2D) Line {
	return Line{base: base, direction: direction}
}

func (line *Line) GetBase() Vector2D {
	return line.base
}

func (line *Line) GetDirection() Vector2D {
	return line.direction
}

func (line *Line) SetBase(base Vector2D) *Line {
	line.base = base
	return line
}

func (line *Line) SetDirection(direction Vector2D) *Line {
	line.direction = direction
	return line
}