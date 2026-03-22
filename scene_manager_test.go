package ludum

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

// mockScene is a test Scene that tracks Enter/Exit calls.
type mockScene struct {
	engine     *Engine
	enterCalls int
	exitCalls  int
}

func newMockScene() *mockScene {
	return &mockScene{}
}

func (s *mockScene) Enter(engine *Engine) {
	s.engine = engine
	s.enterCalls++
}

func (s *mockScene) Exit() {
	s.exitCalls++
}

func (s *mockScene) Update() error {
	return nil
}

func (s *mockScene) Draw(screen *ebiten.Image) {}

func TestNewSceneManager(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	if sm == nil {
		t.Fatal("NewSceneManager returned nil")
	}
	if sm.CurrentScene() != nil {
		t.Error("new SceneManager should have no current scene")
	}
}

func TestSceneManagerPushSceneImmediate(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	s1 := newMockScene()

	sm.PushScene(s1)
	if sm.CurrentScene() != s1 {
		t.Error("CurrentScene should be s1 after PushScene")
	}
	if s1.enterCalls != 1 {
		t.Errorf("Enter should be called once, got %d", s1.enterCalls)
	}
}

func TestSceneManagerReplaceSceneImmediate(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	s1 := newMockScene()
	s2 := newMockScene()

	sm.PushScene(s1)
	sm.ReplaceScene(s2)
	if sm.CurrentScene() != s2 {
		t.Error("CurrentScene should be s2 after ReplaceScene")
	}
	if s1.exitCalls != 1 {
		t.Errorf("s1.Exit should be called, got %d", s1.exitCalls)
	}
	if s2.enterCalls != 1 {
		t.Errorf("s2.Enter should be called, got %d", s2.enterCalls)
	}
}

func TestSceneManagerPopSceneImmediate(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	s1 := newMockScene()
	s2 := newMockScene()

	sm.PushScene(s1)
	sm.PushScene(s2)
	sm.PopScene()
	if sm.CurrentScene() != s1 {
		t.Error("CurrentScene should be s1 after PopScene")
	}
	if s2.exitCalls != 1 {
		t.Errorf("s2.Exit should be called, got %d", s2.exitCalls)
	}
	if s1.enterCalls != 2 {
		t.Errorf("s1.Enter should be called again on pop, got %d", s1.enterCalls)
	}
}

func TestSceneManagerPopEmptyNoPanic(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	sm.PopScene() // should not panic
}

func TestSceneManagerSetTransitionDuration(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)

	sm.SetTransitionDuration(0.3)
	if sm.transitionDuration != 0.3 {
		t.Errorf("transitionDuration: got %v, want 0.3", sm.transitionDuration)
	}
	if sm.transition == nil {
		t.Error("transition should be created when duration > 0")
	}

	sm.SetTransitionDuration(0)
	if sm.transitionDuration != 0 {
		t.Errorf("transitionDuration: got %v, want 0", sm.transitionDuration)
	}
}

func TestSceneManagerSetTransitionColor(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)

	red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	sm.SetTransitionColor(red)
	if sm.transitionColor != red {
		t.Error("SetTransitionColor did not set color")
	}
}

func TestSceneManagerWithTransitionPushScene(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	sm.SetTransitionDuration(0.5)
	s1 := newMockScene()
	s2 := newMockScene()

	sm.PushScene(s1)
	if sm.CurrentScene() != s1 {
		t.Error("first PushScene should be immediate (no prior scene to fade from)")
	}

	// Push s2 with transition: should start fade-out, not change scene yet
	sm.PushScene(s2)
	if sm.CurrentScene() != s1 {
		t.Error("during fade-out, current should still be s1")
	}
	if sm.pending != pendingPush {
		t.Error("pending should be pendingPush")
	}
	if sm.pendingScene != s2 {
		t.Error("pendingScene should be s2")
	}

	// Simulate enough Update frames for fade-out to complete
	for i := 0; i < 30 && sm.pending != pendingNone; i++ {
		_ = sm.Update()
	}
	// After fade-out completes, we do swap and start fade-in. Run a few more.
	for i := 0; i < 35; i++ {
		_ = sm.Update()
	}

	if sm.CurrentScene() != s2 {
		t.Errorf("after transition: CurrentScene should be s2, got %T", sm.CurrentScene())
	}
	if s1.exitCalls != 1 {
		t.Errorf("s1.Exit: got %d, want 1", s1.exitCalls)
	}
	if s2.enterCalls != 1 {
		t.Errorf("s2.Enter: got %d, want 1", s2.enterCalls)
	}
}

func TestSceneManagerDrawWithTransition(t *testing.T) {
	engine := NewEngine()
	sm := NewSceneManager(engine)
	sm.PushScene(newMockScene())
	sm.SetTransitionDuration(0.3)
	sm.PushScene(newMockScene())

	target := ebiten.NewImage(100, 100)
	defer target.Deallocate()
	sm.Draw(target)
	// Should not panic; Draw renders current scene + overlay during transition
}
