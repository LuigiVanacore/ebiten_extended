package utils

import (
	"testing"
)

func TestStackPush_Int(t *testing.T) {
	stack := NewStack[int]()
	if stack.Size() != 0 {
		t.Errorf("Expected size 0, got %d", stack.Size())
	}

	stack.Push(10)
	if stack.Size() != 1 {
		t.Errorf("Expected size 1 after push, got %d", stack.Size())
	}
	top, ok := stack.Peek()
	if !ok || top != 10 {
		t.Errorf("Expected top 10, got %v, ok=%v", top, ok)
	}

	stack.Push(20)
	if stack.Size() != 2 {
		t.Errorf("Expected size 2 after second push, got %d", stack.Size())
	}
	top, ok = stack.Peek()
	if !ok || top != 20 {
		t.Errorf("Expected top 20, got %v, ok=%v", top, ok)
	}
}

func TestStackPush_String(t *testing.T) {
	stack := NewStack[string]()
	stack.Push("foo")
	stack.Push("bar")

	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
	top, ok := stack.Peek()
	if !ok || top != "bar" {
		t.Errorf("Expected top 'bar', got %v, ok=%v", top, ok)
	}
}

func TestStackPop_Int(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	val, ok := stack.Pop()
	if !ok || val != 2 {
		t.Errorf("Expected pop 2, got %v, ok=%v", val, ok)
	}
	if stack.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", stack.Size())
	}

	val, ok = stack.Pop()
	if !ok || val != 1 {
		t.Errorf("Expected pop 1, got %v, ok=%v", val, ok)
	}
	if stack.Size() != 0 {
		t.Errorf("Expected size 0 after second pop, got %d", stack.Size())
	}

	val, ok = stack.Pop()
	if ok {
		t.Errorf("Expected pop to fail on empty stack, got %v, ok=%v", val, ok)
	}
}
