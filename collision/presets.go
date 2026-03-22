package collision

import "github.com/LuigiVanacore/ludum/utils"

// Preset layer identities (power-of-2 bits) for common game object types.
// Use with NewPresetMask or combine with utils.ByteSet for custom masks.
const (
	LayerPlayer utils.ByteSet = 1 << iota
	LayerEnemy
	LayerWorld
	LayerPickup
	LayerProjectile
)

// NewPresetMask creates a CollisionMask: identity is this object's layer; collidesWith are the
// layers it will collide with. Example: NewPresetMask(LayerPlayer, LayerWorld, LayerEnemy, LayerPickup)
// creates a player that collides with world, enemies, and pickups.
func NewPresetMask(identity utils.ByteSet, collidesWith ...utils.ByteSet) CollisionMask {
	var mask utils.ByteSet
	for _, layer := range collidesWith {
		mask = mask.Set(layer)
	}
	return NewCollisionMask(identity, mask)
}

// Common preset masks for typical 2D games. Use directly or as reference for custom masks.
//
// Collision matrix (✓ = pair collides):
//
//	          Player Enemy World Pickup Projectile
//	Player      -     ✓     ✓     ✓        -
//	Enemy       ✓     -     ✓     -        ✓
//	World       ✓     ✓     -     -        ✓
//	Pickup      ✓     -     -     -        -
//	Projectile  -     ✓     ✓     -        -
var (
	MaskPlayer     = NewPresetMask(LayerPlayer, LayerWorld, LayerPickup)
	MaskEnemy      = NewPresetMask(LayerEnemy, LayerWorld, LayerPlayer)
	MaskWorld      = NewPresetMask(LayerWorld, LayerPlayer, LayerEnemy, LayerProjectile)
	MaskPickup     = NewPresetMask(LayerPickup, LayerPlayer)
	MaskProjectile = NewPresetMask(LayerProjectile, LayerWorld, LayerEnemy)
)
