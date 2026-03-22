package ludum

import "time"

// Clock represents a simple time-tracking mechanism to measure elapsed durations.
type Clock struct {
	startTime time.Time
	stopTime  time.Time
}

// NewClock creates a new Clock instance with its start time initialized to the current moment.
func NewClock() *Clock {
	return &Clock{startTime: time.Now()}
}

// Start begins the clock's time tracking if it is not already running.
func (c *Clock) Start() {
	if !c.IsRunning() {
		c.startTime = time.Now()
		c.stopTime = time.Time{}
	}
}

// Restart resets the clock's start time to now and returns the duration elapsed since the last start.
func (c *Clock) Restart() time.Duration {
	elapsedTime := c.GetElapsedTime()
	c.startTime = time.Now()
	c.stopTime = time.Time{}
	return elapsedTime
}

// Stop halts the clock and records the current time as the stop time.
func (c *Clock) Stop() {
	if c.IsRunning() {
		c.stopTime = time.Now()
	}
}

// GetElapsedTime calculates the duration since the clock was started, or the total tracked time if stopped.
func (c *Clock) GetElapsedTime() time.Duration {
	if c.IsRunning() {
		return time.Since(c.startTime)
	}
	return c.stopTime.Sub(c.startTime)
}

// IsRunning returns true if the clock is currently actively tracking time without being stopped.
func (c *Clock) IsRunning() bool {
	return c.stopTime.IsZero()
}
