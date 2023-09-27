package collision
 
import (
	math2D "github.com/LuigiVanacore/ebiten_extended/math2D"
)

func Overlapping(minA, maxA, minB, maxB float64) bool {
	return minB <= maxA && minA <= maxB
}


 


func OnOneSide(axis math2D.Line, segment math2D.Segment) bool{
	d1 := math2D.SubtractVectors( axis.GetBase(), segment.GetStartPoint())
	d2 :=  math2D.SubtractVectors( axis.GetBase(), segment.GetEndPoint())
	
	n := axis.GetDirection().RotateVector90()
	return math2D.DotProduct(n, d1) * math2D.DotProduct(n, d2) > 0
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
	return math2D.NewVector2D( ClampOnRange(p.X(), r.GetPosition().X(), r.GetPosition().X() + r.GetSize().X()),
								ClampOnRange(p.Y(), r.GetPosition().Y(),r.GetPosition().Y() + r.GetSize().Y()))
}

 

func RectangleCorner(r math2D.Rectangle, nr int) math2D.Vector2D {
	corner := r.GetPosition()
	i := nr % 4
	if i == 0 {
		corner.SetX(corner.X() + r.GetSize().X())
	} else 	if i == 1 {
		corner = math2D.AddVectors(corner, r.GetSize())
	} else if i == 2 {
		corner.SetY(corner.Y() + r.GetSize().Y())
	}
	return corner
}

 
func OrientedRectangleCorner(r math2D.OrientedRectangle, nr int ) math2D.Vector2D {
	c := r.GetHalfExtended()
	i := nr % 4
	if i == 0 {
		c.SetX(-c.X())
	} else if i== 2 {
		c.SetY(-c.Y())
	} else {
		c = c.Negate()
	}
	c = c.RotateVector(r.GetRotation())
	return math2D.AddVectors(c, r.GetCenter())
}

 

func OrientedRectangleEdge( r math2D.OrientedRectangle, nr int) math2D.Segment {

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
 

 
func SeparatingAxisForOrientedRectangle( axis math2D.Segment, r math2D.OrientedRectangle) bool {

	rEdge0 := OrientedRectangleEdge(r, 0)
	rEdge2 := OrientedRectangleEdge(r, 2)

	n := math2D.SubtractVectors(axis.GetStartPoint(), axis.GetEndPoint())

	n = n.UnitVector2D()

	axisRange := axis.ProjectSegment(n, true)
	r0Range := rEdge0.ProjectSegment(n, true)
	r2Range := rEdge2.ProjectSegment(n, true)
	rProjection := math2D.RangeHull(r0Range, r2Range)

	return !math2D.OverlappingRanges(axisRange, rProjection)
}


func SeparatingAxisForRectangle( axis math2D.Segment, r math2D.Rectangle) bool {

	n := math2D.SubtractVectors(axis.GetStartPoint(), axis.GetEndPoint())

	n = n.UnitVector2D()

	rEdgeA := math2D.NewSegment(RectangleCorner(r, 0), RectangleCorner(r, 1))
	rEdgeB := math2D.NewSegment(RectangleCorner(r, 2), RectangleCorner(r, 3))

	rEdgeARange := rEdgeA.ProjectSegment(n, true)
	rEdgeBRange := rEdgeB.ProjectSegment(n, true)
	rProjection := math2D.RangeHull(rEdgeARange, rEdgeBRange)

	axisRange := axis.ProjectSegment(n, true)

	return !math2D.OverlappingRanges(axisRange, rProjection)
}

