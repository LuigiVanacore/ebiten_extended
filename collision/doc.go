// Package collision provides 2D collision detection with shapes, layers (masks), and lifecycle events.
//
// # Collision Layers and Masks
//
// Each Collider and Area2D has a CollisionMask that controls which pairs can collide. The mask has two parts:
//
//   - Identity: the layer(s) this object belongs to (what it "is"). Use power-of-2 bits (utils.ByteSet).
//   - Mask: the layer(s) this object responds to (what it can collide with).
//
// A pair (A, B) collides only if BOTH allow it:
//   - A.mask must include B.identity (A wants to collide with B's layer)
//   - B.mask must include A.identity (B wants to collide with A's layer)
//
// Example: Player (identity=LayerPlayer, mask=LayerWorld|LayerPickup) collides with World (identity=LayerWorld)
// because Player.mask has LayerWorld and World.mask has LayerPlayer.
//
// Use NewCollisionMask(identity, mask) for custom bit sets, or NewPresetMask(identity, collidesWith...)
// for the preset layers (LayerPlayer, LayerEnemy, LayerWorld, etc.). Presets MaskPlayer, MaskEnemy,
// MaskWorld, MaskPickup, MaskProjectile are ready-to-use for common 2D games.
//
// Custom layers: define constants as 1<<0, 1<<1, … then NewCollisionMask(utils.ByteSet(myLayer), utils.ByteSet(otherLayer)).
//
// # Shapes and Colliders
//
// Create colliders with NewCollider(shape, mask). Shapes implement CollisionShape (e.g. CollisionCircle,
// CollisionRect, CollisionOrientedRect, CollisionPolygon for convex polygons).
//
// # Events and Manager
//
// CollisionManager runs the broad phase (spatial hash grid) and narrow phase, and emits Enter/Stay/Exit
// events via Collider.OnCollisionEnter, OnCollisionStay, OnCollisionExit (event.Event[*Collider]).
// Add colliders with AddCollider; call CheckCollision each frame (e.g. in World.SetPostUpdate).
//
// # Spatial Queries
//
// OverlapPoint, OverlapCircle, OverlapRect for point/circle/rectangle overlap tests;
// Raycast for segment-vs-shape hits (sorted by distance).
//
// # Movement
//
// Swept collision helpers in movementCollisionDetection.go (e.g. MovingCircleCircleCollide).
package collision
