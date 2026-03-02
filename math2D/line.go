package math2D

// Line represents an infinite 2D line with a base point and a direction vector.
type Line struct {
	base      Vector2D
	direction Vector2D
}

// NewLine returns a line through base with the given direction.
func NewLine(base, direction Vector2D) Line {
	return Line{base: base, direction: direction}
}

// GetBase returns the base point of the line.
func (line *Line) GetBase() Vector2D {
	return line.base
}

// GetDirection returns the direction vector of the line.
func (line *Line) GetDirection() Vector2D {
	return line.direction
}

// SetBase sets the base point of the line.
func (line *Line) SetBase(base Vector2D) *Line {
	line.base = base
	return line
}

// SetDirection sets the direction vector of the line.
func (line *Line) SetDirection(direction Vector2D) *Line {
	line.direction = direction
	return line
}