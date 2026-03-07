// Package collision provides 2D collision detection with shapes, layers (masks), and lifecycle events.
//
// Create colliders with NewCollider(shape, mask). Shapes implement CollisionShape (e.g. CollisionCircle,
// CollisionRect, CollisionOrientedRect, CollisionPolygon for convex polygons). CollisionMask uses identity and mask bits
// to filter which pairs can collide (CanCollideWith). Use presets (MaskPlayer, MaskEnemy, MaskWorld, etc.)
// or NewPresetMask for common layer configurations.
//
// CollisionManager runs the broad phase (spatial hash grid) and narrow phase, and emits Enter/Stay/Exit
// events via Collider.OnCollisionEnter, OnCollisionStay, OnCollisionExit (event.Event[*Collider]).
// Add colliders to a manager with AddCollider; call CheckCollision each frame (e.g. in World.SetPostUpdate).
//
// Spatial queries: OverlapPoint, OverlapCircle, OverlapRect for point/circle/rectangle overlap tests;
// Raycast for segment-vs-shape hits (sorted by distance).
//
// Movement/swept collision helpers are in movementCollisionDetection.go (e.g. MovingCircleCircleCollide).
package collision
