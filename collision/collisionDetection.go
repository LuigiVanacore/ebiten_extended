package collision

import "github.com/LuigiVanacore/ebiten_extended/math2D"
 

func CirclesCollide(a, b math2D.Circle) bool {
	radiusSum := a.GetRadius() + b.GetRadius()
	distance := math2D.SubtractVectors(a.GetCenterPosition(), b.GetCenterPosition())
	return math2D.DotProduct(distance, distance) <= radiusSum * radiusSum
}

 
func CirclePointCollide(c math2D.Circle, p math2D.Vector2D ) bool{
	distance := math2D.SubtractVectors(c.GetCenterPosition(), p)
	radius := c.GetRadius()
	return math2D.DotProduct(distance, distance) <= radius * radius
}

 
func CircleLineCollide(c math2D.Circle, l math2D.Line ) bool {
	lc := math2D.SubtractVectors(c.GetCenterPosition(), l.GetBase())
	p := lc.ProjectVector(l.GetDirection())
	p = math2D.AddVectors(l.GetBase(), p)
	return CirclePointCollide(c, p) 
}

 
func CircleRectangleCollide(c math2D.Circle, r math2D.Rectangle) bool {
	clamped := ClampOnRectangle(c.GetCenterPosition(), r)
	return CirclePointCollide(c, clamped)
}

 
func CircleSegmentCollide(c math2D.Circle, s math2D.Segment) bool {
	if CirclePointCollide(c, s.GetStartPoint()) || CirclePointCollide(c, s.GetEndPoint()) {
		return true
	}
	d := math2D.SubtractVectors(s.GetEndPoint(), s.GetStartPoint())
	lc := math2D.SubtractVectors(c.GetCenterPosition(), s.GetStartPoint())
	p := lc.ProjectVector(d)
	nearest := math2D.AddVectors(s.GetStartPoint(), p)
	return CirclePointCollide(c, nearest) && math2D.DotProduct(p, p) <= math2D.DotProduct(d, d) && 0 <= math2D.DotProduct(p,p)

}


func CircleOrientedRectangleCollide(c math2D.Circle, r math2D.OrientedRectangle) bool {
	lr := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(r.GetHalfExtended().X() * 2, r.GetHalfExtended().Y() * 2))
	lc := math2D.NewCircle(math2D.NewVector2D(0,0), c.GetRadius() )

	lc.SetCenter(math2D.SubtractVectors(c.GetCenterPosition(), r.GetCenter()))
	lc.SetCenter(lc.GetCenterPosition().RotateVector(-r.GetRotation()))
	lc.SetCenter(math2D.AddVectors(lc.GetCenterPosition(), r.GetHalfExtended()))

	return CircleRectangleCollide(lc, lr)
}

 
func RectanglesCollide(a, b math2D.Rectangle) bool {
	return Overlapping(a.GetPosition().X(), a.GetPosition().X() + a.GetSize().X(), b.GetPosition().X(),
					   b.GetPosition().X() + b.GetSize().X()) && Overlapping(a.GetPosition().Y(), 
					   															a.GetPosition().Y() + a.GetSize().Y(), 
																				b.GetPosition().Y(), 
																				b.GetPosition().Y() + b.GetSize().Y())
}

 

func PointsCollide(a, b math2D.Vector2D) bool {
	return a.X() == b.X() && a.Y() == b.Y()
}

 
func LinePointCollide(l math2D.Line, p math2D.Vector2D) bool {
	lp := math2D.SubtractVectors(p, l.GetBase())
	return ( lp.X() == 0 && lp.Y() == 0 ) || lp.IsParallel(l.GetDirection()) 
}


func LinesCollide(a, b math2D.Line) bool {
	return !a.GetDirection().IsParallel(b.GetDirection()) || LinePointCollide(a, b.GetBase())
}

 

func LineSegmentCollide(l math2D.Line, s math2D.Segment) bool {
	return !OnOneSide(l, s)
}

 
func PointSegmentCollide(p math2D.Vector2D, s math2D.Segment) bool {
	d := math2D.SubtractVectors(s.GetEndPoint(), s.GetStartPoint())
	lp := math2D.SubtractVectors(p, s.GetStartPoint())
	pr := lp.ProjectVector(d)
	return PointsCollide(lp, pr) && math2D.DotProduct(pr, pr) <= math2D.DotProduct(d, d) && 0 <= math2D.DotProduct(pr, d)
}




func SegmentCollide(a, b math2D.Segment) bool {

	axisA := math2D.NewLine(a.GetStartPoint(), math2D.SubtractVectors(a.GetEndPoint(), a.GetStartPoint()))
	
	if ( 0 == axisA.GetDirection().X() && 0 == axisA.GetDirection().Y()) {
		return PointSegmentCollide(a.GetStartPoint(), b)
	} 
	if OnOneSide(axisA, b) {
		return false
	}

	axisB := math2D.NewLine(b.GetStartPoint(), math2D.SubtractVectors(b.GetEndPoint(), b.GetStartPoint())) 
	if ( 0 == axisB.GetDirection().X() && 0 == axisB.GetDirection().Y()) {
		return PointSegmentCollide(b.GetStartPoint(), a)
	}
	if OnOneSide(axisB, a) {
		return false
	}

	if axisA.GetDirection().IsParallel(axisB.GetDirection()) {
		d := axisA.GetDirection().UnitVector2D() 
		rangeA := a.ProjectSegment(d, true)
		rangeB := b.ProjectSegment(d, true)
		return math2D.OverlappingRanges(rangeA, rangeB)
	}

	return true
}

 

