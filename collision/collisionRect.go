package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionRect struct {
	rectangle math2D.Rectangle
}



func (c *CollisionRect) UpdateTransform(transform transform.Transform) {
	c.rectangle.SetPosition(transform.GetPosition())
}

func (c *CollisionRect) IsColliding( collisionShape CollisionShape) bool {
	switch other := collisionShape.(type) {
	
	case *CollisionCircle:
		return CircleRectangleCollide(other.circle, c.rectangle )
	default:
		return false
	}
}
