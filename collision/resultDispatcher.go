package collision

// ShapeCollisionResult returns the MTV for two shapes, or (zero result, false) if unsupported.
// Shapes must have UpdateTransform called with world transform before calling this.
func ShapeCollisionResult(a, b CollisionShape) (CollisionResult, bool) {
	ta, okA := a.(typedShape)
	tb, okB := b.(typedShape)
	if !okA || !okB {
		return CollisionResult{}, false
	}
	handler := resultHandlers[ta.shapeKind()][tb.shapeKind()]
	if handler == nil {
		return CollisionResult{}, false
	}
	return handler(a, b), true
}

var resultHandlers = map[shapeKind]map[shapeKind]func(CollisionShape, CollisionShape) CollisionResult{
	kindCircle: {
		kindCircle: func(a, b CollisionShape) CollisionResult {
			return CirclesCollideResult(a.(*CollisionCircle).circle, b.(*CollisionCircle).circle)
		},
		kindRect: func(a, b CollisionShape) CollisionResult {
			return CircleRectangleCollideResult(a.(*CollisionCircle).circle, b.(*CollisionRect).rectangle)
		},
	},
	kindRect: {
		kindCircle: func(a, b CollisionShape) CollisionResult {
			// Circle vs Rect: swap so circle is first, invert normal for consistency
			res := CircleRectangleCollideResult(b.(*CollisionCircle).circle, a.(*CollisionRect).rectangle)
			if res.Overlapping {
				res.Normal = res.Normal.Negate()
			}
			return res
		},
		kindRect: func(a, b CollisionShape) CollisionResult {
			return RectanglesCollideResult(a.(*CollisionRect).rectangle, b.(*CollisionRect).rectangle)
		},
	},
}
