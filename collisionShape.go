package ebiten_extended

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type CollisionShape struct {
	transform     *Transform
	shape         Shape
	collisionMask CollisionMask
	debug         bool
}

func NewCollisionShape(transform *Transform, shape Shape, mask CollisionMask) *CollisionShape {
	collisionShape := &CollisionShape{transform: transform, collisionMask: *NewCollisionMask()}
	shape.SetTransform(transform)
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

func (c *CollisionShape) GetShape() Shape {
	return c.shape
}

func (c *CollisionShape) DrawDebug(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	if c.debug {
		//vector.DrawFilledCircle(target, float32(c.Transform.position.X), float32(c.Transform.position.Y), 15, color.White, false)
		//m := c.transform.GetGeoM()
		if shape, ok := c.shape.(*Circle); ok {
			vector.StrokeCircle(target, float32(c.transform.position.X), float32(c.transform.position.Y), float32(shape.radius), 2, color.White, false)
		}
	}
}

func (c *CollisionShape) SetDebug(isDebug bool) {
	c.debug = isDebug
}

func (c *CollisionShape) IsCollidible(tag int) bool {
	return c.collisionMask.IsCollidible(tag)
}
