package tilemap

import (
	"container/heap"
	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/lafriks/go-tiled"
)

// PathNode is a cell in the path (tile coordinates).
type PathNode struct {
	X, Y int
}

// Pathfinder performs A* pathfinding on a 2D walkability grid.
type Pathfinder struct {
	width          int
	height         int
	walkable       [][]bool
	allowDiagonals bool // if true, uses 8 directions; otherwise 4
}

// NewPathfinder creates a pathfinder with the given grid dimensions.
// All cells start as walkable; use SetWalkable to mark blocked cells.
func NewPathfinder(width, height int) *Pathfinder {
	grid := make([][]bool, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]bool, width)
		for x := 0; x < width; x++ {
			grid[y][x] = true
		}
	}
	return &Pathfinder{width: width, height: height, walkable: grid}
}

// SetAllowDiagonals enables or disables 8-direction (diagonal) movement.
// When enabled, diagonal moves require both adjacent cardinal cells to be walkable (no corner cutting).
func (p *Pathfinder) SetAllowDiagonals(allow bool) {
	p.allowDiagonals = allow
}

// AllowDiagonals returns whether diagonal movement is enabled.
func (p *Pathfinder) AllowDiagonals() bool {
	return p.allowDiagonals
}

// SetWalkable marks the cell at (x, y) as walkable or blocked.
func (p *Pathfinder) SetWalkable(x, y int, walkable bool) {
	if x >= 0 && x < p.width && y >= 0 && y < p.height {
		p.walkable[y][x] = walkable
	}
}

// IsWalkable returns true if the cell can be traversed.
func (p *Pathfinder) IsWalkable(x, y int) bool {
	if x < 0 || x >= p.width || y < 0 || y >= p.height {
		return false
	}
	return p.walkable[y][x]
}

// Width returns the grid width.
func (p *Pathfinder) Width() int {
	return p.width
}

// Height returns the grid height.
func (p *Pathfinder) Height() int {
	return p.height
}

// FindPath returns a path from (startX, startY) to (endX, endY) in tile coordinates.
// Returns nil if no path exists. The path includes start and end.
// Uses 4 or 8 directions depending on AllowDiagonals.
func (p *Pathfinder) FindPath(startX, startY, endX, endY int) []PathNode {
	if !p.IsWalkable(startX, startY) || !p.IsWalkable(endX, endY) {
		return nil
	}
	if startX == endX && startY == endY {
		return []PathNode{{X: startX, Y: startY}}
	}

	heuristic := func(ax, ay, bx, by int) int {
		dx := ax - bx
		if dx < 0 {
			dx = -dx
		}
		dy := ay - by
		if dy < 0 {
			dy = -dy
		}
		if p.allowDiagonals {
			if dx > dy {
				return dx * 10
			}
			return dy * 10
		}
		return (dx + dy) * 10
	}

	openSet := &pfNodeHeap{}
	heap.Init(openSet)
	start := &pfNode{x: startX, y: startY, g: 0, f: heuristic(startX, startY, endX, endY)}
	heap.Push(openSet, start)
	visited := make(map[int]bool)
	key := func(x, y int) int { return y*p.width + x }
	visited[key(startX, startY)] = true

	dirs := p.getDirections()

	for openSet.Len() > 0 {
		cur := heap.Pop(openSet).(*pfNode)
		if cur.x == endX && cur.y == endY {
			var path []PathNode
			for n := cur; n != nil; n = n.parent {
				path = append(path, PathNode{X: n.x, Y: n.y})
			}
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path
		}

		for _, d := range dirs {
			nx, ny := cur.x+d[0], cur.y+d[1]
			if !p.IsWalkable(nx, ny) {
				continue
			}
			cost := 10
			if p.allowDiagonals && d[0] != 0 && d[1] != 0 {
				if !p.IsWalkable(cur.x+d[0], cur.y) || !p.IsWalkable(cur.x, cur.y+d[1]) {
					continue
				}
				cost = 14
			}
			k := key(nx, ny)
			if visited[k] {
				continue
			}
			visited[k] = true
			g := cur.g + cost
			neighbor := &pfNode{x: nx, y: ny, g: g, f: g + heuristic(nx, ny, endX, endY), parent: cur}
			heap.Push(openSet, neighbor)
		}
	}
	return nil
}

func (p *Pathfinder) getDirections() [][2]int {
	if p.allowDiagonals {
		return [][2]int{
			{0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1},
		}
	}
	return [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
}

// PathToWorld converts a path in tile coordinates to world positions (center of each tile).
func PathToWorld(path []PathNode, tileWidth, tileHeight float64) []math2D.Vector2D {
	if len(path) == 0 {
		return nil
	}
	result := make([]math2D.Vector2D, len(path))
	for i, n := range path {
		result[i] = math2D.NewVector2D(
			float64(n.X)*tileWidth+tileWidth/2,
			float64(n.Y)*tileHeight+tileHeight/2,
		)
	}
	return result
}

type pfNode struct {
	x, y   int
	g, f   int
	parent *pfNode
}

type pfNodeHeap []*pfNode

func (h pfNodeHeap) Len() int           { return len(h) }
func (h pfNodeHeap) Less(i, j int) bool { return h[i].f < h[j].f }
func (h pfNodeHeap) Swap(i, j int)     { h[i], h[j] = h[j], h[i] }

func (h *pfNodeHeap) Push(x any) {
	*h = append(*h, x.(*pfNode))
}
func (h *pfNodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}


// BuildPathfinderFromTileLayer creates a Pathfinder from a TileMapNode's tile layer.
// layerName is the Tiled layer name. blockNonEmpty: if true, any non-empty tile blocks;
// if false, only tiles with collision objects (from the tileset) block.
func BuildPathfinderFromTileLayer(tm *TileMapNode, layerName string, blockNonEmpty bool) *Pathfinder {
	if tm == nil || tm.MapData == nil {
		return nil
	}
	m := tm.MapData
	w, h := m.Width, m.Height
	pf := NewPathfinder(w, h)

	var targetLayer *tiled.Layer
	for _, layer := range m.Layers {
		if layer.Name == layerName {
			targetLayer = layer
			break
		}
	}
	if targetLayer == nil || targetLayer.Tiles == nil || len(targetLayer.Tiles) < w*h {
		return pf
	}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			idx := y*w + x
			tile := targetLayer.Tiles[idx]
			blocked := false
			if blockNonEmpty {
				blocked = !tile.IsNil()
			} else {
				if !tile.IsNil() && tile.Tileset != nil {
					if ts, ok := tm.tilesets[tile.Tileset.Name]; ok {
						for _, g := range ts.GetTileCollisionGroups(tile.ID) {
							if len(g.Objects) > 0 {
								blocked = true
								break
							}
						}
					}
				}
			}
			pf.SetWalkable(x, y, !blocked)
		}
	}
	return pf
}
