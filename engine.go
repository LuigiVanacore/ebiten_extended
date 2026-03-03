package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
)

// FIXED_DELTA is the fixed time step used for engine updates, set to 60 FPS (1.0 / 60.0).
const (
	FIXED_DELTA float64 = 1.0 / 60.0
)

// Engine represents the core of the framework, managing the game world, input, resources, audio, clock, and debug systems.
type Engine struct {
	world           *World
	inputManager    *input.InputManager
	resourceManager *ResourceManager
	audioManager    *AudioManager
	clock           *Clock
	debug           *Debug
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
	// The world might need an ops to start with
	op := &ebiten.DrawImageOptions{}
	e.world.Draw(target, op)
}

// Layout accepts a screen size and returns the logical screen size.
// Can be overridden by the actual game to adjust screen proportions.
func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// default placeholder for Layout, actual game overrides this.
	return outsideWidth, outsideHeight
}
