// Package physics provides 2D physics simulation with RigidBody2D and PhysicsWorld.
// RigidBody2D bodies have velocity, optional gravity, and collision resolution (no overlap).
// Add RigidBody2D to both the World (for rendering/hierarchy) and PhysicsWorld.
// Call PhysicsWorld.Step(dt) each frame before CollisionManager.CheckCollision().
package physics
