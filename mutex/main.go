package mutex

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type safeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *safeCounter) inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *safeCounter) value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func UseMutex() {
	c := safeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.value("somekey"))
}
