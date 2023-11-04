package collision

var collisionMaskSize = 10

type CollisionMaskBit uint64

type CollisionMask struct {
	bitmask Bitvector
}

func NewCollisionMask() *CollisionMask {
	return &CollisionMask{bitmask: make(Bitvector, collisionMaskSize)}
}

func (c *CollisionMask) SetBit(n int) {
	err := Bitset(c.bitmask, n)
	if err != nil {
		return
	}
}

func (c *CollisionMask) UnsetBit(n int) {
	err := Bitunset(c.bitmask, n)
	if err != nil {
		return
	}
}

func (c *CollisionMask) IsCollidible(n int) bool {
	result, _ := IsBitSet(c.bitmask, n)
	return result
}
