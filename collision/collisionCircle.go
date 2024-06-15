package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionCircle struct {
	circle math2D.Circle
}

func NewCollisionCircle(circle math2D.Circle) *CollisionCircle {
	return &CollisionCircle{ circle: circle}
}


func (c *CollisionCircle) UpdateTransform(transform transform.Transform) {
	c.circle.SetCenter(transform.GetPosition())
}


func (c *CollisionCircle) IsColliding( collisionShape CollisionShape) bool {
	switch other := collisionShape.(type) {
	case *CollisionCircle:
		return CirclesCollide( c.circle, other.circle)
	case *CollisionRect:
		return CircleRectangleCollide( c.circle, other.rectangle)
	default:
		return false
	}
}

