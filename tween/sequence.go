package tween

import "github.com/LuigiVanacore/ludum"

type Sequence struct {
	ludum.Node
	Tweens []*Tween
	index  int
}

// NewSequence returns a new Sequence object.
func NewSequence(name string, tweens ...*Tween) *Sequence {
	seq := &Sequence{
		Node:   *ludum.NewNode(name),
		Tweens: tweens,
	}
	return seq
}

// Add adds one or more Tweens in order to the Sequence.
func (seq *Sequence) Add(tweens ...*Tween) {
	seq.Tweens = append(seq.Tweens, tweens...)
}

// Remove removes a Tween of the specified index from the Sequence.
func (seq *Sequence) Remove(index int) {
	if index < 0 || index >= len(seq.Tweens) {
		return
	}
	seq.Tweens = append(seq.Tweens[:index], seq.Tweens[index+1:]...)
	if seq.index > len(seq.Tweens) {
		seq.index = len(seq.Tweens)
	}
}

// Step advances the currently active Tween by one frame and returns the interpolated value,
// whether that Tween is complete, and whether the entire Sequence is complete.
func (seq *Sequence) Step() (float32, bool, bool) {
	return seq.stepImpl(func(t *Tween) (float32, bool) { return t.Step() })
}

// StepDelta advances the currently active Tween by delta seconds (frame-rate independent).
func (seq *Sequence) StepDelta(delta float64) (float32, bool, bool) {
	return seq.stepImpl(func(t *Tween) (float32, bool) { return t.StepDelta(delta) })
}

func (seq *Sequence) stepImpl(advance func(*Tween) (float32, bool)) (float32, bool, bool) {
	value := float32(0.0)
	tweenComplete := false
	sequenceComplete := false

	if seq.index < len(seq.Tweens) {
		value, tweenComplete = advance(seq.Tweens[seq.index])
		if tweenComplete {
			seq.Tweens[seq.index].Reset()
			seq.index++
			if seq.index >= len(seq.Tweens) {
				sequenceComplete = true
			}
		}
	} else {
		sequenceComplete = true
	}

	return value, tweenComplete, sequenceComplete
}

// Update advances the sequence by one fixed frame. Implements Updatable.
func (seq *Sequence) Update() {
	seq.StepDelta(ludum.FIXED_DELTA)
}

// Tick advances the sequence by one frame, satisfying the Updatable interface.
// Use Step() if you need the return values.
func (seq *Sequence) Tick() {
	seq.Step()
}

// Index returns the current index of the Sequence. Note that this can exceed the number of Tweens in the Sequence.
func (seq *Sequence) Index() int {
	return seq.index
}

// SetIndex sets the current index of the Sequence, influencing which Tween is active at any given time.
func (seq *Sequence) SetIndex(index int) {
	// Because it's possible to call SetIndex() when the Sequence is at the end.
	if seq.index >= 0 && seq.index < len(seq.Tweens) {
		seq.Tweens[seq.index].Reset()
	}
	seq.index = index
}

// Reset resets the Sequence, resetting all Tweens and setting the Sequence's index back to 0.
func (seq *Sequence) Reset() {
	for _, tween := range seq.Tweens {
		tween.Reset()
	}
	seq.index = 0
}

// HasTweens returns whether the Sequence is populated with Tweens or not.
func (seq *Sequence) HasTweens() bool {
	return len(seq.Tweens) > 0
}
