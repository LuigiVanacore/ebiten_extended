package ebiten_extended

import "time"

type Timer struct {
	startTime time.Time
	duration time.Duration
	looped bool
}

func NewTimer(duration time.Duration, isLooping bool) *Timer {
	return &Timer{duration: duration, looped: isLooping}
}

func (t *Timer) SetDuration(duration time.Duration) *Timer {
	t.duration = duration
	return t
}

func (t *Timer) GetDuration() time.Duration {
	return t.duration
}

func (t *Timer) IsLooped() bool {
	return t.looped
}

func (t *Timer) SetLooping(isLooping bool) *Timer {
	t.looped = isLooping
	return t
}

func (t *Timer) Start() *Timer {
	t.startTime = time.Now()
	return t
}

// IsEnded returns true if the timer duration has elapsed. It does not modify state.
// Call Restart() explicitly if you want to reset the timer.
func (t *Timer) IsEnded() bool {
	return t.GetElapsedTime() >= t.duration
}

func (t *Timer) Restart() time.Duration {
	elapsedTime := time.Since(t.startTime)
	t.startTime = time.Now()
	return elapsedTime
}

func (t *Timer) GetElapsedTime() time.Duration {
	return time.Since(t.startTime)
}

