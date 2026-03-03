package collision

import "math"

type collisionPair struct {
	a CollisionParticipant
	b CollisionParticipant
}

// CollisionManager holds collision participants and runs broad-phase (AABB-based spatial hash) and narrow-phase
// detection each frame. Use AddCollider for Colliders, AddParticipant for Area2D/RigidBody2D.
// Call CheckCollision every frame (e.g. in World.SetPostUpdate).
type CollisionManager struct {
	participants []CollisionParticipant

	// CellSize is the spatial hash grid cell size in pixels (e.g. 100).
	CellSize int

	previousCollisions map[uint64]collisionPair
}

// NewCollisionManager returns a new collision manager with default CellSize 100.
func NewCollisionManager() *CollisionManager {
	return &CollisionManager{
		participants:       make([]CollisionParticipant, 0),
		CellSize:           100,
		previousCollisions: make(map[uint64]collisionPair),
	}
}

// AddCollider registers a collider. Kept for backward compatibility.
func (m *CollisionManager) AddCollider(collider *Collider) {
	m.AddParticipant(collider)
}

// AddParticipant registers a collision participant (Collider, Area2D, or RigidBody2D).
func (m *CollisionManager) AddParticipant(p CollisionParticipant) {
	if p == nil {
		return
	}
	m.participants = append(m.participants, p)
}

func combineIDs(id1, id2 uint64) uint64 {
	if id1 < id2 {
		return (id1 << 32) | id2
	}
	return (id2 << 32) | id1
}

func isOverlapping(a, b CollisionParticipant) bool {
	sa, sb := a.GetShape(), b.GetShape()
	if sa == nil || sb == nil {
		return false
	}
	sa.UpdateTransform(a.GetWorldTransform())
	sb.UpdateTransform(b.GetWorldTransform())
	return ShapeCollides(sa, sb)
}

// CheckCollision evaluates intersections and emits lifecycle events.
func (m *CollisionManager) CheckCollision() {
	grid := make(map[uint64][]CollisionParticipant)
	cellSize := float64(m.CellSize)

	for _, p := range m.participants {
		shape := p.GetShape()
		if shape == nil {
			continue
		}
		shape.UpdateTransform(p.GetWorldTransform())
		minX, minY, maxX, maxY := ShapeAABB(shape)
		cellMinX := int(math.Floor(minX / cellSize))
		cellMaxX := int(math.Floor(maxX / cellSize))
		cellMinY := int(math.Floor(minY / cellSize))
		cellMaxY := int(math.Floor(maxY / cellSize))
		for cx := cellMinX; cx <= cellMaxX; cx++ {
			for cy := cellMinY; cy <= cellMaxY; cy++ {
				key := (uint64(uint32(cx)) << 32) | uint64(uint32(cy))
				grid[key] = append(grid[key], p)
			}
		}
	}

	currentCollisions := make(map[uint64]collisionPair)
	checkedPairs := make(map[uint64]bool)

	for _, cellParticipants := range grid {
		for i := 0; i < len(cellParticipants); i++ {
			for j := i + 1; j < len(cellParticipants); j++ {
				a := cellParticipants[i]
				b := cellParticipants[j]

				if a == b {
					continue
				}

				pairID := combineIDs(a.GetID(), b.GetID())
				if checkedPairs[pairID] {
					continue
				}
				checkedPairs[pairID] = true

				if !a.CanCollideWith(b) || !b.CanCollideWith(a) {
					continue
				}
				if !isOverlapping(a, b) {
					continue
				}

				currentCollisions[pairID] = collisionPair{a: a, b: b}
				m.emitCollisionEvents(a, b, pairID)
			}
		}
	}

	for pairID, pair := range m.previousCollisions {
		if _, still := currentCollisions[pairID]; !still {
			m.emitExitEvents(pair.a, pair.b)
		}
	}

	m.previousCollisions = currentCollisions

	for _, p := range m.participants {
		if coll, ok := p.(*Collider); ok {
			coll.SetWorldCoordinateUpdated(false)
		}
	}
}

func (m *CollisionManager) emitCollisionEvents(a, b CollisionParticipant, pairID uint64) {
	wasColliding := false
	if prev, ok := m.previousCollisions[pairID]; ok {
		wasColliding = (prev.a == a && prev.b == b) || (prev.a == b && prev.b == a)
	}

	// Collider vs Collider
	if colA, okA := a.(*Collider); okA {
		if colB, okB := b.(*Collider); okB {
			if !wasColliding {
				if !colA.OnCollisionEnter.IsEmpty() {
					colA.OnCollisionEnter.Emit(colB)
				}
				if !colB.OnCollisionEnter.IsEmpty() {
					colB.OnCollisionEnter.Emit(colA)
				}
			} else {
				if !colA.OnCollisionStay.IsEmpty() {
					colA.OnCollisionStay.Emit(colB)
				}
				if !colB.OnCollisionStay.IsEmpty() {
					colB.OnCollisionStay.Emit(colA)
				}
			}
			return
		}
	}

	// Area2D vs Body (Collider or RigidBody2D)
	if area, okA := a.(*Area2D); okA {
		if body, okB := b.(Body); okB {
			m.emitArea2DEvents(area, body, wasColliding)
			return
		}
	}
	if area, okB := b.(*Area2D); okB {
		if body, okA := a.(Body); okA {
			m.emitArea2DEvents(area, body, wasColliding)
		}
	}

	// RigidBody2D vs RigidBody2D: no events (handled by PhysicsWorld)
}

func (m *CollisionManager) emitArea2DEvents(area *Area2D, body Body, wasColliding bool) {
	ev := Area2DBodyEvent{Body: body}
	if !wasColliding {
		if !area.OnBodyEntered.IsEmpty() {
			area.OnBodyEntered.Emit(ev)
		}
	} else {
		if !area.OnBodyStay.IsEmpty() {
			area.OnBodyStay.Emit(ev)
		}
	}
}

func (m *CollisionManager) emitExitEvents(a, b CollisionParticipant) {
	if colA, okA := a.(*Collider); okA {
		if colB, okB := b.(*Collider); okB {
			if !colA.OnCollisionExit.IsEmpty() {
				colA.OnCollisionExit.Emit(colB)
			}
			if !colB.OnCollisionExit.IsEmpty() {
				colB.OnCollisionExit.Emit(colA)
			}
			return
		}
	}
	if area, okA := a.(*Area2D); okA {
		if body, okB := b.(Body); okB {
			if !area.OnBodyExited.IsEmpty() {
				area.OnBodyExited.Emit(Area2DBodyEvent{Body: body})
			}
			return
		}
	}
	if area, okB := b.(*Area2D); okB {
		if body, okA := a.(Body); okA {
			if !area.OnBodyExited.IsEmpty() {
				area.OnBodyExited.Emit(Area2DBodyEvent{Body: body})
			}
		}
	}
}
