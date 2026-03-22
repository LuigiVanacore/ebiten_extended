package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

// CollisionPolygon is a convex polygon shape. Vertices are in local space relative to the collider position.
// Vertices should be ordered counter-clockwise. Position acts as the polygon's origin.
type CollisionPolygon struct {
	vertices []math2D.Vector2D
}

// NewCollisionPolygon creates a polygon from local-space vertices (relative to collider position).
// Vertices should be ordered counter-clockwise for correct SAT.
func NewCollisionPolygon(vertices []math2D.Vector2D) *CollisionPolygon {
	if len(vertices) < 3 {
		return nil
	}
	// Copy to avoid external mutation
	local := make([]math2D.Vector2D, len(vertices))
	for i, v := range vertices {
		local[i] = math2D.NewVector2D(v.X(), v.Y())
	}
	return &CollisionPolygon{vertices: local}
}

// GetVertices returns a copy of the local vertices.
func (c *CollisionPolygon) GetVertices() []math2D.Vector2D {
	out := make([]math2D.Vector2D, len(c.vertices))
	for i, v := range c.vertices {
		out[i] = math2D.NewVector2D(v.X(), v.Y())
	}
	return out
}

func (c *CollisionPolygon) shapeKind() shapeKind {
	return kindPolygon
}

func (c *CollisionPolygon) IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool {
	return ShapeCollides(c, tSelf, other, tOther)
}

// polygonWorldVertices returns the polygon vertices in world space.
func polygonWorldVertices(vertices []math2D.Vector2D, pos math2D.Vector2D, rotation float64) []math2D.Vector2D {
	out := make([]math2D.Vector2D, len(vertices))
	for i, v := range vertices {
		rotated := v.RotateVector(rotation)
		out[i] = math2D.AddVectors(pos, rotated)
	}
	return out
}

// projectPolygon projects vertices onto the axis (unit vector) and returns the min/max as Range.
func projectPolygon(vertices []math2D.Vector2D, axis math2D.Vector2D) math2D.Range {
	if len(vertices) == 0 {
		return math2D.Range{}
	}
	dot := math2D.DotProduct(vertices[0], axis)
	minV, maxV := dot, dot
	for i := 1; i < len(vertices); i++ {
		dot = math2D.DotProduct(vertices[i], axis)
		if dot < minV {
			minV = dot
		}
		if dot > maxV {
			maxV = dot
		}
	}
	r := math2D.NewRange(minV, maxV)
	return r.SortRange()
}

// polygonEdgeNormals returns the outward-facing normals for each edge (counter-clockwise vertices).
func polygonEdgeNormals(vertices []math2D.Vector2D) []math2D.Vector2D {
	n := len(vertices)
	if n < 3 {
		return nil
	}
	normals := make([]math2D.Vector2D, n)
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		edge := math2D.SubtractVectors(vertices[next], vertices[i])
		normal := edge.RotateVector90()
		normals[i] = normal.Normalize()
	}
	return normals
}

// PolygonsCollide returns true if two convex polygons overlap (SAT).
func PolygonsCollide(vertsA, vertsB []math2D.Vector2D) bool {
	if len(vertsA) < 3 || len(vertsB) < 3 {
		return false
	}
	normalsA := polygonEdgeNormals(vertsA)
	normalsB := polygonEdgeNormals(vertsB)
	for _, n := range normalsA {
		projA := projectPolygon(vertsA, n)
		projB := projectPolygon(vertsB, n)
		if !math2D.OverlappingRanges(projA, projB) {
			return false
		}
	}
	for _, n := range normalsB {
		projA := projectPolygon(vertsA, n)
		projB := projectPolygon(vertsB, n)
		if !math2D.OverlappingRanges(projA, projB) {
			return false
		}
	}
	return true
}

// PolygonCircleCollide returns true if a convex polygon and circle overlap.
func PolygonCircleCollide(verts []math2D.Vector2D, c math2D.Circle) bool {
	if len(verts) < 3 {
		return false
	}
	center := c.GetCenterPosition()
	radius := c.GetRadius()
	for i := 0; i < len(verts); i++ {
		next := (i + 1) % len(verts)
		edge := math2D.SubtractVectors(verts[next], verts[i])
		normal := edge.RotateVector90().Normalize()
		proj := projectPolygon(verts, normal)
		centerProj := math2D.DotProduct(center, normal)
		if centerProj+radius < proj.GetMinimum() || centerProj-radius > proj.GetMaximum() {
			return false
		}
	}
	closestDist := 1e18
	for i := 0; i < len(verts); i++ {
		next := (i + 1) % len(verts)
		seg := math2D.NewSegment(verts[i], verts[next])
		clamped := ClampOnSegment(center, seg)
		d := math2D.SubtractVectors(center, clamped)
		distSq := math2D.DotProduct(d, d)
		if distSq < closestDist {
			closestDist = distSq
		}
	}
	return closestDist <= radius*radius
}

// ClampOnSegment returns the closest point on the segment to p.
func ClampOnSegment(p math2D.Vector2D, seg math2D.Segment) math2D.Vector2D {
	a := seg.GetStartPoint()
	b := seg.GetEndPoint()
	ab := math2D.SubtractVectors(b, a)
	ap := math2D.SubtractVectors(p, a)
	t := math2D.DotProduct(ap, ab) / (math2D.DotProduct(ab, ab) + 1e-12)
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	return math2D.AddVectors(a, ab.MultiplyScalar(t))
}

// PolygonRectangleCollide returns true if polygon and axis-aligned rectangle overlap.
func PolygonRectangleCollide(verts []math2D.Vector2D, r math2D.Rectangle) bool {
	pos := r.GetPosition()
	sz := r.GetSize()
	rectVerts := []math2D.Vector2D{
		pos,
		math2D.NewVector2D(pos.X()+sz.X(), pos.Y()),
		math2D.AddVectors(pos, sz),
		math2D.NewVector2D(pos.X(), pos.Y()+sz.Y()),
	}
	return PolygonsCollide(verts, rectVerts)
}

// PointInPolygon returns true if the point is inside the convex polygon.
func PointInPolygon(p math2D.Vector2D, verts []math2D.Vector2D) bool {
	if len(verts) < 3 {
		return false
	}
	normals := polygonEdgeNormals(verts)
	for i := 0; i < len(verts); i++ {
		toPoint := math2D.SubtractVectors(p, verts[i])
		if math2D.DotProduct(toPoint, normals[i]) > 0 {
			return false
		}
	}
	return true
}

// PolygonSegmentCollide returns true if the segment intersects the polygon.
func PolygonSegmentCollide(verts []math2D.Vector2D, seg math2D.Segment) bool {
	if len(verts) < 3 {
		return false
	}
	if PointInPolygon(seg.GetStartPoint(), verts) || PointInPolygon(seg.GetEndPoint(), verts) {
		return true
	}
	for i := 0; i < len(verts); i++ {
		next := (i + 1) % len(verts)
		edge := math2D.NewSegment(verts[i], verts[next])
		if SegmentCollide(seg, edge) {
			return true
		}
	}
	return false
}
