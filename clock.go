package ebiten_extended

import "time"

type Clock struct {
	startTime time.Time
	stopTime time.Time
}



func NewClock() *Clock {
	return &Clock{ startTime: time.Now() }
}

func (c *Clock) Start() {
	if (!c.IsRunning()) {
		c.startTime = time.Now()
		c.stopTime = time.Time{}
	}
}

func (c *Clock) Restart() time.Duration {
	elapsedTime := c.GetElapsedTime()
	c.startTime = time.Now()
	c.stopTime = time.Time{}
	return elapsedTime
}

func (c *Clock) Stop() {
	if (c.IsRunning()){
		c.stopTime = time.Now()
	}
}

func (c *Clock) GetElapsedTime() time.Duration {
	if c.IsRunning() {
		return	time.Since(c.startTime)
	}
	return c.stopTime.Sub(c.startTime)
}

func (c *Clock) IsRunning() bool {
	return c.stopTime.IsZero()
}
