package collision

import "github.com/LuigiVanacore/ludum/utils"

// CollisionMask controls which collision pairs are checked. It uses two bit sets (utils.ByteSet):
//
//   - Identity: the layer(s) this object belongs to (what it "is"). Other objects with this layer
//     in their mask can collide with it.
//   - Mask: the layer(s) this object responds to. This object will collide with others whose
//     identity is in this mask.
//
// For a pair (A, B) to collide, both must allow it: A.mask must include B.identity AND
// B.mask must include A.identity. Use power-of-2 bits (1<<0, 1<<1, …) for layers to avoid overlap.
type CollisionMask struct {
	identity utils.ByteSet
	mask     utils.ByteSet
}

// NewCollisionMask creates a mask with the given identity (layer membership) and mask (layers to collide with).
// Example: NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1|2)) = object on layer 1 that collides with layers 1 and 2.
func NewCollisionMask(identity, mask utils.ByteSet) CollisionMask {
	return CollisionMask{identity: identity, mask: mask}
}

// IsCollidable returns true if this shape's mask includes the other's identity.
// Collision between A and B requires both A.IsCollidable(B) and B.IsCollidable(A).
func (t CollisionMask) IsCollidable(other CollisionMask) bool {
	return t.mask.Has(other.GetIdentity())
}

// GetIdentity returns the layer(s) this object belongs to.
func (t CollisionMask) GetIdentity() utils.ByteSet {
	return t.identity
}

// SetIdentity sets the layer(s) this object belongs to.
func (t *CollisionMask) SetIdentity(newIdentity utils.ByteSet) {
	t.identity = newIdentity
}

// GetMask returns the layer(s) this object can collide with.
func (t CollisionMask) GetMask() utils.ByteSet {
	return t.mask
}

// SetMask sets the layer(s) this object can collide with.
func (t *CollisionMask) SetMask(newMask utils.ByteSet) {
	t.mask = newMask
}
