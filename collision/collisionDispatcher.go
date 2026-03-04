package collision

import "github.com/LuigiVanacore/ebiten_extended/transform"

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
func ShapeCollides(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
	ta, okA := a.(typedShape)
	tb, okB := b.(typedShape)
	if !okA || !okB {
		return false
	}
	handler := collisionHandlers[ta.shapeKind()][tb.shapeKind()]
	if handler == nil {
		return false
	}
	return handler(a, tA, b, tB)
}

var collisionHandlers = map[shapeKind]map[shapeKind]func(CollisionShape, transform.Transform, CollisionShape, transform.Transform) bool{
	kindCircle: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ca := a.(*CollisionCircle).circle
			cb := b.(*CollisionCircle).circle
			ca.SetCenter(tA.GetPosition())
			cb.SetCenter(tB.GetPosition())
			return CirclesCollide(ca, cb)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ca := a.(*CollisionCircle).circle
			rb := b.(*CollisionRect).rectangle
			ca.SetCenter(tA.GetPosition())
			rb.SetCenter(tB.GetPosition())
			return CircleRectangleCollide(ca, rb)
		},
	},
	kindRect: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ra := a.(*CollisionRect).rectangle
			cb := b.(*CollisionCircle).circle
			ra.SetCenter(tA.GetPosition())
			cb.SetCenter(tB.GetPosition())
			return CircleRectangleCollide(cb, ra)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ra := a.(*CollisionRect).rectangle
			rb := b.(*CollisionRect).rectangle
			ra.SetCenter(tA.GetPosition())
			rb.SetCenter(tB.GetPosition())
			return RectanglesCollide(ra, rb)
		},
	},
}
