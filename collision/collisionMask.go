package collision

var collisionMaskSize = 64

type CollisionMask uint64



// func NewCollisionMask() *CollisionMask {
// 	return &CollisionMask{bitmask: make(Bitvector, collisionMaskSize)}
// }

// func (c *CollisionMask) SetBit(n int) {
// 	err := Bitset(c.bitmask, n)
// 	if err != nil {
// 		return
// 	}
// }

// func (c *CollisionMask) UnsetBit(n int) {
// 	err := Bitunset(c.bitmask, n)
// 	if err != nil {
// 		return
// 	}
// }

// func (c *CollisionMask) IsCollidible(n int) bool {
// 	result, _ := IsBitSet(c.bitmask, n)
// 	return result
// }





func (c CollisionMask) Set(flag CollisionMask) CollisionMask {
	return c | flag
}

func (c CollisionMask) Clear(flag CollisionMask) CollisionMask {
	return c &^ flag
}

func (c CollisionMask) IsCollidible(flag CollisionMask) bool {
	return c&flag != 0
}
