# Ludum

A 2D gameplay framework for [Ebiten](https://ebiten.org) written in Go. It provides a scene graph, layers, camera, sprites, animations, collision detection, input handling, and resource management so you can focus on game logic.

[![Go Reference](https://pkg.go.dev/badge/github.com/LuigiVanacore/ludum.svg)](https://pkg.go.dev/github.com/LuigiVanacore/ludum)

## Features

- **Scene graph**: Hierarchical nodes ([Node], [Node2D]) with local and world transforms
- **Layers**: Group nodes by layer ID and draw order (priority)
- **Camera**: 2D camera with zoom, position, and world/screen coordinate conversion (top-left origin)
- **Sprites & animation**: [Sprite] for static images, [AnimationPlayer] + [AnimationSet] for sprite-sheet animations
- **Collision**: [collision] package with shapes (circle, rect, oriented rect), masks, and callbacks; optional broad-phase for performance
- **Physics**: [physics] package with `RigidBody2D` (Dynamic/Kinematic), forces, and velocity-based simulation
- **Particles**: [particles] system with `EmitterNode` for highly customizable visual effects (life, size, color, spread, velocity)
- **Tweening**: [tween] package for procedural animations with over 30 easing functions (Linear, Sine, Elastic, Bounce, etc.)
- **Events**: [event] generic event system (`Event[T]`) for decoupled node-to-node communication
- **Time**: [Clock] and [Timer] for elapsed time and delayed/looping actions
- **Resources/Save**: [ResourceManager] for images/fonts; supports **Texture Atlases** (`LoadAtlas`, `GetAtlasRegion`) and **Asynchronous Preloading** (`PreloadBatch`) with progress callbacks; [save] for atomic JSON/Binary (Gob) data persistence.
- **UI**: Interaction-ready components including `PanelNode`, `ButtonNode`, `ProgressBarNode`, `SliderNode`, `CheckboxNode`, `TextInputNode`, and `ScrollPanelNode`. Includes a **FocusManager** for keyboard/gamepad navigation and an **AnchorLayout** system for relative positioning (Top, Center, Bottom, Stretch).
- **Audio**: [AudioManager] for sounds (WAV, OGG, MP3) and playback
- **Input**: Cursor position, key/button state via [input] package; gamepad/joystick support (buttons, sticks, standard layout)
- **State machine**: [fsm] for AI or game states
- **Tile map**: [tilemap] data structures for grid-based maps; supports `BuildCollisions` (per-tile), `BuildCollisionsFromObjectLayer`, and **Pathfinder integration** via `BuildWalkableFromLayer` (automatically configures pathfinding grid from obstacle layers).

## Installation

```bash
go get github.com/LuigiVanacore/ludum
```

Requires Go 1.21+ and [Ebiten v2](https://github.com/hajimehoshi/ebiten).

## Quick start

```go
package main

import (
	"github.com/LuigiVanacore/ludum"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine *ludum.Engine
}

func NewGame() *Game {
	engine := ludum.NewEngine()
	// Add your nodes with World.AddNodeToLayer(node, layerIndex) or AddNodeToDefaultLayer(node)
	return &Game{engine: engine}
}

func (g *Game) Update() error { return g.engine.Update() }
func (g *Game) Draw(screen *ebiten.Image) { g.engine.Draw(screen) }
func (g *Game) Layout(w, h int) (int, int) { return g.engine.Layout(w, h) }

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
```

## Architecture overview

- **Engine**: Owns World, ResourceManager, InputManager, Clock, Debug. Implements the main game loop (Update/Draw).
- **World**: Root scene, Layers (stack-based), Camera. Update walks the scene; Draw queues nodes to Layers and draws through the camera.
- **Layers**: Stack-based draw system. AddNodeToLayer(node, index) assigns a node to a layer index (lower = drawn first).
- **Node2D**: Position, rotation, scale in local space; GetWorldTransform() returns the combined transform from root to this node. Use for sprites, camera, and custom drawables.

## Documentation

- **API (godoc)**: [pkg.go.dev/github.com/LuigiVanacore/ludum](https://pkg.go.dev/github.com/LuigiVanacore/ludum)
- **Subpackages**: [math2d](https://pkg.go.dev/github.com/LuigiVanacore/ludum/math2d), [transform](https://pkg.go.dev/github.com/LuigiVanacore/ludum/transform), [collision](https://pkg.go.dev/github.com/LuigiVanacore/ludum/collision), [physics](https://pkg.go.dev/github.com/LuigiVanacore/ludum/physics), [particles](https://pkg.go.dev/github.com/LuigiVanacore/ludum/particles), [tween](https://pkg.go.dev/github.com/LuigiVanacore/ludum/tween), [event](https://pkg.go.dev/github.com/LuigiVanacore/ludum/event), [input](https://pkg.go.dev/github.com/LuigiVanacore/ludum/input), [fsm](https://pkg.go.dev/github.com/LuigiVanacore/ludum/fsm), [save](https://pkg.go.dev/github.com/LuigiVanacore/ludum/save), [tilemap](https://pkg.go.dev/github.com/LuigiVanacore/ludum/tilemap), [ui](https://pkg.go.dev/github.com/LuigiVanacore/ludum/ui)

## Layers

Layer indices define draw order (lower = drawn first). Use `World.AddNodeToLayer(node, layerIndex)` or `World.AddNodeToDefaultLayer(node)` for the default layer (index 0). Remove nodes with `World.RemoveNode(node)` or clear a layer with `World.ClearLayer(layerIndex)`.

## Collision

1. Create a [collision.CollisionManager] (e.g. NewCollisionManager()) and set its CellSize if needed.
2. Create colliders with [collision.NewCollider](shape, mask) and add them with manager.AddCollider(collider).
3. Subscribe to [collision.Collider].OnCollisionEnter, OnCollisionStay, or OnCollisionExit (event.Event[*Collider]).
4. Each frame, call manager.CheckCollision(); or set World.SetPostUpdate to a function that calls it (avoids import cycles).

## Physics and Collision Integration

To use [physics.PhysicsWorld] together with [collision.CollisionManager], wire both into the game loop via `World.SetPostUpdate`:

```go
engine := ludum.NewEngine()
physicsWorld := physics.NewPhysicsWorld()
collisionMgr := collision.NewCollisionManager()

// Add RigidBody2D / Collider / Area2D to both World and their managers
engine.World().AddNodeToLayer(player, 0)
physicsWorld.AddRigidBody(player)
collisionMgr.AddParticipant(player)

engine.World().SetPostUpdate(func() {
    physicsWorld.Step(ludum.PhysicsDelta()) // fixed timestep matching ebiten.TPS()
    collisionMgr.CheckCollision()                    // emit Enter/Stay/Exit events
})
```

The order matters: run `PhysicsWorld.Step` first (updates positions), then `CollisionManager.CheckCollision` (evaluates collisions at the new positions). See [example/physics](example/physics) for a full runnable demo.

## Development

To run the linter (requires [golangci-lint](https://golangci-lint.run/) to be installed):

```bash
golangci-lint run ./...
```

## Contributing

Contributions of any kind are welcome. Open an issue or a pull request to propose improvements or report problems.

### Naming conventions

- **Packages**: Use short, all-lowercase names with only letters and numbers—no underscores or `mixedCaps` in the package name. Multi-word names are usually a single unbroken word (for example `tabwriter`, not `tab_writer` or `tabWriter`). See [Effective Go — package names](https://go.dev/doc/effective_go#package-names) and [Google Go Style — package names](https://google.github.io/styleguide/go/decisions#package-names).
- **Source files**: `snake_case` file names are widely used in the Go tree and in many projects; matching that style for new files is a good default ([discussion](https://github.com/golang/go/issues/36060)).
- **Module path**: Prefer **flatcase** path segments (this module is `github.com/LuigiVanacore/ludum`). Tools and proxies may use `!`-encoding for uppercase letters in paths; your `go.mod` uses the normal GitHub path. See `go help modules` and the [module reference](https://go.dev/ref/mod).

Changing the module path again would break existing imports. New code in this repo should follow the conventions above where it fits the existing layout.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details. It extends [Ebiten](https://ebiten.org), which uses the same license.
