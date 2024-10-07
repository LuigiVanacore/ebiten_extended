package tilemap


type EbitenMap struct {
	TileWidth  int
	TileHeight int
	MapHeight  int
	MapWidth   int
	Layers     [][]int
}