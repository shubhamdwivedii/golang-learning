package main

import (
	"fmt"
	"time"
)

// Go elegantly supports rate limiting with goroutines, channels and tickers.

func main() {
	requests := make(chan int, 5)

	// Suppose we want to limit our handling of incoming requests.
	// We’ll serve these requests off a channel of the same name.

	for i := 1; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	// limiter channel will receive a value every 200 milliseconds.
	// This is the regulator in our rate limiting scheme.

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	// By blocking on a receive from the limiter channel before serving each request,
	// we limit ourselves to 1 request every 200 milliseconds.

	// We may want to allow short bursts of requests in our rate limiting scheme while preserving the overall rate limit.
	// We can accomplish this by buffering our limiter channel.
	burstyLimiter := make(chan time.Time, 3)
	// This burstyLimiter channel will allow bursts of up to 3 events.

	// Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now simulate 5 more incoming requests.
	// The first 3 of these will benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

// For first batch, every request will handled every 200 milliseconds

// For seconds batch, first 3 request will be handled immediately, rest will thake 200 milliseconds.
