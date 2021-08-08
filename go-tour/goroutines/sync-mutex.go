package main

import (
	"fmt"
	"sync"
	"time"
)

// Channels are great for communication among goroutines

// But if we wan't to make sure only one goroutine can access
// a variable at a time to avoid conflicts ?

// This concept is called "Mutual Exclusion" ("mutex" for short)

// Go provides mutual exclusion with "sync.Mutex" and its two method:
// "Lock" and "Unlock"

type SafeCounter struct { // Safe to use concurrently
	mu sync.Mutex
	v  map[string]int
}

// Increments counter for the given key.
func (c *SafeCounter) Inc(key string) {
	// Lock so only one goroutine at a time can access the map c.v
	c.mu.Lock()

	c.v[key]++
	c.mu.Unlock() // Unlock at the end (can also use defer at top)
}

func (c *SafeCounter) Value(key string) int {
	defer c.mu.Unlock() // Will get executed last (to Unlock)
	c.mu.Lock()         // Lock to ensure mutual exclusion
	return c.v[key]
}

func main() {
	c := SafeCounter{
		v: make(map[string]int),
	}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
