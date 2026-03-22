package collision

import "github.com/LuigiVanacore/ludum/math2d"

// maxBisectionDepth limits the recursive bisection in moving-shape collision tests,
// preventing stack overflows when the motion vector is very large or never shrinks.
const maxBisectionDepth = 20

func MovingCircleCircleCollide(a, b math2d.Circle, moveA math2d.Vector2D) bool {
	bAbsorbedA := math2d.NewCircle(math2d.NewVector2D(b.GetCenterPosition().X(), b.GetCenterPosition().Y()), b.GetRadius()+a.GetRadius())
	travelA := math2d.NewSegment(a.GetCenterPosition(), math2d.AddVectors(a.GetCenterPosition(), moveA))
	return CircleSegmentCollide(bAbsorbedA, travelA)
}

// MovingRectangleRectangleCollide reports whether rectangle a, moving by moveA, would hit rectangle b.
func MovingRectangleRectangleCollide(a math2d.Rectangle, moveA math2d.Vector2D, b math2d.Rectangle) bool {
	return movingRectRectCollide(a, moveA, b, 0)
}

func movingRectRectCollide(a math2d.Rectangle, moveA math2d.Vector2D, b math2d.Rectangle, depth int) bool {
	if depth >= maxBisectionDepth {
		return RectanglesCollide(a, b)
	}

	envelope := a
	envelope.SetPosition(math2d.AddVectors(envelope.GetPosition(), moveA))
	envelope = EnlargeRectangleRectangle(envelope, a)

	if RectanglesCollide(envelope, b) {
		halfMoveA := moveA.DivideScalar(2)
		envelope.SetPosition(math2d.AddVectors(a.GetPosition(), halfMoveA))
		envelope.SetSize(a.GetSize())
		return movingRectRectCollide(a, halfMoveA, b, depth+1) || movingRectRectCollide(envelope, halfMoveA, b, depth+1)
	}

	return false
}

// MovingCircleRectangleCollide reports whether circle a, moving by moveA, would hit rectangle b.
func MovingCircleRectangleCollide(a math2d.Circle, moveA math2d.Vector2D, b math2d.Rectangle) bool {
	return movingCircleRectCollide(a, moveA, b, 0)
}

func movingCircleRectCollide(a math2d.Circle, moveA math2d.Vector2D, b math2d.Rectangle, depth int) bool {
	if depth >= maxBisectionDepth {
		return CircleRectangleCollide(a, b)
	}

	halfMoveA := moveA.DivideScalar(2)
	moveDistance := moveA.Length()

	envelope := a
	envelope.SetCenter(math2d.AddVectors(a.GetCenterPosition(), halfMoveA))
	envelope.SetRadius(a.GetRadius() + moveDistance/2)

	if CircleRectangleCollide(envelope, b) {
		envelope.SetRadius(a.GetRadius())
		return movingCircleRectCollide(a, halfMoveA, b, depth+1) || movingCircleRectCollide(envelope, halfMoveA, b, depth+1)
	}

	return false
}

func MovingRectangleCircleCollide(a math2d.Rectangle, moveA math2d.Vector2D, b math2d.Circle) bool {
	moveB := moveA.Negate()
	return MovingCircleRectangleCollide(b, moveB, a)
}
