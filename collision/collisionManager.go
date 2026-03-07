package collision

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

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
	// internal maps reused across CheckCollision calls to avoid per-frame allocations.
	grid             map[uint64][]CollisionParticipant
	currentCollisions map[uint64]collisionPair
	checkedPairs     map[uint64]bool
}

// NewCollisionManager returns a new collision manager with default CellSize 100.
func NewCollisionManager() *CollisionManager {
	return &CollisionManager{
		participants:      make([]CollisionParticipant, 0),
		CellSize:          100,
		previousCollisions: make(map[uint64]collisionPair),
		grid:              make(map[uint64][]CollisionParticipant),
		currentCollisions: make(map[uint64]collisionPair),
		checkedPairs:      make(map[uint64]bool),
	}
}

// AddCollider registers a collider. Kept for backward compatibility.
func (m *CollisionManager) AddCollider(collider *Collider) {
	m.AddParticipant(collider)
}

// RemoveCollider unregisters a collider. Kept for backward compatibility.
func (m *CollisionManager) RemoveCollider(collider *Collider) {
	m.RemoveParticipant(collider)
}

// AddParticipant registers a collision participant (Collider, Area2D, or RigidBody2D).
func (m *CollisionManager) AddParticipant(p CollisionParticipant) {
	if p == nil {
		return
	}
	m.participants = append(m.participants, p)
}

// RemoveParticipant unregisters a collision participant.
func (m *CollisionManager) RemoveParticipant(p CollisionParticipant) {
	if p == nil {
		return
	}
	for i, current := range m.participants {
		if current == p {
			m.participants[i] = m.participants[len(m.participants)-1]
			m.participants = m.participants[:len(m.participants)-1]
			break
		}
	}
	// Drop stale collision pairs that contain removed participant.
	for pairID, pair := range m.previousCollisions {
		if pair.a == p || pair.b == p {
			delete(m.previousCollisions, pairID)
		}
	}
}

// RaycastResult holds a participant hit by a ray/segment and the distance from the ray origin.
type RaycastResult struct {
	Participant CollisionParticipant
	Distance    float64 // distance from segment start to hit (0 if at start)
}

