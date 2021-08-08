package main

import (
	"fmt"
	"time"
)

// A "goroutine" is a lightweight thread managed by the GO runtime.

// go f(x, y, z) // starts a new goroutine running f(X, y, z)

// The evaluation of f, x, y, z happens in the current goroutine and
// the execution of f happens in the new goroutine.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// Go routines run in same address space, so access to shared memory MUST be synchronized.
// The "sync" package provides useful primitives.
