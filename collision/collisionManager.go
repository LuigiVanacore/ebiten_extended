package collision


type CollisionManager struct {
	colliders []*Collider
}


func NewCollisionManager() *CollisionManager {
	return &CollisionManager{ colliders: make([]*Collider, 0)}
}

func (c *CollisionManager) AddCollider(collider *Collider) {
	c.colliders = append(c.colliders, collider)
}

func (c *CollisionManager) CheckCollision() {
	for i := 0; i < len(c.colliders); i++ {
		for j := i + 1; j < len(c.colliders); j++ {
			collider1 := c.colliders[i]
			collider2 := c.colliders[j]
			if collider1.CanCollideWith(collider2) && collider1.IsColliding(collider2) {
				if collider1.onCollision != nil {
					collider1.onCollision(collider2)
				}
				if collider2.onCollision != nil {
					collider2.onCollision(collider1)
				}
			}
		}
	}
	for _, coll := range c.colliders {
		coll.SetWorldCoordinateUpdated(false)
	}
}