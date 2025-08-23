# ebiten_extended

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

```bash
go get github.com/LuigiVanacore/ebiten_extended
```

## Quick Start

Minimal example:

```go
package main

import (
    e "github.com/LuigiVanacore/ebiten_extended"
)

func main() {
    // TODO: implement game loop
    _ = e.World{}
}
```

## Contributing

Contributions of any kind are welcome. Open an issue or a pull request to propose improvements or report problems.
