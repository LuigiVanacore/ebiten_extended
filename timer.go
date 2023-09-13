package ebiten_extended

import "time"

type Timer struct {
	startTime time.Time
	duration time.Duration
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{duration: duration}
}

func (t *Timer) SetDuration(duration time.Duration) {
	t.duration = duration
}

func (t *Timer) GetDuration() time.Duration {
	return t.duration
}

func (t *Timer) Start() *Timer {
	t.startTime = time.Now()
	return t
}

func (t *Timer) IsEnded() bool {
	if t.GetElapsedTime() >= t.duration {
		t.Restart()
		return true
	}
	return false
}

func (t *Timer) Restart() time.Duration {
	elapsedTime := time.Since(t.startTime)
	t.startTime = time.Now()
	return elapsedTime
}

func (t *Timer) GetElapsedTime() time.Duration {
	return time.Since(t.startTime)
}

