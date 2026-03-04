# ebiten_extended

A 2D gameplay framework for [Ebiten](https://ebiten.org) written in Go. It provides a scene graph, layers, camera, sprites, animations, collision detection, input handling, and resource management so you can focus on game logic.

[![Go Reference](https://pkg.go.dev/badge/github.com/LuigiVanacore/ebiten_extended.svg)](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended)

## Features

- **Scene graph**: Hierarchical nodes ([Node], [Node2D]) with local and world transforms
- **Layers**: Group nodes by layer ID and draw order (priority)
- **Camera**: 2D camera with zoom, position, and world/screen coordinate conversion (top-left origin)
- **Sprites & animation**: [Sprite] for static images, [AnimationPlayer] + [AnimationSet] for sprite-sheet animations
- **Collision**: [collision] package with shapes (circle, rect), masks, and callbacks; optional broad-phase for performance
- **Time**: [Clock] and [Timer] for elapsed time and delayed/looping actions
- **Resources/Save**: [ResourceManager] for images/fonts; [save] for atomic JSON/Binary (Gob) data persistence.
- **UI**: Fast interaction-ready graphical user interface components including `PanelNode`, `ButtonNode` (with image states), `ProgressBarNode`, and `CheckboxNode`.
- **Audio**: [AudioManager] for sounds (WAV, OGG, MP3) and playback
- **Input**: Cursor position, key/button state via [input] package; gamepad/joystick support (buttons, sticks, standard layout)
- **State machine**: [fsm] for AI or game states
- **Tile map**: [tilemap] data structures for grid-based maps

## Installation

```bash
go get github.com/LuigiVanacore/ebiten_extended
```

Requires Go 1.21+ and [Ebiten v2](https://github.com/hajimehoshi/ebiten).

## Quick start

```go
package main

import (
	"github.com/LuigiVanacore/ebiten_extended"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	engine *ebiten_extended.Engine
}

func NewGame() *Game {
	engine := ebiten_extended.NewEngine()
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

- **API (godoc)**: [pkg.go.dev/github.com/LuigiVanacore/ebiten_extended](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended)
- **Subpackages**: [math2D](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/math2D), [transform](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/transform), [collision](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/collision), [input](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/input), [fsm](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/fsm), [save](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/save), [tilemap](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/tilemap), [ui](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/ui)

## Layers

Layer indices define draw order (lower = drawn first). Use `World.AddNodeToLayer(node, layerIndex)` or `World.AddNodeToDefaultLayer(node)` for the default layer (index 0).

## Collision

1. Create a [collision.CollisionManager] (e.g. NewCollisionManager()) and set its CellSize if needed.
2. Create colliders with [collision.NewCollider](shape, mask) and add them with manager.AddCollider(collider).
3. Subscribe to [collision.Collider].OnCollisionEnter, OnCollisionStay, or OnCollisionExit (event.Event[*Collider]).
4. Each frame, call manager.CheckCollision(); or set World.SetPostUpdate to a function that calls it (avoids import cycles).

## Contributing

Contributions of any kind are welcome. Open an issue or a pull request to propose improvements or report problems.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details. It extends [Ebiten](https://ebiten.org), which uses the same license.
