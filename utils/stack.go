package utils


// Stack structure
type Stack[T any] struct {
    data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{ data: make([]T, 0, 10) }
}

// Push adds an element to the stack
func (s *Stack[T]) Push(value T) {
    s.data = append(s.data, value)
}

// Pop removes and returns the top element from the stack
// Returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {

    if len(s.data) == 0 {
        var zero T
        return zero, false 
    }
    // Get the last element
    top := s.data[len(s.data)-1]
    // Remove the last element
    s.data = s.data[:len(s.data)-1]
    return top, true
}

// Peek returns the top element without removing it
// Returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Peek() (T, bool) {
    if len(s.data) == 0 {
        var zero T
        return zero, false // Stack is empty
    }
    return s.data[len(s.data)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
    return len(s.data) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
    return len(s.data)
}