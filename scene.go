package ebiten_extended

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene defines the interface for a game scene (menu, level, etc.).
// Implement Enter to set up the scene (add nodes to World), Exit for cleanup,
// Update for per-frame logic, and Draw for custom rendering (or delegate to World).
type Scene interface {
	// Enter is called when the scene becomes active. Use to add nodes to engine.World().
	Enter(engine *Engine)
	// Exit is called when the scene is left. Use for cleanup (remove nodes, release resources).
	Exit()
	// Update runs each frame. Return an error to propagate to the game loop.
	Update() error
	// Draw renders the scene. Typically calls engine.World().Draw(screen) or custom rendering.
	Draw(screen *ebiten.Image)
}
