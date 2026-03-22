package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
)

type CollisionCircle struct {
	circle math2D.Circle
}

func NewCollisionCircle(circle math2D.Circle) *CollisionCircle {
	return &CollisionCircle{circle: circle}
}

func (c *CollisionCircle) shapeKind() shapeKind {
	return kindCircle
}

func (c *CollisionCircle) IsColliding(tSelf transform.Transform, other CollisionShape, tOther transform.Transform) bool {
	return ShapeCollides(c, tSelf, other, tOther)
}
