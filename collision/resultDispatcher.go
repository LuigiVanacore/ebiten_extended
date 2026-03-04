package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// ShapeCollisionResult returns the MTV for two shapes, or (zero result, false) if unsupported.
// Shapes are evaluated virtually using their provided world transforms.
func ShapeCollisionResult(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) (CollisionResult, bool) {
	ta, okA := a.(typedShape)
	tb, okB := b.(typedShape)
	if !okA || !okB {
		return CollisionResult{}, false
	}
	handler := resultHandlers[ta.shapeKind()][tb.shapeKind()]
	if handler == nil {
		return CollisionResult{}, false
	}
	return handler(a, tA, b, tB), true
}

var resultHandlers = map[shapeKind]map[shapeKind]func(CollisionShape, transform.Transform, CollisionShape, transform.Transform) CollisionResult{
	kindCircle: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ca := a.(*CollisionCircle).circle
			cb := b.(*CollisionCircle).circle
			ca.SetCenter(tA.GetPosition())
			cb.SetCenter(tB.GetPosition())
			return CirclesCollideResult(ca, cb)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ca := a.(*CollisionCircle).circle
			rb := b.(*CollisionRect).rectangle
			ca.SetCenter(tA.GetPosition())
			rb.SetCenter(tB.GetPosition())
			return CircleRectangleCollideResult(ca, rb)
		},
	},
	kindRect: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ra := a.(*CollisionRect).rectangle
			cb := b.(*CollisionCircle).circle
			ra.SetCenter(tA.GetPosition())
			cb.SetCenter(tB.GetPosition())

			// Circle vs Rect: swap so circle is first, invert normal for consistency
			res := CircleRectangleCollideResult(cb, ra)
			if res.Overlapping {
				res.Normal = res.Normal.Negate()
			}
			return res
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ra := a.(*CollisionRect).rectangle
			rb := b.(*CollisionRect).rectangle
			ra.SetCenter(tA.GetPosition())
			rb.SetCenter(tB.GetPosition())
			return RectanglesCollideResult(ra, rb)
		},
	},
}
