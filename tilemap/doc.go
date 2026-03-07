// Package tilemap provides data structures for tile-based 2D maps.
//
// TileMapNode wraps Tiled TMX maps with rendering, animations, and collision building.
// BuildCollisionsFromObjectLayer and BuildCollisions create colliders from Tiled layers.
//
// Pathfinder performs A* pathfinding on a walkability grid. Use NewPathfinder for manual
// grids or BuildPathfinderFromTileLayer to derive walkability from a TileMapNode layer.
package tilemap
