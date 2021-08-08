package main

import "fmt"

// Channels are a typed conduit through which you can send and receive values (using operator <-)
/*
	ch <- v   // Send v to channel ch.
	v := <-ch // Receive from ch, and assign to value to v.
*/

// Like maps and slices, channels must be created before use.

func sum(s []int, c chan int) { // Receives a channel of type int as second argument.
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// Like maps and slices, channels must be created before use:
	ch := make(chan int)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch // receive from channel ch

	// By default, sends and receives block until other side is ready.
	// This allows goroutines to synchronize without explicit locks or condition variables.

	fmt.Println(x, y, x+y) // -5 17 12

	// Buffered Channels ########

	// Channels can be "buffered". (provide the buffer length as second argument to "make")
	chbf := make(chan int, 2)

	// Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.

	chbf <- 1
	chbf <- 2
	// chbf <- 3 // fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-chbf)
	fmt.Println(<-chbf)

	// Range and Close ##########

	// To test whether a channel has been closed
	if v, ok := <-ch; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Channel ch was closed")
	}

	// A sender can "close" a channel to indicate that no more values will be sent.
	c := make(chan int, 10)

	go fibonacci(cap(c), c) // cap(c) returns the capacity of the channel. (just like a slice)

	// The loop "for i := range c" receives values from the channel repeatedly until its closed.
	for i := range c { // Basically iterating over a channel
		// (i := <-c in each iteration)
		fmt.Println(i)
	}
	//NOTE: "range c" loop will never break if the channel isn't closed.
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Closing a channel.
	close(c)
}

// Note: ONLY the sender should close a channel, NEVER the receiver.
// Sending on a closed channel will cause panic.

// Channels aren't like files; you don't usually need to close them.
// Closing IS ONLY necessary when the receiver must be told there are no more values coming
// Such as to terminate the "range" loop.
