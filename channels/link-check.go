package main

import (
	"fmt"
	"net/http"
	"time"
)

// Website status checker continuous
func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://shubhshubh.org",
	}

	c := make(chan string) // "chan" keyword to create channel (type string means this channel can only communicate in strings)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for { // Infinite Loop. Will check link status indefinitely
	// 	go checkLink(<-c, c) // Starts a new go routine as soon as link is received from channel.
	// }

	// This will still be an infinite loop. It will just wait for channel to emit value before each iteration.
	for l := range c { // In Go we can iterate through channels too.
		// An iteration executes when a value (link) is received via channel.
		// Each iteration will wait for channel to emit some value.

		// time.Sleep(time.Second * 2) // This will wait 2 second to start new go routine
		// go checkLink(l, c)

		// A function literal in Go is an unamed function to wrap some code to execute it later.
		// Function Literal are similar to Python's lambda expression or JS's arrow function (anonymous)
		/* go func() { // We add pause inside a new go routine so that there would be no wait before starting new checkLink routine.
			time.Sleep(time.Second * 2) // Will pause the current go routine (func literal routine here) for 2 seconds (time.Second == 1 Second)
			checkLink(l, c)             // The "l" here is referencing a variable in above scope (for loop)
			// This "l" might change value by the time its used by checkLink (due to the time.Sleep)
			// Note: We are not passing "l" by value to the function literal.
		}() // Incorrect Way */

		// Correct Way
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) { // will now also receive a string type channel
	_, err := http.Get(link) // we only need err to verify if website is online
	// http.Get() is a blocking code.
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
