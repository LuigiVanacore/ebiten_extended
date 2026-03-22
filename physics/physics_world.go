package physics

import (
	"errors"
	"math"

	"github.com/LuigiVanacore/ludum/collision"
	"github.com/LuigiVanacore/ludum/math2d"
)

// PhysicsWorld manages RigidBody2D and runs integration + collision resolution each frame.
type PhysicsWorld struct {
	rigidBodies  []*RigidBody2D
	Gravity      math2d.Vector2D
	CellSize     int
	checkedPairs map[uint64]bool           // reused each Step to avoid per-frame allocation
	grid         map[uint64][]*RigidBody2D // reused for broad-phase spatial hash
}

// NewPhysicsWorld creates a PhysicsWorld with default CellSize 100.
func NewPhysicsWorld() *PhysicsWorld {
	return &PhysicsWorld{
		rigidBodies:  make([]*RigidBody2D, 0),
		Gravity:      math2d.NewVector2D(0, 980),
		CellSize:     100,
		checkedPairs: make(map[uint64]bool),
		grid:         make(map[uint64][]*RigidBody2D),
	}
}

// AddRigidBody adds a body to the simulation.
func (w *PhysicsWorld) AddRigidBody(body *RigidBody2D) error {
	if body == nil || body.GetShape() == nil {
		return errors.New("physics: AddRigidBody body and shape must not be nil")
	}
	w.rigidBodies = append(w.rigidBodies, body)
	return nil
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
	// 1. Integrate (skip kinematic and static)
	for _, body := range w.rigidBodies {
		if body.Static || body.Kinematic {
			continue
		}
		if body.UsesGravity {
			g := math2d.NewVector2D(
				w.Gravity.X()*body.GravityScale*dt,
				w.Gravity.Y()*body.GravityScale*dt,
			)
			body.velocity = math2d.AddVectors(body.velocity, g)
		}
		pos := body.GetPosition()
		body.SetPosition(
			pos.X()+body.velocity.X()*dt,
			pos.Y()+body.velocity.Y()*dt,
		)
	}

	// 2. Resolve overlaps (iterate 3 times for stability)
	for iter := 0; iter < 3; iter++ {
		w.broadPhase()
		clear(w.checkedPairs)

		for _, candidates := range w.grid {
			for i := 0; i < len(candidates); i++ {
				for j := i + 1; j < len(candidates); j++ {
					a, b := candidates[i], candidates[j]
					if a == b {
						continue
					}
					pairID := collision.CombineIDs(a.GetID(), b.GetID())
					if w.checkedPairs[pairID] {
						continue
					}
					w.checkedPairs[pairID] = true

					sa, sb := a.GetShape(), b.GetShape()

					res, ok := collision.ShapeCollisionResult(sa, a.GetWorldTransform(), sb, b.GetWorldTransform())
					if !ok || !res.Overlapping {
						continue
					}

					w.resolveOverlap(a, b, res)
				}
			}
		}
	}
}

func (w *PhysicsWorld) broadPhase() {
	clear(w.grid)
	cellSize := float64(w.CellSize)
	for _, body := range w.rigidBodies {
		sa := body.GetShape()
		minX, minY, maxX, maxY := collision.ShapeAABB(sa, body.GetWorldTransform())
		cellMinX := int(math.Floor(minX / cellSize))
		cellMaxX := int(math.Floor(maxX / cellSize))
		cellMinY := int(math.Floor(minY / cellSize))
		cellMaxY := int(math.Floor(maxY / cellSize))
		for cx := cellMinX; cx <= cellMaxX; cx++ {
			for cy := cellMinY; cy <= cellMaxY; cy++ {
				key := (uint64(uint32(cx)) << 32) | uint64(uint32(cy))
				w.grid[key] = append(w.grid[key], body)
			}
		}
	}
}

func (w *PhysicsWorld) resolveOverlap(a, b *RigidBody2D, res collision.CollisionResult) {
	// Static and kinematic bodies are not pushed
	solidA := a.Static || a.Kinematic
	solidB := b.Static || b.Kinematic
	if solidA && solidB {
		return
	}
	var pushA, pushB math2d.Vector2D
	if solidA {
		pushA = math2d.ZeroVector2D()
		pushB = res.Normal.Negate().MultiplyScalar(res.Depth)
	} else if solidB {
		pushA = res.Normal.MultiplyScalar(res.Depth)
		pushB = math2d.ZeroVector2D()
	} else {
		pushA = res.Normal.MultiplyScalar(res.Depth * 0.5)
		pushB = res.Normal.Negate().MultiplyScalar(res.Depth * 0.5)
	}

	posA := a.GetPosition()
	a.SetPosition(posA.X()+pushA.X(), posA.Y()+pushA.Y())

	posB := b.GetPosition()
	b.SetPosition(posB.X()+pushB.X(), posB.Y()+pushB.Y())

	// Velocity response: restitution (bounce) on normal, friction on tangent.
	// Restitution: use the dynamic body's value when one is static, else min of both.
	// Friction: average of both.
	restitution := b.Restitution
	if b.Static {
		restitution = a.Restitution
	} else if !a.Static && a.Restitution < restitution {
		restitution = a.Restitution
	}
	friction := (a.Friction + b.Friction) / 2
	if friction > 1 {
		friction = 1
	}

	applyVelocityResponse := func(body *RigidBody2D, v math2d.Vector2D, normalCompIntoSurface float64) {
		if body.Static || body.Kinematic {
			return
		}
		// Normal: reflect with restitution only when moving into surface.
		var newNormalComp float64
		if normalCompIntoSurface < 0 {
			newNormalComp = -restitution * normalCompIntoSurface
		} else {
			newNormalComp = normalCompIntoSurface
		}

		// Tangent: always reduce by friction when in contact.
		tangent := math2d.SubtractVectors(v, res.Normal.MultiplyScalar(normalCompIntoSurface))
		vNew := math2d.AddVectors(
			res.Normal.MultiplyScalar(newNormalComp),
			tangent.MultiplyScalar(1-friction),
		)
		body.SetVelocity(vNew)
	}

	va := a.GetVelocity()
	normalCompA := math2d.DotProduct(va, res.Normal)
	if !solidA {
		applyVelocityResponse(a, va, normalCompA)
	}

	vb := b.GetVelocity()
	normalCompB := math2d.DotProduct(vb, res.Normal)
	if !solidB {
		applyVelocityResponse(b, vb, normalCompB)
	}
}
