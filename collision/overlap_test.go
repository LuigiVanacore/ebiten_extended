package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestOverlapCircle_Empty(t *testing.T) {
	mgr := NewCollisionManager()
	results := mgr.OverlapCircle(math2D.NewVector2D(50, 50), 20)
	if len(results) != 0 {
		t.Errorf("empty manager: got %d overlaps", len(results))
	}
}

func TestOverlapCircle_HitsCircle(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	circ := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 15))
	col, _ := NewCollider("c", circ, mask)
	col.SetPosition(50, 50)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	results := mgr.OverlapCircle(math2D.NewVector2D(55, 55), 10)
	if len(results) != 1 {
		t.Errorf("circle overlap: got %d overlaps, want 1", len(results))
	}
}

func TestOverlapCircle_NoHit(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	rect := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	col, _ := NewCollider("r", rect, mask)
	col.SetPosition(100, 100)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	results := mgr.OverlapCircle(math2D.NewVector2D(0, 0), 5)
	if len(results) != 0 {
		t.Errorf("no overlap: got %d overlaps", len(results))
	}
}

func TestOverlapRect_Empty(t *testing.T) {
	mgr := NewCollisionManager()
	rect := math2D.NewRectangle(math2D.NewVector2D(0, 0), math2D.NewVector2D(100, 100))
	results := mgr.OverlapRect(rect)
	if len(results) != 0 {
		t.Errorf("empty manager: got %d overlaps", len(results))
	}
}

func TestOverlapRect_HitsRect(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	rect := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	col, _ := NewCollider("r", rect, mask)
	col.SetPosition(50, 50)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	query := math2D.NewRectangle(math2D.NewVector2D(40, 40), math2D.NewVector2D(30, 30))
	results := mgr.OverlapRect(query)
	if len(results) != 1 {
		t.Errorf("rect overlap: got %d overlaps, want 1", len(results))
	}
}

func TestOverlapRect_HitsCircle(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	circ := NewCollisionCircle(math2D.NewCircle(math2D.ZeroVector2D(), 15))
	col, _ := NewCollider("c", circ, mask)
	col.SetPosition(50, 50)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	query := math2D.NewRectangle(math2D.NewVector2D(45, 45), math2D.NewVector2D(20, 20))
	results := mgr.OverlapRect(query)
	if len(results) != 1 {
		t.Errorf("rect vs circle: got %d overlaps, want 1", len(results))
	}
}
