package utils

// Stack is a generic LIFO stack.
type Stack[T any] struct {
	data []T
}

// NewStack creates a new stack with an initial capacity.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0, 10)}
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element from the stack.
// It returns the zero value of T and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	index := len(s.data) - 1
	val := s.data[index]
	s.data = s.data[:index]
	return val, true
}

// Peek returns the top element without removing it.
// It returns the zero value of T and false if the stack is empty.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

// IsEmpty reports whether the stack contains no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return len(s.data)
}
