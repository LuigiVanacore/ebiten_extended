package collision

import "github.com/LuigiVanacore/ebiten_extended/collision"

type collisionManager struct {
	colliders []*collision.Collider
}

var collisionManager_instance *collisionManager

func CollisionManager() *collisionManager {
	if collisionManager_instance == nil {
		collisionManager_instance = newCollisionManager()
	}

	return collisionManager_instance
}


func newCollisionManager() *collisionManager {
	return &collisionManager{ colliders: make([]*collision.Collider, 0)}
}

func (c *collisionManager) AddCollider(collider *collision.Collider) {
	c.colliders = append(c.colliders, collider)
}

func (c *collisionManager) CheckCollision() {
	
}