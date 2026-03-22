// Package ludum provides a 2D gameplay framework built on top of
// [Ebiten](https://ebiten.org). It adds scene graph management, layers, camera,
// sprites, animations, collision detection, input handling, resource loading,
// particles, pooling, scene manager, and transitions so you can build 2D games
// without reimplementing common engine features.
//
// # Core concepts
//
//   - [Engine]: Entry point that owns the [World], [ResourceManager], [AudioManager], input, and clock.
//     Implement ebiten.Game by delegating Update/Draw/Layout to the engine.
//   - [World]: Holds the scene tree, [Layers], and [Camera]. Update runs the game tick;
//     Draw renders the scene through the camera.
//   - [SceneNode]: Interface for nodes in the scene graph ([Node], [Node2D]).
//   - [Node2D]: 2D node with local transform (position, rotation, scale) and cached
//     world transform. Base for [SpriteNode], [Camera], and custom drawables.
//   - [Layers]: Stack-based draw system. Use [World.AddNodeToLayer] with a layer index
//     (lower = drawn first). [World.AddNodeToDefaultLayer] uses index 0.
//   - [Drawable] / [Updatable]: Interfaces for nodes that are drawn or updated each frame.
//
// # Quick start
//
// Create an engine, add a layer and a node, then run the game:
//
//	engine := ludum.NewEngine()
//	engine.World().AddNodeToLayer(myNode, 0)
//	// In your game's Update: engine.Update()
//	// In your game's Draw: engine.Draw(screen)
//
// # Collision
//
// Use the [collision] package: create [collision.Collider]s with shapes and masks,
// subscribe to [collision.Collider.OnCollisionEnter], OnCollisionStay, OnCollisionExit.
// Add a [collision.CollisionManager], add colliders to it, and call CheckCollision each
// frame (e.g. in [World.SetPostUpdate]) to run broad/narrow phase and emit events.
//
// # Resources and animation
//
// [ResourceManager] loads images (AddImage) and fonts (LoadFont). [AnimationSet] and
// [AnimationPlayer] handle sprite-sheet animations; attach an AnimationPlayer as a
// child of a Node2D to draw animated sprites.
//
// # Scene manager and transitions
//
// [SceneManager] manages a stack of [Scene]s and delegates Engine Update/Draw to the
// current scene. Set it on the engine with [Engine.SetSceneManager]. Use PushScene,
// ReplaceScene, and PopScene to navigate. [Transition] provides fade effects between
// scene changes: call SetTransitionDuration (e.g. 0.3 seconds) to enable fade-out,
// swap, then fade-in. SetTransitionColor customizes the overlay (default black).
//
// # Particles
//
// The [particles] package provides [particles.ParticleEmitter] for configurable particle
// effects (lifetime, velocity, scale, color). Use [particles.ParticleEmitterNode] as a
// Node2D child to position and draw particles in the scene. Set emission rate for
// continuous emission, or call Burst for one-shot effects.
//
// # Pooling
//
// [SpritePool] reuses [Sprite] instances to reduce allocations (e.g. projectiles, bullets).
// Create with NewSpritePool, then Get/Put sprites. [utils.Pool] is a generic pool for
// any type when you need custom factory and reset logic.
//
// # Subpackages
//
//   - [math2d]: 2D math (vectors, shapes, segments, rectangles, circles).
//   - [transform]: Local transform (position, pivot, rotation, scale).
//   - [collision]: Colliders, shapes, masks, and collision detection.
//   - [input]: Input manager and cursor position (enable with SetMouseEnabled).
//   - [fsm]: Generics powered state machine for AI or game states.
//   - [tilemap]: Tile map data structures.
//   - [particles]: ParticleEmitter and ParticleEmitterNode for visual effects.
//   - [utils]: Generic Pool for object reuse.
//   - [event]: Event bus (if enabled).
package ludum
