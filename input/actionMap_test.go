package input

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestActionMapSetKeyBinding(t *testing.T) {
	am := NewActionMap()
	am.AddAction(1, NewKeyAction(ebiten.KeySpace, ActionHold))

	key, ok := am.GetKeyForAction(1)
	if !ok || key != ebiten.KeySpace {
		t.Errorf("GetKeyForAction: got key %v ok=%v, want Space", key, ok)
	}

	err := am.SetKeyBinding(1, ebiten.KeyEnter)
	if err != nil {
		t.Fatalf("SetKeyBinding: %v", err)
	}

	key, ok = am.GetKeyForAction(1)
	if !ok || key != ebiten.KeyEnter {
		t.Errorf("After SetKeyBinding: got key %v ok=%v, want Enter", key, ok)
	}
}

func TestActionMapSetKeyBindingNotFound(t *testing.T) {
	am := NewActionMap()
	err := am.SetKeyBinding(99, ebiten.KeySpace)
	if err == nil {
		t.Error("SetKeyBinding for non-existent action should return error")
	}
	if err != ErrActionNotFound {
		t.Errorf("expected ErrActionNotFound, got %v", err)
	}
}

func TestActionMapGetKeyForActionNotKeyAction(t *testing.T) {
	am := NewActionMap()
	am.AddAction(1, NewMouseButtonAction(ebiten.MouseButtonLeft, ActionHold))

	_, ok := am.GetKeyForAction(1)
	if ok {
		t.Error("GetKeyForAction for MouseButtonAction should return false")
	}
}