func LineRectangleCollide(l math2D.Line, r math2D.Rectangle) bool {
	n := l.GetDirection().RotateVector90()

	c1 := r.GetPosition()
	c2 := math2D.AddVectors(c1, r.GetSize())
	c3 := math2D.NewVector2D(c2.X(), c1.Y())
	c4 := math2D.NewVector2D(c1.X(), c2.Y())

	c1 = math2D.SubtractVectors(c1, l.GetBase())
	c2 = math2D.SubtractVectors(c2, l.GetBase())
	c3 = math2D.SubtractVectors(c3, l.GetBase())
	c4 = math2D.SubtractVectors(c4, l.GetBase())

	dp1 := math2D.DotProduct(n, c1)
	dp2 := math2D.DotProduct(n, c2)
	dp3 := math2D.DotProduct(n, c3)
	dp4 := math2D.DotProduct(n, c4)

	return (dp1 * dp2 <= 0) || (dp2 * dp3 <= 0) || (dp3 * dp4 <= 0)
}


 

func LineOrientedRectangleCollide( l math2D.Line, r math2D.OrientedRectangle) bool {
	lr := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(r.GetHalfExtended().X() * 2, r.GetHalfExtended().Y() * 2))

	base := math2D.SubtractVectors(l.GetBase(), r.GetCenter())
	base = base.RotateVector(-r.GetRotation())
	base = math2D.AddVectors(base, r.GetHalfExtended())
	direction := l.GetDirection().RotateVector(-r.GetRotation())
	line := math2D.NewLine(base, direction)
	
	return LineRectangleCollide(line, lr)
}

func OrientedRectanglesCollide(a, b math2D.OrientedRectangle) bool {
	edge := OrientedRectangleEdge(a, 0)

	if SeparatingAxisForOrientedRectangle(edge, b) {
		return false
	}

	edge = OrientedRectangleEdge(a, 1)
	if SeparatingAxisForOrientedRectangle(edge, b) {
		return false
	}

	edge =  OrientedRectangleEdge(b, 0)
	if SeparatingAxisForOrientedRectangle(edge, a) {
		return false
	}

	edge = OrientedRectangleEdge(b, 1) 
	return !SeparatingAxisForOrientedRectangle(edge, a)
}

func PointRectangleCollide(p math2D.Vector2D, r math2D.Rectangle) bool {
	left := r.GetPosition().X()
	right := left + r.GetSize().X()
	bottom := r.GetPosition().Y()
	top := bottom + r.GetSize().Y()
	return left <= p.X() && bottom <= p.Y() && p.X() <= right && p.Y() <= top
}


func OrientedRectanglePointCollide( r math2D.OrientedRectangle, p math2D.Vector2D) bool {
	lr := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D( r.GetHalfExtended().X() * 2, r.GetHalfExtended().Y() * 2))
	lp := math2D.SubtractVectors(p, r.GetCenter())
	lp = lp.RotateVector(-r.GetRotation())
	lp = math2D.AddVectors(lp, r.GetHalfExtended())
	return PointRectangleCollide(lp, lr)
}


func OrientedRectangleRectangleCollide(or math2D.OrientedRectangle, r math2D.Rectangle) bool {
	orHull := OrientedRectangleRectangleHull(or)
	if !RectanglesCollide(orHull, r){
		return false
	}

	edge := OrientedRectangleEdge(or, 0)
	if SeparatingAxisForRectangle(edge, r) {
		return false
	}

	edge = OrientedRectangleEdge(or, 1) 
	return !SeparatingAxisForRectangle(edge, r)
}

 

func RectangleSegmentCollide(r math2D.Rectangle, s math2D.Segment) bool {
	rRange := math2D.NewRange(r.GetPosition().X(), r.GetPosition().X() + r.GetSize().X())
	sRange := math2D.NewRange( s.GetStartPoint().X(), s.GetEndPoint().X())
	sRange = sRange.SortRange()
	if !math2D.OverlappingRanges(rRange, sRange) {
		return false
	}
	
	rRange.SetMinimun(r.GetPosition().Y())
	rRange.SetMaximum(r.GetPosition().Y() + r.GetSize().Y())
	sRange.SetMinimun(s.GetStartPoint().Y())
	sRange.SetMaximum( s.GetEndPoint().Y())
	sRange = sRange.SortRange()
	if !math2D.OverlappingRanges(rRange, sRange) {
		return false
	}

	sLine := math2D.NewLine( s.GetStartPoint(), math2D.SubtractVectors(s.GetEndPoint(), s.GetStartPoint()))
	return LineRectangleCollide(sLine, r)
}

func OrientedRectangleSegmentCollide(r math2D.OrientedRectangle, s math2D.Segment) bool{
	lr := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(r.GetHalfExtended().X() * 2, r.GetHalfExtended().Y() * 2))

	ls := math2D.Segment{}
	ls.SetStartPoint( math2D.SubtractVectors(s.GetStartPoint(), r.GetCenter()))
	ls.SetStartPoint( ls.GetStartPoint().RotateVector(-r.GetRotation()))
	ls.SetStartPoint(math2D.AddVectors(ls.GetStartPoint(), r.GetHalfExtended()))
	ls.SetEndPoint( math2D.SubtractVectors(s.GetEndPoint(), r.GetCenter()))
	ls.SetEndPoint( ls.GetEndPoint().RotateVector(-r.GetRotation()))
	ls.SetEndPoint( math2D.AddVectors( ls.GetEndPoint(), r.GetHalfExtended()))

	return RectangleSegmentCollide(lr, ls)
}