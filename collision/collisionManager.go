package collision


type collisionManager struct {
	colliders []*Collider
}

var collisionManager_instance *collisionManager

func CollisionManager() *collisionManager {
	if collisionManager_instance == nil {
		collisionManager_instance = newCollisionManager()
	}

	return collisionManager_instance
}


func newCollisionManager() *collisionManager {
	return &collisionManager{ colliders: make([]*Collider, 0)}
}

func (c *collisionManager) AddCollider(collider *Collider) {
	c.colliders = append(c.colliders, collider)
}

func (c *collisionManager) CheckCollision() {
	for i := 0; i < len(c.colliders); i++ {
		for j := i + 1; j < len(c.colliders); j++ {
			collider1 := c.colliders[i]
			collider2 := c.colliders[j]
		//	if collider1.IsAlive() && collider2.IsAlive() {
				if collider1.CanCollideWith(collider2) {
					if collider1.IsColliding(collider2) {

					}
				//	collider1.OnDestroy()
				//	fmt.Printf("first entity %s is colliding with second entity %s", collider1.ToString(), collider2.ToString())
				}
				if collider2.CanCollideWith(collider1) {
				//	collider2.OnDestroy()
				//	fmt.Printf("second entity %s is colliding with first entity %s \n", collider2.ToString(), collider1.ToString())
				}
		//	}
		}
	}
	for _, coll := range c.colliders {
		coll.SetWorldCordinateUpdated(false)
	}
}