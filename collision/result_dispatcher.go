package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
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
			caCopy := math2D.NewCircle(tA.GetPosition(), ca.GetRadius())
			cbCopy := math2D.NewCircle(tB.GetPosition(), cb.GetRadius())
			return CirclesCollideResult(caCopy, cbCopy)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ca := a.(*CollisionCircle).circle
			rb := b.(*CollisionRect).rectangle
			caCopy := math2D.NewCircle(tA.GetPosition(), ca.GetRadius())
			rbCopy := math2D.NewRectangle(rb.GetPosition(), rb.GetSize())
			rbCopy.SetCenter(tB.GetPosition())
			return CircleRectangleCollideResult(caCopy, rbCopy)
		},
	},
	kindRect: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ra := a.(*CollisionRect).rectangle
			cb := b.(*CollisionCircle).circle
			raCopy := math2D.NewRectangle(ra.GetPosition(), ra.GetSize())
			raCopy.SetCenter(tA.GetPosition())
			cbCopy := math2D.NewCircle(tB.GetPosition(), cb.GetRadius())

			// Circle vs Rect: swap so circle is first, invert normal for consistency
			res := CircleRectangleCollideResult(cbCopy, raCopy)
			if res.Overlapping {
				res.Normal = res.Normal.Negate()
			}
			return res
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) CollisionResult {
			ra := a.(*CollisionRect).rectangle
			rb := b.(*CollisionRect).rectangle
			raCopy := math2D.NewRectangle(ra.GetPosition(), ra.GetSize())
			rbCopy := math2D.NewRectangle(rb.GetPosition(), rb.GetSize())
			raCopy.SetCenter(tA.GetPosition())
			rbCopy.SetCenter(tB.GetPosition())
			return RectanglesCollideResult(raCopy, rbCopy)
		},
	},
}
