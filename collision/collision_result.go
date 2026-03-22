package collision

import (
	"math"

	"github.com/LuigiVanacore/ludum/math2d"
)

const epsilon = 1e-9

// CollisionResult holds the minimum translation vector (MTV) for resolving overlap.
// Normal points from B toward A (direction to push A out of B).
// Depth is the penetration amount to correct.
type CollisionResult struct {
	Overlapping bool
	Normal      math2d.Vector2D
	Depth       float64
}

// CirclesCollideResult returns MTV for two circles.
// Normal points from B toward A (direction to push A out of B).
// Edge case: coincident circles use Normal (1,0) as fallback.
func CirclesCollideResult(a, b math2d.Circle) CollisionResult {
	centerA := a.GetCenterPosition()
	centerB := b.GetCenterPosition()
	radiusSum := a.GetRadius() + b.GetRadius()
	delta := math2d.SubtractVectors(centerA, centerB) // A - B, points from B to A
	distSq := math2d.DotProduct(delta, delta)

	if distSq >= radiusSum*radiusSum {
		return CollisionResult{Overlapping: false}
	}

	dist := math.Sqrt(distSq)
	var normal math2d.Vector2D
	if dist < epsilon {
		normal = math2d.NewVector2D(1, 0)
	} else {
		normal = delta.DivideScalar(dist)
	}
	depth := radiusSum - dist
	return CollisionResult{
		Overlapping: true,
		Normal:      normal,
		Depth:       depth,
	}
}

// CircleRectangleCollideResult returns MTV for circle vs axis-aligned rectangle.
func CircleRectangleCollideResult(c math2d.Circle, r math2d.Rectangle) CollisionResult {
	center := c.GetCenterPosition()
	clamped := ClampOnRectangle(center, r)
	delta := math2d.SubtractVectors(center, clamped)
	distSq := math2d.DotProduct(delta, delta)
	radius := c.GetRadius()

	if distSq >= radius*radius {
		return CollisionResult{Overlapping: false}
	}

	dist := math.Sqrt(distSq)
	var normal math2d.Vector2D
	if dist < epsilon {
		normal = math2d.NewVector2D(1, 0)
	} else {
		normal = delta.DivideScalar(dist)
	}
	depth := radius - dist
	return CollisionResult{
		Overlapping: true,
		Normal:      normal,
		Depth:       depth,
	}
}

// RectanglesCollideResult returns MTV for two axis-aligned rectangles.
func RectanglesCollideResult(a, b math2d.Rectangle) CollisionResult {
	aLeft := a.GetPosition().X()
	aRight := aLeft + a.GetSize().X()
	aTop := a.GetPosition().Y()
	aBottom := aTop + a.GetSize().Y()

	bLeft := b.GetPosition().X()
	bRight := bLeft + b.GetSize().X()
	bTop := b.GetPosition().Y()
	bBottom := bTop + b.GetSize().Y()

	overlapX := math.Min(aRight, bRight) - math.Max(aLeft, bLeft)
	overlapY := math.Min(aBottom, bBottom) - math.Max(aTop, bTop)

	if overlapX <= 0 || overlapY <= 0 {
		return CollisionResult{Overlapping: false}
	}

	centerA := a.GetCenter()
	centerB := b.GetCenter()
	dx := centerA.X() - centerB.X()
	dy := centerA.Y() - centerB.Y()

	var normal math2d.Vector2D
	var depth float64
	if overlapX < overlapY {
		depth = overlapX
		if dx > 0 {
			normal = math2d.NewVector2D(1, 0)
		} else {
			normal = math2d.NewVector2D(-1, 0)
		}
	} else {
		depth = overlapY
		if dy > 0 {
			normal = math2d.NewVector2D(0, 1)
		} else {
			normal = math2d.NewVector2D(0, -1)
		}
	}

	return CollisionResult{
		Overlapping: true,
		Normal:      normal,
		Depth:       depth,
	}
}
