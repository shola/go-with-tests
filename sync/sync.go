package main

import "sync"

// Use channels when passing ownership of data
// Use mutexes for managing state
type Counter struct {
	mu sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// does locking "disable" concurrency because only a single goroutine can operate on the
	// value at a given time?
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}