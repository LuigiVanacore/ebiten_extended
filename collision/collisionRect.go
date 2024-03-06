package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionRect struct {
	math2D.Rectangle
	transform *transform.Transform
}

func (c *CollisionRect) SetTransform(transform *transform.Transform) {
	c.transform = transform
}

func (c *CollisionRect) Update() {
	c.updatePosition() 
}

func (c *CollisionRect) updatePosition() {
	
}

func (c *CollisionRect) IsColliding( collisionShape CollisionShape) bool {
	switch other := collisionShape.(type) {
	
	case *CollisionCircle:
		return CircleRectangleCollide(other.circle, c.Rectangle )
	default:
		return false
	}
}
