package collision

import (
	"github.com/LuigiVanacore/ebiten_extended/transform"
	"github.com/hajimehoshi/ebiten/v2"
)

type Collider struct {
	transform     transform.Transform
	collisionShape  CollisionShape
	mask CollisionMask
}



func NewCollider(transform transform.Transform, shape CollisionShape, mask CollisionMask) *Collider {
	c := Collider{transform: transform,  mask: mask }
	shape.SetTransform(&c.transform)
	return &c
}

func (c *Collider) Update() {
	c.collisionShape.Update()
}

func (c *Collider) IsColliding(collider Collider) bool {
	if c.IsCollidible(collider.GetCollisionMask()) {
		return c.collisionShape.IsColliding(collider.GetShape())
	}
	return false
}

func (c *Collider) GetCollisionMask() CollisionMask {
	return c.mask
}

func (c *Collider) SetCollisionMask(mask CollisionMask) {
	c.mask = mask
}

func (c *Collider) GetShape() CollisionShape {
	return c.collisionShape
}

func (c *Collider) DrawDebug(target *ebiten.Image, op *ebiten.DrawImageOptions) {
	//if c.debug {
		//vector.DrawFilledCircle(target, float32(c.Transform.position.X), float32(c.Transform.position.Y), 15, color.White, false)
		//m := c.transform.GetGeoM()
		/* if shape, ok := c.shape.(*math2D.Circle); ok {
			vector.StrokeCircle(target, float32(c.transform.GetPosition().X()), float32(c.transform.GetPosition().Y()), float32(shape.GetRadius()), 2, color.White, false)
		} */
	//}
}



func (c *Collider) IsCollidible(mask CollisionMask) bool {
	return c.mask.IsCollidible(mask)
}
