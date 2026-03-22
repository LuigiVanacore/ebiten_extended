package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type shapeKind int

const (
	kindCircle shapeKind = iota
	kindRect
	kindOrientedRect
	kindPolygon
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
			caCopy := math2D.NewCircle(tA.GetPosition(), ca.GetRadius())
			cbCopy := math2D.NewCircle(tB.GetPosition(), cb.GetRadius())
			return CirclesCollide(caCopy, cbCopy)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ca := a.(*CollisionCircle).circle
			rb := b.(*CollisionRect).rectangle
			caCopy := math2D.NewCircle(tA.GetPosition(), ca.GetRadius())
			rbCopy := math2D.NewRectangle(rb.GetPosition(), rb.GetSize())
			rbCopy.SetCenter(tB.GetPosition())
			return CircleRectangleCollide(caCopy, rbCopy)
		},
		kindOrientedRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ca := a.(*CollisionCircle).circle
			ob := b.(*CollisionOrientedRect).rectangle
			caCopy := math2D.NewCircle(tA.GetPosition(), ca.GetRadius())
			obCopy := math2D.NewOrientedRectangle(tB.GetPosition(), ob.GetHalfExtended(), tB.GetRotation())
			return CircleOrientedRectangleCollide(caCopy, obCopy)
		},
		kindPolygon: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ca := math2D.NewCircle(tA.GetPosition(), a.(*CollisionCircle).circle.GetRadius())
			vb := polygonWorldVertices(b.(*CollisionPolygon).vertices, tB.GetPosition(), tB.GetRotation())
			return PolygonCircleCollide(vb, ca)
		},
	},
	kindRect: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ra := a.(*CollisionRect).rectangle
			cb := b.(*CollisionCircle).circle
			raCopy := math2D.NewRectangle(ra.GetPosition(), ra.GetSize())
			raCopy.SetCenter(tA.GetPosition())
			cbCopy := math2D.NewCircle(tB.GetPosition(), cb.GetRadius())
			return CircleRectangleCollide(cbCopy, raCopy)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ra := a.(*CollisionRect).rectangle
			rb := b.(*CollisionRect).rectangle
			raCopy := math2D.NewRectangle(ra.GetPosition(), ra.GetSize())
			rbCopy := math2D.NewRectangle(rb.GetPosition(), rb.GetSize())
			raCopy.SetCenter(tA.GetPosition())
			rbCopy.SetCenter(tB.GetPosition())
			return RectanglesCollide(raCopy, rbCopy)
		},
		kindOrientedRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ra := a.(*CollisionRect).rectangle
			ob := b.(*CollisionOrientedRect).rectangle
			raCopy := math2D.NewRectangle(ra.GetPosition(), ra.GetSize())
			raCopy.SetCenter(tA.GetPosition())
			obCopy := math2D.NewOrientedRectangle(tB.GetPosition(), ob.GetHalfExtended(), tB.GetRotation())
			return OrientedRectangleRectangleCollide(obCopy, raCopy)
		},
		kindPolygon: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			ra := math2D.NewRectangle(a.(*CollisionRect).rectangle.GetPosition(), a.(*CollisionRect).rectangle.GetSize())
			ra.SetCenter(tA.GetPosition())
			vb := polygonWorldVertices(b.(*CollisionPolygon).vertices, tB.GetPosition(), tB.GetRotation())
			return PolygonRectangleCollide(vb, ra)
		},
	},
	kindOrientedRect: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			oa := a.(*CollisionOrientedRect).rectangle
			cb := b.(*CollisionCircle).circle
			oaCopy := math2D.NewOrientedRectangle(tA.GetPosition(), oa.GetHalfExtended(), tA.GetRotation())
			cbCopy := math2D.NewCircle(tB.GetPosition(), cb.GetRadius())
			return CircleOrientedRectangleCollide(cbCopy, oaCopy)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			oa := a.(*CollisionOrientedRect).rectangle
			rb := b.(*CollisionRect).rectangle
			oaCopy := math2D.NewOrientedRectangle(tA.GetPosition(), oa.GetHalfExtended(), tA.GetRotation())
			rbCopy := math2D.NewRectangle(rb.GetPosition(), rb.GetSize())
			rbCopy.SetCenter(tB.GetPosition())
			return OrientedRectangleRectangleCollide(oaCopy, rbCopy)
		},
		kindOrientedRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			oa := a.(*CollisionOrientedRect).rectangle
			ob := b.(*CollisionOrientedRect).rectangle
			oaCopy := math2D.NewOrientedRectangle(tA.GetPosition(), oa.GetHalfExtended(), tA.GetRotation())
			obCopy := math2D.NewOrientedRectangle(tB.GetPosition(), ob.GetHalfExtended(), tB.GetRotation())
			return OrientedRectanglesCollide(oaCopy, obCopy)
		},
		kindPolygon: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			oa := a.(*CollisionOrientedRect).rectangle
			or := math2D.NewOrientedRectangle(tA.GetPosition(), oa.GetHalfExtended(), tA.GetRotation())
			va := orientedRectToVertices(or)
			vb := polygonWorldVertices(b.(*CollisionPolygon).vertices, tB.GetPosition(), tB.GetRotation())
			return PolygonsCollide(va, vb)
		},
	},
	kindPolygon: {
		kindCircle: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			va := polygonWorldVertices(a.(*CollisionPolygon).vertices, tA.GetPosition(), tA.GetRotation())
			cb := math2D.NewCircle(tB.GetPosition(), b.(*CollisionCircle).circle.GetRadius())
			return PolygonCircleCollide(va, cb)
		},
		kindRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			va := polygonWorldVertices(a.(*CollisionPolygon).vertices, tA.GetPosition(), tA.GetRotation())
			rb := b.(*CollisionRect).rectangle
			rbCopy := math2D.NewRectangle(rb.GetPosition(), rb.GetSize())
			rbCopy.SetCenter(tB.GetPosition())
			return PolygonRectangleCollide(va, rbCopy)
		},
		kindOrientedRect: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			va := polygonWorldVertices(a.(*CollisionPolygon).vertices, tA.GetPosition(), tA.GetRotation())
			ob := b.(*CollisionOrientedRect).rectangle
			or := math2D.NewOrientedRectangle(tB.GetPosition(), ob.GetHalfExtended(), tB.GetRotation())
			vb := orientedRectToVertices(or)
			return PolygonsCollide(va, vb)
		},
		kindPolygon: func(a CollisionShape, tA transform.Transform, b CollisionShape, tB transform.Transform) bool {
			va := polygonWorldVertices(a.(*CollisionPolygon).vertices, tA.GetPosition(), tA.GetRotation())
			vb := polygonWorldVertices(b.(*CollisionPolygon).vertices, tB.GetPosition(), tB.GetRotation())
			return PolygonsCollide(va, vb)
		},
	},
}

func orientedRectToVertices(or math2D.OrientedRectangle) []math2D.Vector2D {
	return []math2D.Vector2D{
		OrientedRectangleCorner(or, 0),
		OrientedRectangleCorner(or, 1),
		OrientedRectangleCorner(or, 2),
		OrientedRectangleCorner(or, 3),
	}
}
