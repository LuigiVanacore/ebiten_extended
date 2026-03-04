package tween

import (
	"testing"
)

func TestNewSequence(t *testing.T) {
	tw1 := NewTween("a", 0, 1, 1.0, Linear)
	tw2 := NewTween("b", 0, 1, 1.0, Linear)
	seq := NewSequence("seq", tw1, tw2)
	if seq == nil {
		t.Fatal("NewSequence should not return nil")
	}
	if len(seq.Tweens) != 2 || seq.index != 0 {
		t.Errorf("NewSequence: len(Tweens)=%v index=%v", len(seq.Tweens), seq.index)
	}
}

func TestNewSequence_Empty(t *testing.T) {
	seq := NewSequence("empty")
	if len(seq.Tweens) != 0 {
		t.Errorf("NewSequence() empty: len(Tweens)=%v", len(seq.Tweens))
	}
}

func TestSequence_Add(t *testing.T) {
	seq := NewSequence("seq")
	tw := NewTween("t", 0, 1, 1.0, Linear)
	seq.Add(tw)
	if len(seq.Tweens) != 1 {
		t.Errorf("Add: len(Tweens)=%v", len(seq.Tweens))
	}
	seq.Add(NewTween("t2", 0, 1, 1.0, Linear))
	if len(seq.Tweens) != 2 {
		t.Errorf("Add twice: len(Tweens)=%v", len(seq.Tweens))
	}
}

func TestSequence_Remove(t *testing.T) {
	tw1 := NewTween("a", 0, 1, 1.0, Linear)
	tw2 := NewTween("b", 0, 1, 1.0, Linear)
	tw3 := NewTween("c", 0, 1, 1.0, Linear)
	seq := NewSequence("seq", tw1, tw2, tw3)
	seq.Remove(1)
	if len(seq.Tweens) != 2 {
		t.Errorf("Remove(1): len(Tweens)=%v", len(seq.Tweens))
	}
	if seq.Tweens[0] != tw1 || seq.Tweens[1] != tw3 {
		t.Error("Remove(1) should remove middle tween")
	}
}

func TestSequence_HasTweens(t *testing.T) {
	empty := NewSequence("empty")
	if empty.HasTweens() {
		t.Error("Empty sequence should not have tweens")
	}
	seq := NewSequence("seq", NewTween("t", 0, 1, 1.0, Linear))
	if !seq.HasTweens() {
		t.Error("Sequence with tween should have tweens")
	}
}

func TestSequence_Index_SetIndex(t *testing.T) {
	seq := NewSequence("seq",
		NewTween("a", 0, 1, 1.0, Linear),
		NewTween("b", 0, 1, 1.0, Linear),
	)
	if seq.Index() != 0 {
		t.Errorf("Index() = %v, want 0", seq.Index())
	}
	seq.SetIndex(1)
	if seq.Index() != 1 {
		t.Errorf("SetIndex(1): Index() = %v", seq.Index())
	}
}

func TestSequence_Reset(t *testing.T) {
	tw := NewTween("t", 0, 100, 1.0, Linear)
	seq := NewSequence("seq", tw)
	tw.Set(0.5)
	seq.SetIndex(0)
	seq.Reset()
	if seq.Index() != 0 {
		t.Errorf("Reset: Index = %v, want 0", seq.Index())
	}
	current, _ := tw.Set(0)
	if current != 0 {
		t.Errorf("Reset: tween value = %v, want 0", current)
	}
}

func TestSequence_Update(t *testing.T) {
	tw := NewTween("t", 0, 100, 2.0, Linear)
	seq := NewSequence("seq", tw)
	// First Update advances one frame.
	value, tweenComplete, seqComplete := seq.Update()
	if value != 50 {
		t.Errorf("Update at start: value = %v, want 50", value)
	}
	if tweenComplete || seqComplete {
		t.Errorf("Update at start: tweenComplete=%v seqComplete=%v", tweenComplete, seqComplete)
	}
}

func TestSequence_Remove_OutOfBoundsNoPanic(t *testing.T) {
	seq := NewSequence("seq", NewTween("a", 0, 1, 1.0, Linear))
	seq.Remove(-1)
	seq.Remove(5)
	if len(seq.Tweens) != 1 {
		t.Fatalf("out-of-bounds remove should not modify sequence, got %d tweens", len(seq.Tweens))
	}
}

func TestSequence_Update_AdvancesWhenTweenComplete(t *testing.T) {
	tw1 := NewTween("t1", 0, 10, 0.001, Linear)
	tw2 := NewTween("t2", 100, 200, 0.001, Linear)
	seq := NewSequence("seq", tw1, tw2)
	// Advance tween1 to completion
	tw1.Set(0.001)
	value, tweenComplete, seqComplete := seq.Update()
	if value != 10 {
		t.Errorf("Update when t1 done: value = %v, want 10", value)
	}
	if !tweenComplete {
		t.Error("tweenComplete should be true when tween finishes")
	}
	if seqComplete {
		t.Error("seqComplete should be false (tween2 still pending)")
	}
	if seq.Index() != 1 {
		t.Errorf("Index should be 1 after first tween done, got %v", seq.Index())
	}
}

func TestSequence_Update_SequenceComplete(t *testing.T) {
	tw := NewTween("t", 0, 10, 0.001, Linear)
	seq := NewSequence("seq", tw)
	tw.Set(0.001) // complete the only tween
	_, _, seqComplete := seq.Update()
	if !seqComplete {
		t.Error("seqComplete should be true when last tween finishes")
	}
}

func TestSequence_Update_EmptySequence(t *testing.T) {
	seq := NewSequence("empty")
	_, _, seqComplete := seq.Update()
	if !seqComplete {
		t.Error("Empty sequence Update should return sequenceComplete=true")
	}
}

func TestSequence_SetIndex_ResetsCurrentTween(t *testing.T) {
	tw1 := NewTween("t1", 0, 10, 1.0, Linear)
	tw2 := NewTween("t2", 0, 10, 1.0, Linear)
	seq := NewSequence("seq", tw1, tw2)
	tw1.Set(0.5)
	seq.SetIndex(1)
	seq.SetIndex(0)
	current, _ := tw1.Set(0)
	if current != 0 {
		t.Errorf("SetIndex back to 0 should reset tween1, got %v", current)
	}
}