// Raycast tests a segment against all participants and returns those hit, sorted by distance from start.
// Use for line-of-sight, shooting, pathfinding. The segment is in world coordinates.
func (m *CollisionManager) Raycast(start, end math2D.Vector2D) []RaycastResult {
	seg := math2D.NewSegment(start, end)
	dir := math2D.SubtractVectors(end, start)
	segLen := math.Sqrt(math2D.DotProduct(dir, dir))
	if segLen < 1e-9 {
		segLen = 1e-9
	}
	dirNorm := dir.Normalize()
	var results []RaycastResult
	for _, p := range m.participants {
		shape := p.GetShape()
		if shape == nil {
			continue
		}
		if !segmentOverlapsShape(seg, shape, p.GetWorldTransform()) {
			continue
		}
		// Approximate distance: use closest point on segment to shape center
		center := p.GetWorldPosition()
		toCenter := math2D.SubtractVectors(center, start)
		proj := math2D.DotProduct(toCenter, dirNorm)
		if proj < 0 {
			proj = 0
		}
		if proj > segLen {
			proj = segLen
		}
		closest := math2D.AddVectors(start, dirNorm.MultiplyScalar(proj))
		dist := math.Sqrt(math2D.DotProduct(
			math2D.SubtractVectors(closest, start),
			math2D.SubtractVectors(closest, start),
		))
		results = append(results, RaycastResult{Participant: p, Distance: dist})
	}
	// Sort by distance
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Distance < results[i].Distance {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
	return results
}

// OverlapPoint returns all participants whose shape contains the given world-space point.
func (m *CollisionManager) OverlapPoint(point math2D.Vector2D) []CollisionParticipant {
	var result []CollisionParticipant
	for _, p := range m.participants {
		shape := p.GetShape()
		if shape == nil {
			continue
		}
		if overlapPointShape(point, shape, p.GetWorldTransform()) {
			result = append(result, p)
		}
	}
	return result
}

// OverlapCircle returns all participants whose shape overlaps the given world-space circle.
func (m *CollisionManager) OverlapCircle(center math2D.Vector2D, radius float64) []CollisionParticipant {
	query := math2D.NewCircle(center, radius)
	var result []CollisionParticipant
	for _, p := range m.participants {
		shape := p.GetShape()
		if shape == nil {
			continue
		}
		if overlapCircleShape(query, shape, p.GetWorldTransform()) {
			result = append(result, p)
		}
	}
	return result
}

// OverlapRect returns all participants whose shape overlaps the given world-space axis-aligned rectangle.
// The rectangle is defined by position (top-left) and size.
func (m *CollisionManager) OverlapRect(rect math2D.Rectangle) []CollisionParticipant {
	var result []CollisionParticipant
	for _, p := range m.participants {
		shape := p.GetShape()
		if shape == nil {
			continue
		}
		if overlapRectShape(rect, shape, p.GetWorldTransform()) {
			result = append(result, p)
		}
	}
	return result
}

// CombineIDs produces a single deterministic uint64 key from two node IDs,
// independent of argument order. Used for collision pair deduplication.
func CombineIDs(id1, id2 uint64) uint64 {
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
	return ShapeCollides(sa, a.GetWorldTransform(), sb, b.GetWorldTransform())
}

// CheckCollision evaluates intersections and emits lifecycle events.
func (m *CollisionManager) CheckCollision() {
	clear(m.grid)
	cellSize := float64(m.CellSize)

	for _, p := range m.participants {
		shape := p.GetShape()
		if shape == nil {
			continue
		}
		minX, minY, maxX, maxY := ShapeAABB(shape, p.GetWorldTransform())
		cellMinX := int(math.Floor(minX / cellSize))
		cellMaxX := int(math.Floor(maxX / cellSize))
		cellMinY := int(math.Floor(minY / cellSize))
		cellMaxY := int(math.Floor(maxY / cellSize))
		for cx := cellMinX; cx <= cellMaxX; cx++ {
			for cy := cellMinY; cy <= cellMaxY; cy++ {
				key := (uint64(uint32(cx)) << 32) | uint64(uint32(cy))
				m.grid[key] = append(m.grid[key], p)
			}
		}
	}

	clear(m.currentCollisions)
	clear(m.checkedPairs)

	for _, cellParticipants := range m.grid {
		for i := 0; i < len(cellParticipants); i++ {
			for j := i + 1; j < len(cellParticipants); j++ {
				a := cellParticipants[i]
				b := cellParticipants[j]

				if a == b {
					continue
				}

				pairID := CombineIDs(a.GetID(), b.GetID())
				if m.checkedPairs[pairID] {
					continue
				}
				m.checkedPairs[pairID] = true

				if !a.CanCollideWith(b) || !b.CanCollideWith(a) {
					continue
				}
				if !isOverlapping(a, b) {
					continue
				}

				m.currentCollisions[pairID] = collisionPair{a: a, b: b}
				m.emitCollisionEvents(a, b, pairID)
			}
		}
	}

	for pairID, pair := range m.previousCollisions {
		if _, still := m.currentCollisions[pairID]; !still {
			m.emitExitEvents(pair.a, pair.b)
		}
	}

	m.previousCollisions, m.currentCollisions = m.currentCollisions, m.previousCollisions
	clear(m.currentCollisions)

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
				if !colA.onCollisionEnter.IsEmpty() {
					colA.onCollisionEnter.Emit(colB)
				}
				if !colB.onCollisionEnter.IsEmpty() {
					colB.onCollisionEnter.Emit(colA)
				}
			} else {
				if !colA.onCollisionStay.IsEmpty() {
					colA.onCollisionStay.Emit(colB)
				}
				if !colB.onCollisionStay.IsEmpty() {
					colB.onCollisionStay.Emit(colA)
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
		if !area.onBodyEntered.IsEmpty() {
			area.onBodyEntered.Emit(ev)
		}
	} else {
		if !area.onBodyStay.IsEmpty() {
			area.onBodyStay.Emit(ev)
		}
	}
}

func (m *CollisionManager) emitExitEvents(a, b CollisionParticipant) {
	if colA, okA := a.(*Collider); okA {
		if colB, okB := b.(*Collider); okB {
			if !colA.onCollisionExit.IsEmpty() {
				colA.onCollisionExit.Emit(colB)
			}
			if !colB.onCollisionExit.IsEmpty() {
				colB.onCollisionExit.Emit(colA)
			}
			return
		}
	}
	if area, okA := a.(*Area2D); okA {
		if body, okB := b.(Body); okB {
			if !area.onBodyExited.IsEmpty() {
				area.onBodyExited.Emit(Area2DBodyEvent{Body: body})
			}
			return
		}
	}
	if area, okB := b.(*Area2D); okB {
		if body, okA := a.(Body); okA {
			if !area.onBodyExited.IsEmpty() {
				area.onBodyExited.Emit(Area2DBodyEvent{Body: body})
			}
		}
	}
}
