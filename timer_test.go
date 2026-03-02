package ebiten_extended

import (
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	duration := 2 * time.Second
	isLooping := true
	timer := NewTimer(duration, isLooping)

	if timer.GetDuration() != duration {
		t.Errorf("expected duration %v, got %v", duration, timer.GetDuration())
	}

	if timer.IsLooped() != isLooping {
		t.Errorf("expected looping %v, got %v", isLooping, timer.IsLooped())
	}
}

func TestSetDuration(t *testing.T) {
	timer := NewTimer(1*time.Second, false)
	newDuration := 3 * time.Second
	timer.SetDuration(newDuration)

	if timer.GetDuration() != newDuration {
		t.Errorf("expected duration %v, got %v", newDuration, timer.GetDuration())
	}
}

func TestSetLooping(t *testing.T) {
	timer := NewTimer(1*time.Second, false)
	timer.SetLooping(true)

	if !timer.IsLooped() {
		t.Errorf("expected looping to be true, got false")
	}
}

func TestStartAndElapsedTime(t *testing.T) {
	timer := NewTimer(1*time.Second, false)
	timer.Start()
	time.Sleep(100 * time.Millisecond)

	elapsed := timer.GetElapsedTime()
	if elapsed < 100*time.Millisecond || elapsed > 200*time.Millisecond {
		t.Errorf("expected elapsed time around 100ms, got %v", elapsed)
	}
}

func TestIsEnded(t *testing.T) {
	timer := NewTimer(100*time.Millisecond, false)
	timer.Start()
	time.Sleep(150 * time.Millisecond)

	if !timer.IsEnded() {
		t.Errorf("expected timer to have ended, but it did not")
	}
}

func TestRestart(t *testing.T) {
	timer := NewTimer(1*time.Second, false)
	timer.Start()
	time.Sleep(100 * time.Millisecond)

	elapsed := timer.Restart()
	if elapsed < 100*time.Millisecond || elapsed > 200*time.Millisecond {
		t.Errorf("expected elapsed time around 100ms, got %v", elapsed)
	}

	if timer.GetElapsedTime() > 10*time.Millisecond {
		t.Errorf("expected elapsed time after restart to be close to 0, got %v", timer.GetElapsedTime())
	}
}