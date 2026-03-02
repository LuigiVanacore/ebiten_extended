# ebiten_extended

<<<<<<< HEAD
A 2D gameplay framework for [Ebiten](https://ebiten.org) written in Go. It provides a scene graph, layers, camera, sprites, animations, collision detection, input handling, and resource management so you can focus on game logic.

[![Go Reference](https://pkg.go.dev/badge/github.com/LuigiVanacore/ebiten_extended.svg)](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended)

## Features

- **Scene graph**: Hierarchical nodes ([Node], [Node2D]) with local and world transforms
- **Layers**: Group nodes by layer ID and draw order (priority)
- **Camera**: 2D camera with zoom, position, and world/screen coordinate conversion
- **Sprites & animation**: [SpriteNode] for static images, [AnimationPlayer] + [AnimationSet] for sprite-sheet animations
- **Collision**: [collision] package with shapes (circle, rect), masks, and callbacks; optional broad-phase for performance
- **Time**: [Clock] and [Timer] for elapsed time and delayed/looping actions
- **Resources**: [ResourceManager] for images and fonts (embed or load from bytes)
- **Input**: Cursor position and key/button state via [input] package
- **State machine**: [stateMachine] for AI or game states
- **Tile map**: [tilemap] data structures for grid-based maps

## Installation

=======
`ebiten_extended` is a support framework for [Ebiten](https://ebiten.org) that simplifies 2D game development in Go. The project provides ready-to-use gameplay components to extend the features of the original engine.

## Features

- Sprite rendering
- Animations
- Scene graph
- Collision detection
- Time management
- Resource management
- Audio management

## Installation

Ensure you have Go 1.22 or later installed, then add the module to your project:

>>>>>>> 153f371edcb4dcf68c2d6633071e13a31c6b0c07
```bash
go get github.com/LuigiVanacore/ebiten_extended
```

<<<<<<< HEAD
Requires Go 1.21+ and [Ebiten v2](https://github.com/hajimehoshi/ebiten).

## Quick start
=======
## Quick Start

Minimal example:
>>>>>>> 153f371edcb4dcf68c2d6633071e13a31c6b0c07

```go
package main

import (
<<<<<<< HEAD
    "github.com/LuigiVanacore/ebiten_extended"
    "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
    engine *ebiten_extended.Engine
}

func NewGame() *Game {
    engine := ebiten_extended.NewEngine()
    layer := ebiten_extended.NewLayer(ebiten_extended.MinLayerID, 0, "main")
    // Add your nodes to layer, then:
    engine.World().AddLayer(layer)
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
- **World**: Root scene, list of Layers, Camera. Update walks the scene and calls Updatable.Update; Draw walks and draws Drawable nodes through the camera.
- **Layer**: Has an ID (≥ MinLayerID), priority, and a root SceneNode. All nodes you add go under that root.
- **Node2D**: Position, rotation, scale in local space; GetWorldTransform() returns the combined transform from root to this node. Use for sprites, camera, and custom drawables.

## Documentation

- **API (godoc)**: [pkg.go.dev/github.com/LuigiVanacore/ebiten_extended](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended)
- **Subpackages**: [math2D](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/math2D), [transform](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/transform), [collision](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/collision), [input](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/input), [stateMachine](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/stateMachine), [tilemap](https://pkg.go.dev/github.com/LuigiVanacore/ebiten_extended/tilemap)

## Layer IDs

Use layer IDs ≥ **MinLayerID** (2). IDs 0 and 1 are reserved. Add a node to the default layer with World.AddNodeToDefaultLayer.

## Collision

1. Create a [collision.CollisionManager] (e.g. NewCollisionManager()) and set its CellSize if needed.
2. Create colliders with [collision.NewCollider](shape, mask) and add them with manager.AddCollider(collider).
3. Subscribe to [collision.Collider].OnCollisionEnter, OnCollisionStay, or OnCollisionExit (event.Event[*Collider]).
4. Each frame, call manager.CheckCollision(); or set World.SetPostUpdate to a function that calls it (avoids import cycles).

## License

See the repository for license information. This project extends Ebiten, which is licensed under the Apache License 2.0.
=======
    e "github.com/LuigiVanacore/ebiten_extended"
)

func main() {
    // TODO: implement game loop
    _ = e.World{}
}
```

## Contributing

Contributions of any kind are welcome. Open an issue or a pull request to propose improvements or report problems.
>>>>>>> 153f371edcb4dcf68c2d6633071e13a31c6b0c07
