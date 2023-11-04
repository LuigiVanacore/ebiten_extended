package collision

import (

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionShape struct {
	transform     *transform.Transform
	shape         math2D.Shape
	collisionMask CollisionMask
	debug         bool
}

func NewCollisionShape(transform *transform.Transform, shape math2D.Shape, mask CollisionMask) *CollisionShape {
	collisionShape := &CollisionShape{transform: transform, collisionMask: *NewCollisionMask()}
	//shape.SetTransform(transform)
	collisionShape.shape = shape
	collisionShape.collisionMask = mask
	return collisionShape
}

func (c *CollisionShape) Update() {

}

func (c *CollisionShape) IsColliding(collider Collidable) bool {
	if c.IsCollidible(collider.GetTag()) {
		return c.shape.Intersect(collider.GetShape())
	}
	return false
}

func (c *CollisionShape) GetShape() math2D.Shape {
	return c.shape
}

func (c *CollisionShape) DrawDebug(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if c.debug {
		//vector.DrawFilledCircle(target, float32(c.Transform.position.X), float32(c.Transform.position.Y), 15, color.White, false)
		//m := c.transform.GetGeoM()
		/* if shape, ok := c.shape.(*math2D.Circle); ok {
			vector.StrokeCircle(target, float32(c.transform.GetPosition().X()), float32(c.transform.GetPosition().Y()), float32(shape.GetRadius()), 2, color.White, false)
		} */
	}
}

func (c *CollisionShape) SetDebug(isDebug bool) {
	c.debug = isDebug
}

func (c *CollisionShape) IsCollidible(tag int) bool {
	return c.collisionMask.IsCollidible(tag)
}
