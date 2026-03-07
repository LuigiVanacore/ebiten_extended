package collision

import (
	"testing"

	"github.com/LuigiVanacore/ebiten_extended/math2D"
	"github.com/LuigiVanacore/ebiten_extended/utils"
)

func TestRaycast_Empty(t *testing.T) {
	mgr := NewCollisionManager()
	results := mgr.Raycast(math2D.NewVector2D(0, 0), math2D.NewVector2D(100, 100))
	if len(results) != 0 {
		t.Errorf("empty manager: got %d results", len(results))
	}
}

func TestRaycast_HitsRect(t *testing.T) {
	mask := NewCollisionMask(utils.ByteSet(1), utils.ByteSet(1))
	rect := NewCollisionRect(math2D.NewRectangle(math2D.ZeroVector2D(), math2D.NewVector2D(20, 20)))
	col, _ := NewCollider("box", rect, mask)
	col.SetPosition(50, 50)
	mgr := NewCollisionManager()
	mgr.AddParticipant(col)
	results := mgr.Raycast(math2D.NewVector2D(0, 50), math2D.NewVector2D(100, 50))
	if len(results) != 1 {
		t.Fatalf("expected 1 hit, got %d", len(results))
	}
	if results[0].Participant != col {
		t.Error("wrong participant")
	}
}
