package ludum

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// pendingAction represents a queued scene change to perform after the fade-out transition.
type pendingAction int

const (
	pendingNone pendingAction = iota
	pendingPush
	pendingReplace
	pendingPop
)

// SceneManager manages a stack of scenes and delegates Engine Update/Draw to the current scene.
// When set on an Engine via SetSceneManager, the Engine delegates to the manager instead of the World directly.
// Use SetTransitionDuration to enable fade transitions between scenes.
type SceneManager struct {
	engine     *Engine
	scenes     []Scene
	current    Scene
	useManager bool // when true, Engine delegates Update/Draw to this manager

	// Transition support: when transitionDuration > 0, PushScene/ReplaceScene/PopScene use fade.
	transition         *Transition
	transitionDuration float32
	transitionColor    color.Color
	pending            pendingAction
	pendingScene       Scene
}

// NewSceneManager creates a SceneManager for the given engine.
func NewSceneManager(engine *Engine) *SceneManager {
	return &SceneManager{
		engine:             engine,
		scenes:             make([]Scene, 0),
		transitionColor:    color.Black,
		transitionDuration: 0,
	}
}

// SetTransitionDuration enables fade transitions for scene changes. Duration is in seconds (e.g. 0.3 for 300ms).
// When 0, scene changes are immediate (no transition).
func (m *SceneManager) SetTransitionDuration(duration float32) {
	m.transitionDuration = duration
	if duration > 0 && m.transition == nil {
		m.transition = NewTransition(duration)
	}
}

// SetTransitionColor sets the fade overlay color (default is black).
func (m *SceneManager) SetTransitionColor(c color.Color) {
	m.transitionColor = c
}

// PushScene pushes a new scene onto the stack. The previous scene's Exit is called, then the new scene's Enter.
// If SetTransitionDuration was called with a value > 0, a fade-out transition plays first, then the scene changes and fades in.
// When there is no current scene (first push), the scene is shown immediately without fade-out.
func (m *SceneManager) PushScene(s Scene) {
	if m.transitionDuration > 0 && m.transition != nil && m.current != nil {
		m.pending = pendingPush
		m.pendingScene = s
		m.transition.StartOut()
		return
	}
	m.doPushScene(s)
}

// PopScene removes the current scene and activates the previous one.
// Calls current.Exit(), then if a previous scene exists, calls previous.Enter().
// If transitions are enabled, a fade-out plays before the change.
func (m *SceneManager) PopScene() {
	if len(m.scenes) == 0 {
		return
	}
	if m.transitionDuration > 0 && m.transition != nil {
		m.pending = pendingPop
		m.pendingScene = nil
		m.transition.StartOut()
		return
	}
	m.doPopScene()
}

// ReplaceScene replaces the current scene with a new one (pop then push).
// If transitions are enabled and there is a current scene, a fade-out plays before the change.
func (m *SceneManager) ReplaceScene(s Scene) {
	if m.transitionDuration > 0 && m.transition != nil && m.current != nil {
		m.pending = pendingReplace
		m.pendingScene = s
		m.transition.StartOut()
		return
	}
	m.doReplaceScene(s)
}

// doPushScene performs an immediate push (used when no transition or after fade-out).
func (m *SceneManager) doPushScene(s Scene) {
	if m.current != nil {
		m.current.Exit()
	}
	m.scenes = append(m.scenes, s)
	m.current = s
	s.Enter(m.engine)
}

// doPopScene performs an immediate pop.
func (m *SceneManager) doPopScene() {
	if len(m.scenes) == 0 {
		return
	}
	if m.current != nil {
		m.current.Exit()
	}
	m.scenes = m.scenes[:len(m.scenes)-1]
	m.current = nil
	if len(m.scenes) > 0 {
		m.current = m.scenes[len(m.scenes)-1]
		m.current.Enter(m.engine)
	}
}

// doReplaceScene performs an immediate replace.
func (m *SceneManager) doReplaceScene(s Scene) {
	if m.current != nil {
		m.current.Exit()
	}
	if len(m.scenes) > 0 {
		m.scenes[len(m.scenes)-1] = s
	} else {
		m.scenes = append(m.scenes, s)
	}
	m.current = s
	s.Enter(m.engine)
}

// CurrentScene returns the currently active scene, or nil if none.
func (m *SceneManager) CurrentScene() Scene {
	return m.current
}

// Update runs the current scene's Update and the world's Update.
// When a transition is in progress, it advances the transition; on fade-out complete, performs the pending scene change and starts fade-in.
func (m *SceneManager) Update() error {
	delta := m.engine.ScaledDelta()

	// Handle transition: when fading out, don't run scene/world Update; when fading in, do run them.
	if m.pending != pendingNone && m.transition != nil {
		done := m.transition.Update(delta)
		if done && m.transition.State == TransitionDone {
			// Fade-out finished: perform the pending scene change, then start fade-in.
			switch m.pending {
			case pendingPush:
				m.doPushScene(m.pendingScene)
			case pendingReplace:
				m.doReplaceScene(m.pendingScene)
			case pendingPop:
				m.doPopScene()
			}
			m.pending = pendingNone
			m.pendingScene = nil
			m.transition.StartIn()
		}
		// While transitioning out, skip scene/world update so the old scene stays frozen
		if m.pending != pendingNone {
			return nil
		}
		// Fade-in in progress: fall through to run scene/world update
	}

	if m.current != nil {
		if err := m.current.Update(); err != nil {
			return err
		}
	}
	if !m.engine.paused {
		m.engine.world.Update()
	}

	// Advance fade-in transition; when complete, reset to idle
	if m.transition != nil && m.transition.State == TransitionIn {
		if m.transition.Update(delta) {
			m.transition.State = TransitionIdle
		}
	}

	return nil
}

// Draw renders the current scene. Typically the scene draws via the World.
// When a transition is active, draws the fade overlay on top.
func (m *SceneManager) Draw(target *ebiten.Image) {
	if m.current != nil {
		m.current.Draw(target)
	} else {
		m.engine.world.Draw(target, nil)
	}
	if m.transition != nil && (m.transition.State == TransitionOut || m.transition.State == TransitionIn) {
		m.transition.DrawFadeOverlay(target, m.transitionColor)
	}
}
