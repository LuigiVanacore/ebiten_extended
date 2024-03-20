package collision

import "github.com/LuigiVanacore/ebiten_extended/utils"


// CollisionMask identifies the kind of the shape
// to collide other shapes.
type CollisionMask struct {
	// Identity determines which other
	// shapes can collide the present shape.
	identity utils.ByteSet
	// Mask determines which other shapes
	// the present shape can collide.
	mask utils.ByteSet
}


func NewCollisionMask(identity, mask utils.ByteSet) CollisionMask {
	return CollisionMask{ identity: identity, mask: mask }
}
// ShouldCollide returns true if the shape should
// collide another one accodring to their tags.
func (t CollisionMask) IsCollidible(other CollisionMask) bool {
	return t.mask.Has(other.GetIdentity())
}

// GetIdentity returns the valye of the shape identity.
func (t CollisionMask) GetIdentity() utils.ByteSet {
	return t.identity
}

// SetIdentity assigns a new value to the tag identity.
func (t *CollisionMask) SetIdentity(newIdentity utils.ByteSet) {
	t.identity = newIdentity
}

// GetMask the value of the shape mask.
func (t CollisionMask) GetMask() utils.ByteSet {
	return t.mask
}

// SetMask assigns a new value to the tag mask.
func (t *CollisionMask) SetMask(newMask utils.ByteSet) {
	t.mask = newMask
}
