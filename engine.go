package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
)

// FIXED_DELTA is the fixed time step used for engine updates, set to 60 FPS (1.0 / 60.0).
// For physics simulation, prefer PhysicsDelta() which matches Ebiten's TPS.
const (
	FIXED_DELTA float64 = 1.0 / 60.0
)

// PhysicsDelta returns the fixed timestep for physics, matching Ebiten's TPS.
// Use this in physicsWorld.Step(ebiten_extended.PhysicsDelta()) for correct
// integration with Ebiten's Update frequency (ebiten.SetTPS affects this).
func PhysicsDelta() float64 {
	return 1.0 / float64(ebiten.TPS())
}

// Engine represents the core of the framework, managing the game world, input, resources, audio, clock, and debug systems.
type Engine struct {
	world           *World
	inputManager    *input.InputManager
	resourceManager *ResourceManager
	audioManager    *AudioManager
	clock           *Clock
	debug           *Debug

	// logicalSize is used for Layout when SetLogicalSize was called; (0,0) = passthrough.
	logicalWidth  int
	logicalHeight int
}

// NewEngine creates and initializes a new Engine instance with default systems.
func NewEngine() *Engine {
	e := &Engine{
		world:           NewWorld(),
		inputManager:    input.NewInputManager(),
		resourceManager: NewResourceManager(),
		audioManager:    NewAudioManager(),
		clock:           NewClock(),
		debug:           NewDebug(false),
	}

	return e
}

// World returns the game world associated with the engine.
func (e *Engine) World() *World {
	return e.world
}

// Input returns the input manager responsible for handling user inputs.
func (e *Engine) Input() *input.InputManager {
	return e.inputManager
}

// Resources returns the resource manager handling game assets.
func (e *Engine) Resources() *ResourceManager {
	return e.resourceManager
}

// Audio returns the audio manager for sounds and music.
func (e *Engine) Audio() *AudioManager {
	return e.audioManager
}

// Clock returns the engine clock.
func (e *Engine) Clock() *Clock {
	return e.clock
}

// IsDebug returns whether debug mode is currently enabled.
func (e *Engine) IsDebug() bool {
	return e.debug.Enabled()
}

// SetIsDebug enables or disables debug mode.
func (e *Engine) SetIsDebug(debugFlag bool) {
	e.debug.SetEnabled(debugFlag)
}

// Update advances the engine state by one tick, updating input and the game world.
func (e *Engine) Update() error {
	e.inputManager.Update()
	e.world.Update()
	return nil
}

// Draw renders the game world onto the target screen image.
// Ensure ebitengine signature compatibility.
func (e *Engine) Draw(target *ebiten.Image) {
	e.world.Draw(target, nil)
}

// SetLogicalSize sets the fixed logical resolution for scaling.
// When non-zero, Layout returns (logicalW, logicalH) so the game renders at a fixed resolution
// while Ebiten scales to the window. Use in your Game.Layout: return e.engine.Layout(outsideWidth, outsideHeight).
func (e *Engine) SetLogicalSize(logicalWidth, logicalHeight int) {
	e.logicalWidth = logicalWidth
	e.logicalHeight = logicalHeight
}

// Layout accepts a screen size and returns the logical screen size.
// If SetLogicalSize was called with non-zero values, returns those; otherwise passthrough.
// Resizes the world camera to match the returned dimensions.
func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	var w, h int
	if e.logicalWidth > 0 && e.logicalHeight > 0 {
		w, h = e.logicalWidth, e.logicalHeight
	} else {
		w, h = outsideWidth, outsideHeight
	}
	e.world.Camera().Resize(uint(w), uint(h))
	return w, h
}
