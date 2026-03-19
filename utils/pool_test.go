package utils

import (
	"sync"
	"testing"
)

func TestPoolGetPut(t *testing.T) {
	created := 0
	factory := func() int {
		created++
		return created
	}
	reset := func(i int) {
		// no-op for int
	}

	pool := NewPool(factory, reset)

	// Get creates new
	a := pool.Get()
	if a != 1 {
		t.Errorf("first Get: got %d, want 1", a)
	}
	b := pool.Get()
	if b != 2 {
		t.Errorf("second Get: got %d, want 2", b)
	}

	// Put returns to pool
	pool.Put(a)
	pool.Put(b)

	// Get reuses
	c := pool.Get()
	if c != 2 {
		t.Errorf("Get after Put: got %d, want 2 (reused)", c)
	}
	d := pool.Get()
	if d != 1 {
		t.Errorf("Get after Put: got %d, want 1 (reused)", d)
	}
}

func TestPoolReset(t *testing.T) {
	factory := func() int { return 42 }
	resetCalls := 0
	reset := func(i int) {
		resetCalls++
	}

	pool := NewPool(factory, reset)
	pool.Put(pool.Get())

	if resetCalls != 1 {
		t.Errorf("Put should call reset once, got %d", resetCalls)
	}
}

func TestPoolClear(t *testing.T) {
	pool := NewPool(func() int { return 1 }, nil)
	pool.Put(pool.Get())
	pool.Put(pool.Get())

	pool.Clear()
	if pool.Size() != 0 {
		t.Errorf("Clear: Size got %d, want 0", pool.Size())
	}

	// Get after Clear creates new
	v := pool.Get()
	if v != 1 {
		t.Errorf("Get after Clear: got %d", v)
	}
}

func TestPoolSize(t *testing.T) {
	pool := NewPool(func() int { return 0 }, nil)
	if pool.Size() != 0 {
		t.Errorf("empty pool Size: got %d", pool.Size())
	}

	pool.Put(pool.Get())
	if pool.Size() != 1 {
		t.Errorf("after Put Size: got %d, want 1", pool.Size())
	}

	pool.Get()
	if pool.Size() != 0 {
		t.Errorf("after Get Size: got %d", pool.Size())
	}
}

func TestPoolConcurrent(t *testing.T) {
	pool := NewPool(func() int { return 1 }, nil)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v := pool.Get()
			pool.Put(v)
		}()
	}
	wg.Wait()
}
