package sync

import "sync"

// summary
// mutex allows us to add locks to our data
// waitgroup is a means of waiting for goroutines to finish jobs

// use channels when passing ownership of data
// use mutexes for managing state

// mutex is a mutual exclusion locl. the zero value of a mutex is an unlocked mutex
type Counter struct {
	mu    sync.Mutex
	value int
}

// any goroutine calling Inc will acquire the lock on Counter if they are first. All the other goroutines will have to wait for it to be Unlocked before getting access
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}
