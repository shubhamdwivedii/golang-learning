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


	// Channel Buffering ####################
	
	// By default channels are "unbuffered", meaning that they will only accept sends (channel <-) if there is a corresponding receive (<- channel) ready to receive the sent value. 

	// "Buffered channels" accept a limited number of values without a corresponding receiver for those values. 
		
	messages2 := make(chan string, 2) // Here we make a channel of strings buffering up to 2 values.

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages2 <- "buffered" 
	messages2 <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages2)
	fmt.Println(<-messages2)
}