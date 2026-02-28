# ebiten_extended

A gameplay framework extending [Ebiten](https://ebiten.org) for 2D games.

## Implemented features

- Sprite
- Animation
- Scene graph
- Collisions (register callbacks with `Collider.SetOnCollision`; run `collision.CollisionManager().CheckCollision()` each frame, or set `GameManager().World().SetPostUpdate(func() { collision.CollisionManager().CheckCollision() })`)
- Time management (Clock, Timer)
- Resource management
- Text / labels
- Audio (see roadmap)

## Layer IDs

Use `ebiten_extended.MinLayerID` (2) or higher for custom layers. Add nodes to the default layer with `ebiten_extended.AddNodeToDefaultLayer(node)`.
