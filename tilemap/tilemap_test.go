package tilemap

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
)

func TestNewTileMapNodeInvalidPath(t *testing.T) {
	result, err := NewTileMapNode("nonexistent_map_does_not_exist.tmx")
	if err == nil {
		t.Error("expected error for nonexistent path")
	}
	if result != nil {
		t.Error("expected nil TileMapNode on error")
	}
}

func TestPathfinder_FindPath(t *testing.T) {
	pf := NewPathfinder(5, 5)
	// Block center
	pf.SetWalkable(2, 2, false)
	path := pf.FindPath(0, 0, 4, 4)
	if path == nil {
		t.Fatal("expected path around obstacle")
	}
	if path[0].X != 0 || path[0].Y != 0 {
		t.Errorf("path start: got (%d,%d)", path[0].X, path[0].Y)
	}
	if path[len(path)-1].X != 4 || path[len(path)-1].Y != 4 {
		t.Errorf("path end: got (%d,%d)", path[len(path)-1].X, path[len(path)-1].Y)
	}
}

func TestPathfinder_NoPath(t *testing.T) {
	pf := NewPathfinder(3, 3)
	pf.SetWalkable(1, 0, false)
	pf.SetWalkable(1, 1, false)
	pf.SetWalkable(1, 2, false)
	path := pf.FindPath(0, 1, 2, 1)
	if path != nil {
		t.Error("expected nil path when blocked")
	}
}

func TestPathfinder_SameCell(t *testing.T) {
	pf := NewPathfinder(5, 5)
	path := pf.FindPath(2, 2, 2, 2)
	if len(path) != 1 || path[0].X != 2 || path[0].Y != 2 {
		t.Errorf("same cell path: got %v", path)
	}
}

func TestPathfinder_Diagonals(t *testing.T) {
	pf := NewPathfinder(5, 5)
	pf.SetAllowDiagonals(true)
	path := pf.FindPath(0, 0, 4, 4)
	if path == nil {
		t.Fatal("expected diagonal path")
	}
	if path[0].X != 0 || path[0].Y != 0 {
		t.Errorf("path start: got (%d,%d)", path[0].X, path[0].Y)
	}
	if path[len(path)-1].X != 4 || path[len(path)-1].Y != 4 {
		t.Errorf("path end: got (%d,%d)", path[len(path)-1].X, path[len(path)-1].Y)
	}
	pf.SetAllowDiagonals(false)
	path4 := pf.FindPath(0, 0, 4, 4)
	if path4 == nil {
		t.Fatal("expected 4-dir path")
	}
	if len(path4) < len(path) {
		t.Errorf("4-dir path should not be shorter than 8-dir when open")
	}
}

func TestPathToWorld(t *testing.T) {
	path := []PathNode{{0, 0}, {1, 1}}
	world := PathToWorld(path, 32, 32)
	if len(world) != 2 {
		t.Fatalf("PathToWorld: got %d points", len(world))
	}
	if world[0].X() != 16 || world[0].Y() != 16 {
		t.Errorf("first point: got (%.0f,%.0f)", world[0].X(), world[0].Y())
	}
	exp := math2D.NewVector2D(48, 48)
	if world[1].X() != exp.X() || world[1].Y() != exp.Y() {
		t.Errorf("second point: got (%.0f,%.0f)", world[1].X(), world[1].Y())
	}
}
