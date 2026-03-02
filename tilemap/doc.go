// Package tilemap provides data structures for tile-based 2D maps.
//
// EbitenMap holds grid dimensions (TileWidth, TileHeight, MapWidth, MapHeight)
// and Layers: each layer is a 2D slice of tile IDs (row-major or as needed).
// Rendering and collision from tile data are left to the game or other packages.
package tilemap
