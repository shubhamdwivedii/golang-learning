package main

import (
	"fmt"
	"net/http"
)

// Concurrency IS NOT Parallelism

/*
	Concurrency => Multiple threads executing code per CPU Core.
	If one thread is blocked, another is picked up and worked on.

	Parallelism => Multiple threads execute at same time in multiple CPU Cores.
*/

// Website status checker
func main() { // Main routine is created when we launch the program.
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
		// checkLink(link) // each link status is print after a delay
		// go checkLink(link) // using "go" before function call will start a new "Go Routine"
		// These go routines (started with "go") are called child routines (to our main routine).

		// Channels are used to communicate between different Go Routines.
		go checkLink(link, c)
	}

	// fmt.Println(<-c) // use "<- channel" to receive data from channel
	// Note: using "<-channel" will wait for the data (ie: it is a blocking code)

	// This will execute when the next value is received from the channel.
	// fmt.Println(<-c) // If no value is received the main routine will keep on waiting here. (Basically hangs)

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) // We will still wait for the channel to receive value.
	}

}

func checkLink(link string, c chan string) { // will now also receive a string type channel
	_, err := http.Get(link) // we only need err to verify if website is online
	// http.Get() is a blocking code.
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might Be down??" // use "channel <-" to send data to channel
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yup, Its up!"
}
