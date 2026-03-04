package ebiten_extended

import "time"

// Timer is a utility for measuring specific durations and conditionally tracking loops.
type Timer struct {
	startTime time.Time
	duration time.Duration
	looped bool
}

// NewTimer initializes a new Timer with a target duration and a looping preference.
func NewTimer(duration time.Duration, isLooping bool) *Timer {
	return &Timer{duration: duration, looped: isLooping}
}

// SetDuration updates the target countdown duration of the timer.
func (t *Timer) SetDuration(duration time.Duration) *Timer {
	t.duration = duration
	return t
}

// GetDuration returns the current target duration.
func (t *Timer) GetDuration() time.Duration {
	return t.duration
}

// IsLooped returns true if the timer is configured to loop back after ending.
func (t *Timer) IsLooped() bool {
	return t.looped
}

// SetLooping adjusts the timer's looping behavior.
func (t *Timer) SetLooping(isLooping bool) *Timer {
	t.looped = isLooping
	return t
}

// Start begins the timer countdown from the current moment.
func (t *Timer) Start() *Timer {
	t.startTime = time.Now()
	return t
}

// IsEnded returns true if the timer duration has elapsed. It does not modify state.
// Call Restart() explicitly if you want to reset the timer.
func (t *Timer) IsEnded() bool {
	return t.GetElapsedTime() >= t.duration
}

// Restart resets the timer's start time and returns the time elapsed since the previous start.
func (t *Timer) Restart() time.Duration {
	elapsedTime := time.Since(t.startTime)
	t.startTime = time.Now()
	return elapsedTime
}

// GetElapsedTime returns the time passed since the timer was started.
func (t *Timer) GetElapsedTime() time.Duration {
	return time.Since(t.startTime)
}

// Update checks timer state and applies loop restart when enabled.
// Returns true when the timer has reached its duration in this update call.
func (t *Timer) Update() bool {
	if t.startTime.IsZero() {
		return false
	}
	if !t.IsEnded() {
		return false
	}
	if t.looped {
		t.Restart()
	}
	return true
}

