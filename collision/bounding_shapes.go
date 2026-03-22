package collision

import (
	"github.com/LuigiVanacore/ludum/math2d"
)

func EnlargeRectanglePoint(r math2d.Rectangle, p math2d.Vector2D) math2d.Rectangle {
	enlarged := math2d.NewRectangle(math2d.NewVector2D(math2d.Min(r.GetPosition().X(), p.X()),
		math2d.Min(r.GetPosition().Y(), p.Y())),
		math2d.NewVector2D(math2d.Max(r.GetPosition().X()+r.GetSize().X(), p.X()),
			math2d.Max(r.GetPosition().Y()+r.GetSize().Y(), p.Y())))
	size := math2d.SubtractVectors(enlarged.GetSize(), enlarged.GetPosition())
	enlarged.SetSize(size)
	return enlarged
}

func EnlargeRectangleRectangle(r, extender math2d.Rectangle) math2d.Rectangle {
	maxCorner := math2d.AddVectors(extender.GetPosition(), extender.GetSize())
	enlarged := EnlargeRectanglePoint(r, maxCorner)
	return EnlargeRectanglePoint(enlarged, extender.GetPosition())
}

func OrientedRectangleRectangleHull(r math2d.OrientedRectangle) math2d.Rectangle {
	h := math2d.NewRectangle(r.GetCenter(), math2d.ZeroVector2D())

	for nr := 1; nr < 4; nr++ {
		corner := OrientedRectangleCorner(r, nr)
		h = EnlargeRectanglePoint(h, corner)
	}
	return h
}

func RectanglesRectangleHull(rectangles []math2d.Rectangle, count int) math2d.Rectangle {
	h := math2d.NewRectangle(math2d.ZeroVector2D(), math2d.ZeroVector2D())
	if 0 == count || len(rectangles) == 0 {
		return h
	}

	h = rectangles[0]
	for i := 1; i < count; i++ {
		h = EnlargeRectangleRectangle(h, rectangles[i])
	}
	return h
}

func OrientedRectangleCircleHull(r math2d.OrientedRectangle) math2d.Circle {
	return math2d.NewCircle(r.GetCenter(), r.GetHalfExtended().Length())
}

// CirclesRectangleHull returns the AABB that contains all given circles.
func CirclesRectangleHull(circles []math2d.Circle, count int) math2d.Rectangle {
	h := math2d.NewRectangle(math2d.ZeroVector2D(), math2d.ZeroVector2D())
	if 0 == count || len(circles) == 0 {
		return h
	}

	h.SetPosition(circles[0].GetCenterPosition())
	for i := 1; i < count; i++ {
		halfExtend := math2d.NewVector2D(circles[i].GetRadius(), circles[i].GetRadius())
		minP := math2d.SubtractVectors(circles[i].GetCenterPosition(), halfExtend)
		maxP := math2d.AddVectors(circles[i].GetCenterPosition(), halfExtend)
		h = EnlargeRectanglePoint(h, minP)
		h = EnlargeRectanglePoint(h, maxP)
	}
	return h
}

func CirclesCircleHull(circles []math2d.Circle, count int) math2d.Circle {
	h := math2d.NewCircle(math2d.NewVector2D(0, 0), 0)
	rh := CirclesRectangleHull(circles, count)
	h.SetCenter(rh.GetCenter())

	for i := 1; i < count; i++ {
		d := math2d.SubtractVectors(circles[i].GetCenterPosition(), h.GetCenterPosition())
		h.SetRadius(math2d.Max(d.Length()+circles[i].GetRadius(), h.GetRadius()))
	}
	return h
}
