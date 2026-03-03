// Package ebiten_extended provides a 2D gameplay framework built on top of
// [Ebiten](https://ebiten.org). It adds scene graph management, layers, camera,
// sprites, animations, collision detection, input handling, and resource loading
// so you can build 2D games without reimplementing common engine features.
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
//	engine := ebiten_extended.NewEngine()
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
// # Subpackages
//
//   - [math2D]: 2D math (vectors, shapes, segments, rectangles, circles).
//   - [transform]: Local transform (position, pivot, rotation, scale).
//   - [collision]: Colliders, shapes, masks, and collision detection.
//   - [input]: Input manager and cursor position (enable with SetMouseEnabled).
//   - [stateMachine]: Simple state machine for AI or game states.
//   - [tilemap]: Tile map data structures.
//   - [event]: Event bus (if enabled).
package ebiten_extended
