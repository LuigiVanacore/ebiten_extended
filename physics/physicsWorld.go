package physics

import (
	"math"

	"github.com/LuigiVanacore/ebiten_extended/collision"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

// PhysicsWorld manages RigidBody2D and runs integration + collision resolution each frame.
type PhysicsWorld struct {
	rigidBodies []*RigidBody2D
	Gravity     math2D.Vector2D
	CellSize    int
}

// NewPhysicsWorld creates a PhysicsWorld with default CellSize 100.
func NewPhysicsWorld() *PhysicsWorld {
	return &PhysicsWorld{
		rigidBodies: make([]*RigidBody2D, 0),
		Gravity:     math2D.NewVector2D(0, 980),
		CellSize:    100,
	}
}

// AddRigidBody adds a body to the simulation. Panics if body or its shape is nil.
func (w *PhysicsWorld) AddRigidBody(body *RigidBody2D) {
	if body == nil || body.GetShape() == nil {
		panic("physics: AddRigidBody body and shape must not be nil")
	}
	w.rigidBodies = append(w.rigidBodies, body)
}

// RemoveRigidBody removes a body from the simulation.
func (w *PhysicsWorld) RemoveRigidBody(body *RigidBody2D) {
	for i, b := range w.rigidBodies {
		if b == body {
			w.rigidBodies[i] = w.rigidBodies[len(w.rigidBodies)-1]
			w.rigidBodies = w.rigidBodies[:len(w.rigidBodies)-1]
			return
		}
	}
}

// Step advances the simulation by dt. Integrates velocity, applies gravity, and resolves overlaps.
func (w *PhysicsWorld) Step(dt float64) {
	// 1. Integrate
	for _, body := range w.rigidBodies {
		if body.UsesGravity {
			g := math2D.NewVector2D(
				w.Gravity.X()*body.GravityScale*dt,
				w.Gravity.Y()*body.GravityScale*dt,
			)
			body.velocity = math2D.AddVectors(body.velocity, g)
		}
		pos := body.GetPosition()
		body.SetPosition(
			pos.X()+body.velocity.X()*dt,
			pos.Y()+body.velocity.Y()*dt,
		)
	}

	// 2. Resolve overlaps (iterate 3 times for stability)
	for iter := 0; iter < 3; iter++ {
		grid := w.broadPhase()
		checked := make(map[uint64]bool)

		for _, candidates := range grid {
			for i := 0; i < len(candidates); i++ {
				for j := i + 1; j < len(candidates); j++ {
					a, b := candidates[i], candidates[j]
					if a == b {
						continue
					}
					pairID := combineIDs(a.GetID(), b.GetID())
					if checked[pairID] {
						continue
					}
					checked[pairID] = true

					sa, sb := a.GetShape(), b.GetShape()
					sa.UpdateTransform(a.GetWorldTransform())
					sb.UpdateTransform(b.GetWorldTransform())

					res, ok := collision.ShapeCollisionResult(sa, sb)
					if !ok || !res.Overlapping {
						continue
					}

					w.resolveOverlap(a, b, res)
				}
			}
		}
	}
}

func (w *PhysicsWorld) broadPhase() map[uint64][]*RigidBody2D {
	grid := make(map[uint64][]*RigidBody2D)
	cellSize := float64(w.CellSize)
	for _, body := range w.rigidBodies {
		sa := body.GetShape()
		sa.UpdateTransform(body.GetWorldTransform())
		minX, minY, maxX, maxY := collision.ShapeAABB(sa)
		cellMinX := int(math.Floor(minX / cellSize))
		cellMaxX := int(math.Floor(maxX / cellSize))
		cellMinY := int(math.Floor(minY / cellSize))
		cellMaxY := int(math.Floor(maxY / cellSize))
		for cx := cellMinX; cx <= cellMaxX; cx++ {
			for cy := cellMinY; cy <= cellMaxY; cy++ {
				key := (uint64(uint32(cx)) << 32) | uint64(uint32(cy))
				grid[key] = append(grid[key], body)
			}
		}
	}
	return grid
}

func (w *PhysicsWorld) resolveOverlap(a, b *RigidBody2D, res collision.CollisionResult) {
	// Static bodies are not pushed; dynamic body gets full separation
	var pushA, pushB math2D.Vector2D
	if a.Static && b.Static {
		return
	}
	if a.Static {
		pushA = math2D.ZeroVector2D()
		pushB = res.Normal.Negate().MultiplyScalar(res.Depth)
	} else if b.Static {
		pushA = res.Normal.MultiplyScalar(res.Depth)
		pushB = math2D.ZeroVector2D()
	} else {
		pushA = res.Normal.MultiplyScalar(res.Depth * 0.5)
		pushB = res.Normal.Negate().MultiplyScalar(res.Depth * 0.5)
	}

	posA := a.GetPosition()
	a.SetPosition(posA.X()+pushA.X(), posA.Y()+pushA.Y())

	posB := b.GetPosition()
	b.SetPosition(posB.X()+pushB.X(), posB.Y()+pushB.Y())

	// Zero velocity component along collision normal so bodies "land" instead of jittering.
	// Normal points from B toward A; skip for static bodies.
	va := a.GetVelocity()
	if !a.Static && math2D.DotProduct(va, res.Normal) < 0 {
		a.SetVelocity(math2D.SubtractVectors(va, res.Normal.MultiplyScalar(math2D.DotProduct(va, res.Normal))))
	}
	vb := b.GetVelocity()
	if !b.Static && math2D.DotProduct(vb, res.Normal) > 0 {
		b.SetVelocity(math2D.SubtractVectors(vb, res.Normal.MultiplyScalar(math2D.DotProduct(vb, res.Normal))))
	}
}

func combineIDs(id1, id2 uint64) uint64 {
	if id1 < id2 {
		return (id1 << 32) | id2
	}
	return (id2 << 32) | id1
}
