package collision

import "math"

type collisionPair struct {
	c1 *Collider
	c2 *Collider
}

// CollisionManager holds colliders and runs broad-phase (spatial hash) and narrow-phase
// detection each frame. Set CellSize (e.g. 100) for the grid cell size in pixels.
// Call CheckCollision every frame (e.g. in World.SetPostUpdate).
type CollisionManager struct {
	colliders []*Collider

	// CellSize is the spatial hash grid cell size in pixels (e.g. 100).
	CellSize int

	// Tracks collisions that happened in the previous frame to deduce Enter/Stay/Exit
	// The uint64 key is an immutable hash of: min(id1, id2)<<32 | max(id1, id2).
	previousCollisions map[uint64]collisionPair
}

// NewCollisionManager returns a new collision manager with default CellSize 100.
func NewCollisionManager() *CollisionManager {
	return &CollisionManager{
		colliders:          make([]*Collider, 0),
		CellSize:           100, // Hardcoded default, can be exposed later
		previousCollisions: make(map[uint64]collisionPair),
	}
}

// AddCollider registers a collider so it is included in CheckCollision.
func (c *CollisionManager) AddCollider(collider *Collider) {
	c.colliders = append(c.colliders, collider)
}

func combineIDs(id1, id2 uint64) uint64 {
	if id1 < id2 {
		return (id1 << 32) | id2
	}
	return (id2 << 32) | id1
}

// CheckCollision evaluates intersections using a Spatial Hash Grid Broadphase
// and calculates lifecycle events (Enter, Stay, Exit) based on the frame memory.
func (c *CollisionManager) CheckCollision() {
	// 1. BROADPHASE: Populate Spatial Hash Grid
	grid := make(map[uint64][]*Collider)

	for _, coll := range c.colliders {
		pos := coll.GetWorldPosition()
		// Calculate the core cell using simple truncation
		cellX := int(math.Floor(pos.X() / float64(c.CellSize)))
		cellY := int(math.Floor(pos.Y() / float64(c.CellSize)))

		// To be safe with boundaries, we insert the collider in a 3x3 neighbor grid
		// (A strict AABB cell insertion would be better, but this fits dynamic shapes).
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				// We map X and Y into a single uint64 key
				x := uint64(uint32(cellX + dx))
				y := uint64(uint32(cellY + dy))
				cellKey := (x << 32) | y
				grid[cellKey] = append(grid[cellKey], coll)
			}
		}
	}

	// 2. NARROWPHASE && EVENT DISPATCHING
	currentCollisions := make(map[uint64]collisionPair)
	// We use a set string to deduplicate pairs checked multiple times in overlapping cells
	checkedPairs := make(map[uint64]bool)

	for _, cellColliders := range grid {
		for i := 0; i < len(cellColliders); i++ {
			for j := i + 1; j < len(cellColliders); j++ {
				col1 := cellColliders[i]
				col2 := cellColliders[j]

				if col1 == col2 {
					continue
				}

				pairID := combineIDs(col1.GetID(), col2.GetID())

				// Skip if already processed in another neighboring cell
				if checkedPairs[pairID] {
					continue
				}
				checkedPairs[pairID] = true

				if col1.CanCollideWith(col2) && col1.IsColliding(col2) {
					currentCollisions[pairID] = collisionPair{col1, col2}

					// Emit Lifecycle Events
					_, wasColliding := c.previousCollisions[pairID]
					if !wasColliding {
						// Enter
						if !col1.OnCollisionEnter.IsEmpty() {
							col1.OnCollisionEnter.Emit(col2)
						}
						if !col2.OnCollisionEnter.IsEmpty() {
							col2.OnCollisionEnter.Emit(col1)
						}
					} else {
						// Stay
						if !col1.OnCollisionStay.IsEmpty() {
							col1.OnCollisionStay.Emit(col2)
						}
						if !col2.OnCollisionStay.IsEmpty() {
							col2.OnCollisionStay.Emit(col1)
						}
					}
				}
			}
		}
	}

	// 3. EXIT EVENT RESOLUTION
	// If a pair existed yesterday but doesn't exist today, emit Exit!
	for pairID, pair := range c.previousCollisions {
		if _, stillColliding := currentCollisions[pairID]; !stillColliding {
			if !pair.c1.OnCollisionExit.IsEmpty() {
				pair.c1.OnCollisionExit.Emit(pair.c2)
			}
			if !pair.c2.OnCollisionExit.IsEmpty() {
				pair.c2.OnCollisionExit.Emit(pair.c1)
			}
		}
	}

	// Flip frame buffers
	c.previousCollisions = currentCollisions

	// Reset physics flag
	for _, coll := range c.colliders {
		coll.SetWorldCoordinateUpdated(false)
	}
}
