package collision

type shapeKind int

const (
	kindCircle shapeKind = iota
	kindRect
)

// typedShape extends CollisionShape with a kind identifier for dispatcher routing.
type typedShape interface {
	CollisionShape
	shapeKind() shapeKind
}

// ShapeCollides delegates overlap detection to the registered handler for the shape pair.
// Returns false if either shape is not a typedShape or no handler is registered.
func ShapeCollides(a, b CollisionShape) bool {
	ta, okA := a.(typedShape)
	tb, okB := b.(typedShape)
	if !okA || !okB {
		return false
	}
	handler := collisionHandlers[ta.shapeKind()][tb.shapeKind()]
	if handler == nil {
		return false
	}
	return handler(a, b)
}

var collisionHandlers = map[shapeKind]map[shapeKind]func(CollisionShape, CollisionShape) bool{
	kindCircle: {
		kindCircle: func(a, b CollisionShape) bool {
			return CirclesCollide(a.(*CollisionCircle).circle, b.(*CollisionCircle).circle)
		},
		kindRect: func(a, b CollisionShape) bool {
			return CircleRectangleCollide(a.(*CollisionCircle).circle, b.(*CollisionRect).rectangle)
		},
	},
	kindRect: {
		kindCircle: func(a, b CollisionShape) bool {
			return CircleRectangleCollide(b.(*CollisionCircle).circle, a.(*CollisionRect).rectangle)
		},
		kindRect: func(a, b CollisionShape) bool {
			return RectanglesCollide(a.(*CollisionRect).rectangle, b.(*CollisionRect).rectangle)
		},
	},
}
