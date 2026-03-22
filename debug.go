package ludum

// Debug represents a global indicator for logic tracking and diagnostic feedback.
type Debug struct {
	enabled bool
}

// NewDebug instantiates a new Debug state representation initialized with the specified flag.
func NewDebug(enabled bool) *Debug {
	return &Debug{enabled: enabled}
}

// Enabled returns whether the system is presently flagged for debugging.
func (d *Debug) Enabled() bool {
	return d.enabled
}

// SetEnabled adjusts the diagnostic tracking mode.
func (d *Debug) SetEnabled(enabled bool) {
	d.enabled = enabled
}
