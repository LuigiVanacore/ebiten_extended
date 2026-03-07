package collision

import "github.com/LuigiVanacore/ebiten_extended/utils"

// Preset layer identities for common game object types.
// Use with NewPresetMask to create collision masks.
const (
	LayerPlayer    utils.ByteSet = 1 << iota
	LayerEnemy
	LayerWorld
	LayerPickup
	LayerProjectile
)

// NewPresetMask creates a CollisionMask from identity and a variadic list of layers to collide with.
// Example: NewPresetMask(LayerPlayer, LayerWorld, LayerEnemy, LayerPickup)
// means a player that collides with world, enemies, and pickups.
func NewPresetMask(identity utils.ByteSet, collidesWith ...utils.ByteSet) CollisionMask {
	var mask utils.ByteSet
	for _, layer := range collidesWith {
		mask = mask.Set(layer)
	}
	return NewCollisionMask(identity, mask)
}

// Common preset masks for typical 2D games.
var (
	// MaskPlayer collides with world and pickups.
	MaskPlayer = NewPresetMask(LayerPlayer, LayerWorld, LayerPickup)
	// MaskEnemy collides with world and player.
	MaskEnemy = NewPresetMask(LayerEnemy, LayerWorld, LayerPlayer)
	// MaskWorld collides with player, enemy, and projectile.
	MaskWorld = NewPresetMask(LayerWorld, LayerPlayer, LayerEnemy, LayerProjectile)
	// MaskPickup collides only with player.
	MaskPickup = NewPresetMask(LayerPickup, LayerPlayer)
	// MaskProjectile collides with world and enemy.
	MaskProjectile = NewPresetMask(LayerProjectile, LayerWorld, LayerEnemy)
)
