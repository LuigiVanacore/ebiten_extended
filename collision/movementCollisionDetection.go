package collision

import "github.com/LuigiVanacore/ebiten_extended/math2D"

 
func MovingCircleCircleCollide( a, b math2D.Circle, moveA math2D.Vector2D) bool {
	bAbsorbedA := math2D.NewCircle( math2D.NewVector2D( b.GetCenterPosition().X(), b.GetCenterPosition().Y()), b.GetRadius() + a.GetRadius())
	travelA := math2D.NewSegment(a.GetCenterPosition(), math2D.AddVectors(a.GetCenterPosition(), moveA))
	return CircleSegmentCollide(bAbsorbedA, travelA)
} 


func MovingRectangleRectangleCollide( a math2D.Rectangle, moveA math2D.Vector2D, b math2D.Rectangle) bool {
	envelope := a
	envelope.SetPosition(math2D.AddVectors(envelope.GetPosition(), moveA))
	envelope = EnlargeRectangleRectangle(envelope, a)

	if RectanglesCollide(envelope, b) {
		halfMoveA := moveA.DivideScalar(2)

		envelope.SetPosition(math2D.AddVectors(a.GetPosition(), halfMoveA))
		envelope.SetSize(a.GetSize())
		return MovingRectangleRectangleCollide(a, halfMoveA, b) || MovingRectangleRectangleCollide(envelope, halfMoveA, b)
	}

	return false
}


func MovingCircleRectangleCollide( a math2D.Circle, moveA math2D.Vector2D, b math2D.Rectangle) bool {
	envelope := a
	halfMoveA := moveA.DivideScalar(2)
	moveDistance := moveA.Length()
	envelope.SetCenter(math2D.AddVectors(a.GetCenterPosition(), halfMoveA))
	envelope.SetRadius(a.GetRadius() + moveDistance / 2)

	if CircleRectangleCollide(envelope, b) {
		envelope.SetRadius(a.GetRadius())
		return MovingCircleRectangleCollide(a, halfMoveA, b) || MovingCircleRectangleCollide(envelope, halfMoveA, b)	
	}

	return false
}

func MovingRectangleCircleCollide( a math2D.Rectangle, moveA math2D.Vector2D, b math2D.Circle) bool {
	moveB := moveA.Negate()
	return MovingCircleRectangleCollide(b, moveB, a)
}

