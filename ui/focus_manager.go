package ui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Focusable represents a UI component that can receive keyboard/gamepad focus.
type Focusable interface {
	SetFocused(focused bool)
	IsFocused() bool
	OnFocusAction()
}

// FocusManager handles keyboard and gamepad navigation between interactive UI elements.
type FocusManager struct {
	items        []Focusable
	focusedIndex int
}

// NewFocusManager creates a new FocusManager instance.
func NewFocusManager() *FocusManager {
	return &FocusManager{
		items:        make([]Focusable, 0),
		focusedIndex: -1,
	}
}

// Register adds a Focusable item to the manager's list.
func (fm *FocusManager) Register(f Focusable) {
	fm.items = append(fm.items, f)
	// If it's the first item added, automatically focus it
	if fm.focusedIndex == -1 {
		fm.SetFocusIndex(0)
	}
}

// Unregister removes a Focusable item from the manager's list.
func (fm *FocusManager) Unregister(f Focusable) {
	idx := -1
	for i, item := range fm.items {
		if item == f {
			idx = i
			break
		}
	}
	if idx != -1 {
		fm.items = append(fm.items[:idx], fm.items[idx+1:]...)
		if fm.focusedIndex == idx {
			fm.focusedIndex = -1
			if len(fm.items) > 0 {
				fm.SetFocusIndex(0)
			}
		} else if fm.focusedIndex > idx {
			fm.focusedIndex--
		}
	}
}

// SetFocusIndex forces focus to a specific index in the list.
func (fm *FocusManager) SetFocusIndex(index int) {
	if fm.focusedIndex >= 0 && fm.focusedIndex < len(fm.items) {
		fm.items[fm.focusedIndex].SetFocused(false)
	}
	fm.focusedIndex = index
	if fm.focusedIndex >= 0 && fm.focusedIndex < len(fm.items) {
		fm.items[fm.focusedIndex].SetFocused(true)
	}
}

// Update processes keyboard/gamepad inputs for navigation.
// It cycles through the registered Focusable elements using Up/Down/Left/Right or Tab,
// and triggers the action via Enter or Space.
func (fm *FocusManager) Update() {
	if len(fm.items) == 0 {
		return
	}

	// Navigation Forward (Down / Right / Tab)
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		nextIndex := (fm.focusedIndex + 1) % len(fm.items)
		fm.SetFocusIndex(nextIndex)
	}

	// Navigation Backward (Up / Left)
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		nextIndex := fm.focusedIndex - 1
		if nextIndex < 0 {
			nextIndex = len(fm.items) - 1
		}
		fm.SetFocusIndex(nextIndex)
	}

	// Action (Enter / Space)
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if fm.focusedIndex >= 0 && fm.focusedIndex < len(fm.items) {
			fm.items[fm.focusedIndex].OnFocusAction()
		}
	}
}
