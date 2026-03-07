package collision

import (
	math2D "github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

func Overlapping(minA, maxA, minB, maxB float64) bool {
	return minB <= maxA && minA <= maxB
}

func OnOneSide(axis math2D.Line, segment math2D.Segment) bool {
	d1 := math2D.SubtractVectors(axis.GetBase(), segment.GetStartPoint())
	d2 := math2D.SubtractVectors(axis.GetBase(), segment.GetEndPoint())

	n := axis.GetDirection().RotateVector90()
	return math2D.DotProduct(n, d1)*math2D.DotProduct(n, d2) > 0
}

func ClampOnRange(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if max < x {
		return max
	}
	return x
}

func ClampOnRectangle(p math2D.Vector2D, r math2D.Rectangle) math2D.Vector2D {
	return math2D.NewVector2D(ClampOnRange(p.X(), r.GetPosition().X(), r.GetPosition().X()+r.GetSize().X()),
		ClampOnRange(p.Y(), r.GetPosition().Y(), r.GetPosition().Y()+r.GetSize().Y()))
}

// segmentOverlapsShape returns true if the world-space segment intersects the shape at the given transform.
func segmentOverlapsShape(seg math2D.Segment, shape CollisionShape, t transform.Transform) bool {
	pos := t.GetPosition()
	switch s := shape.(type) {
	case *CollisionCircle:
		c := math2D.NewCircle(pos, s.circle.GetRadius())
		return CircleSegmentCollide(c, seg)
	case *CollisionRect:
		sz := s.rectangle.GetSize()
		tl := math2D.NewVector2D(pos.X()-sz.X()/2, pos.Y()-sz.Y()/2)
		r := math2D.NewRectangle(tl, sz)
		return RectangleSegmentCollide(r, seg)
	case *CollisionOrientedRect:
		or := math2D.NewOrientedRectangle(pos, s.rectangle.GetHalfExtended(), t.GetRotation())
		return OrientedRectangleSegmentCollide(or, seg)
	case *CollisionPolygon:
		verts := polygonWorldVertices(s.vertices, pos, t.GetRotation())
		return PolygonSegmentCollide(verts, seg)
	default:
		return false
	}
}

// overlapPointShape returns true if the world-space point is inside the shape at the given transform.
func overlapPointShape(point math2D.Vector2D, shape CollisionShape, t transform.Transform) bool {
	pos := t.GetPosition()
	switch s := shape.(type) {
	case *CollisionCircle:
		c := math2D.NewCircle(pos, s.circle.GetRadius())
		return CirclePointCollide(c, point)
	case *CollisionRect:
		sz := s.rectangle.GetSize()
		tl := math2D.NewVector2D(pos.X()-sz.X()/2, pos.Y()-sz.Y()/2)
		r := math2D.NewRectangle(tl, sz)
		return PointRectangleCollide(point, r)
	case *CollisionOrientedRect:
		or := math2D.NewOrientedRectangle(pos, s.rectangle.GetHalfExtended(), t.GetRotation())
		return OrientedRectanglePointCollide(or, point)
	case *CollisionPolygon:
		verts := polygonWorldVertices(s.vertices, pos, t.GetRotation())
		return PointInPolygon(point, verts)
	default:
		return false
	}
}

// overlapCircleShape returns true if the world-space query circle overlaps the shape at the given transform.
func overlapCircleShape(query math2D.Circle, shape CollisionShape, t transform.Transform) bool {
	pos := t.GetPosition()
	switch s := shape.(type) {
	case *CollisionCircle:
		c := math2D.NewCircle(pos, s.circle.GetRadius())
		return CirclesCollide(query, c)
	case *CollisionRect:
		sz := s.rectangle.GetSize()
		tl := math2D.NewVector2D(pos.X()-sz.X()/2, pos.Y()-sz.Y()/2)
		r := math2D.NewRectangle(tl, sz)
		return CircleRectangleCollide(query, r)
	case *CollisionOrientedRect:
		or := math2D.NewOrientedRectangle(pos, s.rectangle.GetHalfExtended(), t.GetRotation())
		return CircleOrientedRectangleCollide(query, or)
	case *CollisionPolygon:
		verts := polygonWorldVertices(s.vertices, pos, t.GetRotation())
		return PolygonCircleCollide(verts, query)
	default:
		return false
	}
}

// overlapRectShape returns true if the world-space query rectangle (AABB) overlaps the shape at the given transform.
func overlapRectShape(query math2D.Rectangle, shape CollisionShape, t transform.Transform) bool {
	pos := t.GetPosition()
	switch s := shape.(type) {
	case *CollisionCircle:
		c := math2D.NewCircle(pos, s.circle.GetRadius())
		return CircleRectangleCollide(c, query)
	case *CollisionRect:
		sz := s.rectangle.GetSize()
		tl := math2D.NewVector2D(pos.X()-sz.X()/2, pos.Y()-sz.Y()/2)
		r := math2D.NewRectangle(tl, sz)
		return RectanglesCollide(query, r)
	case *CollisionOrientedRect:
		or := math2D.NewOrientedRectangle(pos, s.rectangle.GetHalfExtended(), t.GetRotation())
		return OrientedRectangleRectangleCollide(or, query)
	case *CollisionPolygon:
		verts := polygonWorldVertices(s.vertices, pos, t.GetRotation())
		return PolygonRectangleCollide(verts, query)
	default:
		return false
	}
}

// ShapeAABB returns the axis-aligned bounding box (minX, minY, maxX, maxY) of the shape in world space.
// It applies the given transform locally to extract the visual bounds without mutating the shape itself.
func ShapeAABB(shape CollisionShape, t transform.Transform) (minX, minY, maxX, maxY float64) {
	switch s := shape.(type) {
	case *CollisionCircle:
		center := s.circle.GetCenterPosition()
		center.SetX(t.GetPosition().X())
		center.SetY(t.GetPosition().Y())
		r := s.circle.GetRadius()
		return center.X() - r, center.Y() - r, center.X() + r, center.Y() + r
	case *CollisionRect:
		// Pos acts as center in CollisionRect context
		pos := s.rectangle.GetPosition()
		pos.SetX(t.GetPosition().X())
		pos.SetY(t.GetPosition().Y())
		size := s.rectangle.GetSize()
		// Convert center to top-left to compute AABB properly
		topLeftX := pos.X() - size.X()/2
		topLeftY := pos.Y() - size.Y()/2
		return topLeftX, topLeftY, topLeftX + size.X(), topLeftY + size.Y()
	case *CollisionOrientedRect:
		or := math2D.NewOrientedRectangle(t.GetPosition(), s.rectangle.GetHalfExtended(), t.GetRotation())
		hull := OrientedRectangleRectangleHull(or)
		pos := hull.GetPosition()
		sz := hull.GetSize()
		return pos.X(), pos.Y(), pos.X() + sz.X(), pos.Y() + sz.Y()
	case *CollisionPolygon:
		verts := polygonWorldVertices(s.vertices, t.GetPosition(), t.GetRotation())
		if len(verts) == 0 {
			return 0, 0, 0, 0
		}
		minX, minY := verts[0].X(), verts[0].Y()
		maxX, maxY := minX, minY
		for _, v := range verts[1:] {
			if v.X() < minX {
				minX = v.X()
			}
			if v.X() > maxX {
				maxX = v.X()
			}
			if v.Y() < minY {
				minY = v.Y()
			}
			if v.Y() > maxY {
				maxY = v.Y()
			}
		}
		return minX, minY, maxX, maxY
	default:
		return 0, 0, 0, 0
	}
}

func RectangleCorner(r math2D.Rectangle, nr int) math2D.Vector2D {
	corner := r.GetPosition()
	i := nr % 4
	if i == 0 {
		corner.SetX(corner.X() + r.GetSize().X())
	} else if i == 1 {
		corner = math2D.AddVectors(corner, r.GetSize())
	} else if i == 2 {
		corner.SetY(corner.Y() + r.GetSize().Y())
	}
	return corner
}

func OrientedRectangleCorner(r math2D.OrientedRectangle, nr int) math2D.Vector2D {
	c := r.GetHalfExtended()
	i := nr % 4
	if i == 0 {
		c.SetX(-c.X())
	} else if i == 2 {
		c.SetY(-c.Y())
	} else {
		c = c.Negate()
	}
	c = c.RotateVector(r.GetRotation())
	return math2D.AddVectors(c, r.GetCenter())
}

func OrientedRectangleEdge(r math2D.OrientedRectangle, nr int) math2D.Segment {

	a := r.GetHalfExtended()
	b := r.GetHalfExtended()

	i := nr % 4
	if i == 0 {
		a.SetX(-a.X())
	} else if i == 1 {
		b.SetY(-b.Y())
	} else if i == 2 {
		a.SetY(-a.Y())
		b = b.Negate()
	} else {
		a = a.Negate()
		b.SetX(-b.X())
	}

	a = a.RotateVector(r.GetRotation())
	a = math2D.AddVectors(a, r.GetCenter())

	b = b.RotateVector(r.GetRotation())
	b = math2D.AddVectors(b, r.GetCenter())

	return math2D.NewSegment(a, b)
}

func SeparatingAxisForOrientedRectangle(axis math2D.Segment, r math2D.OrientedRectangle) bool {

	rEdge0 := OrientedRectangleEdge(r, 0)
	rEdge2 := OrientedRectangleEdge(r, 2)

	n := math2D.SubtractVectors(axis.GetStartPoint(), axis.GetEndPoint())

	n = n.Normalize()

	axisRange := axis.ProjectSegment(n, true)
	r0Range := rEdge0.ProjectSegment(n, true)
	r2Range := rEdge2.ProjectSegment(n, true)
	rProjection := math2D.RangeHull(r0Range, r2Range)

	return !math2D.OverlappingRanges(axisRange, rProjection)
}

func SeparatingAxisForRectangle(axis math2D.Segment, r math2D.Rectangle) bool {

	n := math2D.SubtractVectors(axis.GetStartPoint(), axis.GetEndPoint())

	n = n.Normalize()

	rEdgeA := math2D.NewSegment(RectangleCorner(r, 0), RectangleCorner(r, 1))
	rEdgeB := math2D.NewSegment(RectangleCorner(r, 2), RectangleCorner(r, 3))

	rEdgeARange := rEdgeA.ProjectSegment(n, true)
	rEdgeBRange := rEdgeB.ProjectSegment(n, true)
	rProjection := math2D.RangeHull(rEdgeARange, rEdgeBRange)

	axisRange := axis.ProjectSegment(n, true)

	return !math2D.OverlappingRanges(axisRange, rProjection)
}
