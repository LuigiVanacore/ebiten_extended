package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

 

func EnlargeRectanglePoint( r math2D.Rectangle, p math2D.Vector2D) math2D.Rectangle {
	enlarged := math2D.NewRectangle( math2D.NewVector2D(math2D.Min(r.GetPosition().X(),p.X()),
									 					math2D.Min(r.GetPosition().Y(), p.Y())),
									 math2D.NewVector2D(math2D.Max(r.GetPosition().X() + r.GetSize().X(), p.X()),
														math2D.Max(r.GetPosition().Y() + r.GetSize().Y(), p.Y())))
	size := math2D.SubtractVectors(enlarged.GetSize(), enlarged.GetPosition())
	enlarged.SetSize(size)
	return enlarged
}

 

func EnlargeRectangleRectangle(r, extender math2D.Rectangle) math2D.Rectangle {
	maxCorner := math2D.AddVectors(extender.GetPosition(), extender.GetSize())
	enlarged := EnlargeRectanglePoint(r, maxCorner)
	return EnlargeRectanglePoint(enlarged, extender.GetPosition())
}



func OrientedRectangleRectangleHull(r math2D.OrientedRectangle) math2D.Rectangle {
	h := math2D.NewRectangle(r.GetCenter(), math2D.ZeroVector2D())

	for nr := 1; nr < 4; nr++ {
		corner := OrientedRectangleCorner(r, nr)
		h = EnlargeRectanglePoint(h, corner)
	}
	return h
}


func RectanglesRectangleHull(rectangles []math2D.Rectangle, count int ) math2D.Rectangle {
	h := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.ZeroVector2D())
	if ( 0 == count || len(rectangles) == 0 ) {
		return h
	}

	h = rectangles[0]
	for i := 1; i < count; i++ {
		h = EnlargeRectangleRectangle(h , rectangles[i])
	}
	return h
}
  

func OrientedRectangleCircleHull(r math2D.OrientedRectangle) math2D.Circle {
	return math2D.NewCircle(r.GetCenter(), r.GetHalfExtended().Length())
}
 
func CirlcesRectangleHull(circles []math2D.Circle, count int) math2D.Rectangle {
	h := math2D.NewRectangle(math2D.ZeroVector2D(), math2D.ZeroVector2D())
	if (0 == count || len(circles) == 0 ) {
		return h
	}

	h.SetPosition(circles[0].GetCenterPosition())
	for i:=1; i < count; i++ {
		halfExtend := math2D.NewVector2D(circles[i].GetRadius(), circles[i].GetRadius())
		minP := math2D.SubtractVectors(circles[i].GetCenterPosition(), halfExtend)
		maxP := math2D.AddVectors(circles[i].GetCenterPosition(), halfExtend)
		h = EnlargeRectanglePoint(h, minP)
		h = EnlargeRectanglePoint(h, maxP)
	}
	return h
}
 

func CirclesCircleHull(circles []math2D.Circle, count int) math2D.Circle {
	h := math2D.NewCircle(math2D.NewVector2D(0, 0), 0)
	rh := CirlcesRectangleHull(circles, count)
	h.SetCenter(rh.GetCenter())

	for i:=1; i < count; i++ {
		d := math2D.SubtractVectors(circles[i].GetCenterPosition(), h.GetCenterPosition())
		h.SetRadius(math2D.Max(d.Length() + circles[i].GetRadius(), h.GetRadius()))
	}
	return h
}