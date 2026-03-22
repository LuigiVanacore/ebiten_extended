package utils

import (
	"testing"
)

func TestByteSet_Set(t *testing.T) {
	var a ByteSet = 0b0001
	var b ByteSet = 0b0010
	result := a.Set(b)
	expected := ByteSet(0b0011)
	if result != expected {
		t.Errorf("Set failed: expected %b, got %b", expected, result)
	}
}

func TestByteSet_Clear(t *testing.T) {
	var a ByteSet = 0b0111
	var b ByteSet = 0b0010
	result := a.Clear(b)
	expected := ByteSet(0b0101)
	if result != expected {
		t.Errorf("Clear failed: expected %b, got %b", expected, result)
	}
}

func TestByteSet_Has(t *testing.T) {
	var a ByteSet = 0b1010
	tests := []struct {
		flag     ByteSet
		expected bool
	}{
		{0b0010, true},
		{0b1000, true},
		{0b0001, false},
		{0b0100, false},
	}
	for _, tt := range tests {
		if got := a.Has(tt.flag); got != tt.expected {
			t.Errorf("Has(%b): expected %v, got %v", tt.flag, tt.expected, got)
		}
	}
}
