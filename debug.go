package ebiten_extended



type Debug struct {
	enabled bool
}


func NewDebug(enabled bool) *Debug {
	return &Debug{enabled: enabled}
}

func (d *Debug) Enabled() bool {
	return d.enabled
}

func (d *Debug) SetEnabled(enabled bool) {
	d.enabled = enabled
}