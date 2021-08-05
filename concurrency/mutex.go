package main

import (
	"fmt"
	"sync"
	"time"
)

type MyStructUnitialized struct {
	name string
	age  uint8
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu                  sync.Mutex
	v                   map[string]int
	nyStructUnitialized MyStructUnitialized
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	//If i dont initialize some property, these will be created with default values in your props
	c := SafeCounter{v: make(map[string]int)}

	fmt.Printf("%p=myUnitialized %s", &c.nyStructUnitialized, "\n")

	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	var teste interface{}
	teste = SafeCounter{v: make(map[string]int)}
	_, ok := teste.(sync.Mutex)

	fmt.Println("Ok type assertion, ", ok)
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
