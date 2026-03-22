package collision

import (
	"github.com/LuigiVanacore/ludum/math2d"
	"github.com/LuigiVanacore/ludum/transform"
)

type CollisionCircle struct {
	circle math2d.Circle
}

func NewCollisionCircle(circle math2d.Circle) *CollisionCircle {
	return &CollisionCircle{circle: circle}
}

func (c *CollisionCircle) shapeKind() shapeKind {
	return kindCircle
}

func (c *CollisionCircle) IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool {
	return ShapeCollides(c, tSelf, other, tOther)
}
