package main 

import "fmt" 

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine

func main() {

	// creating new channel. Channels are typed by the values they convey.
	messages := make(chan string)


	// "Send" a value into a channel using the channel <- syntax. 
	go func() {
		messages <- "ping" // Here we send "ping" to messages channel, from a new goroutine.
	}()

	// the <- channel syntax recieves a value from the channel. Here we'll receive the "ping" message we sent above and print it out.
	msg := <- messages
	fmt.Println(msg) 

	// Read more about this. 
}