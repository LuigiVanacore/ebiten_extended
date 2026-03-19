package utils

import "sync"

// Pool is a generic object pool for reusing instances and reducing allocations.
// Use New to create a pool with a factory and optional reset function.
type Pool[T any] struct {
	mu      sync.Mutex
	items   []T
	factory func() T
	reset   func(T)
}

// NewPool creates a pool with the given factory and reset functions.
// factory is called when Get needs a new instance; reset is called when Put returns an instance.
// reset may be nil if no reset is needed.
func NewPool[T any](factory func() T, reset func(T)) *Pool[T] {
	return &Pool[T]{
		items:   make([]T, 0),
		factory: factory,
		reset:   reset,
	}
}

// Get returns an instance from the pool, or creates a new one via the factory if the pool is empty.
func (p *Pool[T]) Get() T {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.items) > 0 {
		last := len(p.items) - 1
		item := p.items[last]
		p.items = p.items[:last]
		return item
	}
	return p.factory()
}

// Put returns an instance to the pool. If reset is set, it is called before storing.
func (p *Pool[T]) Put(item T) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.reset != nil {
		p.reset(item)
	}
	p.items = append(p.items, item)
}

// Clear removes all pooled instances. Does not call reset on them.
func (p *Pool[T]) Clear() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.items = p.items[:0]
}

// Size returns the number of instances currently in the pool.
func (p *Pool[T]) Size() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.items)
}
