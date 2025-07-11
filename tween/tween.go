package tween

import "github.com/LuigiVanacore/ebiten_extended"

type Tween struct {
	ebiten_extended.Node
	duration float32
	time     float32
	begin    float32
	end      float32
	change   float32
	easing   TweenFunc
}

func NewTween(name string, begin, end, duration float32, easing TweenFunc) *Tween {
	return &Tween{
		Node:     *ebiten_extended.NewNode(name),
		begin:    begin,
		end:      end,
		change:   end - begin,
		duration: duration,
		easing:   easing,
	}
}
 
func (tween *Tween) Set(time float32) (current float32, isFinished bool) {
	if time <= 0 {
		tween.time = 0
		current = tween.begin
	} else if time >= tween.duration {
		tween.time = tween.duration
		current = tween.end
	} else {
		tween.time = time
		current = tween.easing(tween.time, tween.begin, tween.change, tween.duration)
	}

	return current, tween.time >= tween.duration
}
 
func (tween *Tween) Reset() {
	tween.Set(0)
}
 
func (tween *Tween) Update() (current float32, isFinished bool) {
	return tween.Set(tween.time)
}