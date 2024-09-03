package collision

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
)

type Collider struct {
	ebiten_extended.Node2D
	collisionShape          CollisionShape
	mask                    CollisionMask
	isWorldCordinateUpdated bool
}

func NewCollider(shape CollisionShape, mask CollisionMask) *Collider {
	c := &Collider{mask: mask}
	CollisionManager().AddCollider(c)
	return c
}

func (c *Collider) IsWorldCordinateUpdated() bool {
	return c.isWorldCordinateUpdated
}

func (c *Collider) SetWorldCordinateUpdated(flag bool) {
	c.isWorldCordinateUpdated = flag
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

func (c *Collider) CanCollideWith(collider *Collider) bool {
	return c.mask.IsCollidible(collider.GetCollisionMask())
}

func (c *Collider) IsColliding(collider *Collider) bool {
	leftWorldTransf := c.GetWorldTransform()
	rightWorldTransf := collider.GetWorldTransform()
	c.collisionShape.UpdateTransform(leftWorldTransf)
	collider.collisionShape.UpdateTransform(rightWorldTransf)
	return c.collisionShape.IsColliding(collider.collisionShape)
}
