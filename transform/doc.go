// Package transform provides 2D transform data (position, pivot, rotation, scale)
// and the Transformable interface used by the scene graph for spatial hierarchy.
//
// Transform holds local position (Vector2D), pivot, rotation (radians), and an optional
// ebiten GeoM for scale/skew. Concat composes two transforms (parent * child).
// Transformable is implemented by Node2D and allows the engine to query GetTransform
// and GetWorldTransform for rendering and collision.
package transform
