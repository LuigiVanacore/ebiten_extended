package ebiten_extended

import (
	"github.com/LuigiVanacore/ebiten_extended/input"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FIXED_DELTA float64 = 1.0 / 60.0
)

type Engine struct {
	world           *World
	inputManager    *input.InputManager
	resourceManager *ResourceManager
	clock           *Clock
	debug           *Debug
}

func NewEngine() *Engine {
	e := &Engine{
		world:           NewWorld(),
		inputManager:    input.NewInputManager(),
		resourceManager: NewResourceManager(),
		clock:           NewClock(),
		debug:           NewDebug(false),
	}

	return e
}

func (e *Engine) World() *World {
	return e.world
}

func (e *Engine) Input() *input.InputManager {
	return e.inputManager
}

func (e *Engine) Resources() *ResourceManager {
	return e.resourceManager
}

func (e *Engine) IsDebug() bool {
	return e.debug.Enabled()
}

func (e *Engine) SetIsDebug(debugFlag bool) {
	e.debug.SetEnabled(debugFlag)
}

func (e *Engine) Update() error {
	e.inputManager.Update()
	e.world.Update()
	return nil
}

// Ensure ebitengine signature compatibility
func (e *Engine) Draw(target *ebiten.Image) {
	// The world might need an ops to start with
	op := &ebiten.DrawImageOptions{}
	e.world.Draw(target, op)
}

func (e *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	// default placeholder for Layout, actual game overrides this.
	return outsideWidth, outsideHeight
}
