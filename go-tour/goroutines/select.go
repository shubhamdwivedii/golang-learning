package main

import (
	"fmt"
	"time"
)

// The "select" statement lets a goroutine wait on multiple communication operations.

func fibonnaci(ch, quit chan int) {
	x, y := 0, 1
	for {
		// A select blocks until one of its cases can run, then it executes that case.
		select { // It chooses one case at random if multiple are ready.
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() { // A function literal (anonymous function basically) are often used to run new goroutines.
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonnaci(c, quit)

	// Default Selection ##########

	// The "default" case in a "select" is run if no other case is ready. (Or channel is blocked)

	// Use a "default" case to try a send or receive without blocking:
	/*
				select {
				case i := <-c:
		    		// use i
				default:
					// receiving from c would block
				}
	*/

	tick := time.Tick(100 * time.Millisecond)  // like setInterval()
	boom := time.After(500 * time.Millisecond) // like setTimeout()

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
