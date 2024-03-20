package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionCircle struct {
	circle math2D.Circle
	transform *transform.Transform
}

func NewCollisionCircle(circle math2D.Circle, transform *transform.Transform) *CollisionCircle {
	return &CollisionCircle{ circle: circle, transform: transform}
}

func (c *CollisionCircle) GetTransform() *transform.Transform {
	return c.transform
}

func (c *CollisionCircle) SetTransform(transform *transform.Transform) {
	c.transform = transform
}

func (c *CollisionCircle) Update() {
	c.updatePosition()
}

func (c *CollisionCircle) updatePosition() {
	c.circle.SetCenter(c.transform.GetPosition())
}


func (c *CollisionCircle) IsColliding( collisionShape CollisionShape) bool {
	switch other := collisionShape.(type) {
	case *CollisionCircle:
		return CirclesCollide( c.circle, other.circle)
	case *CollisionRect:
		return CircleRectangleCollide( c.circle, other.Rectangle)
	default:
		return false
	}
}

