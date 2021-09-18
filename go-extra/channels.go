package main

// Do not think of Channels as a Queue (data structure)
// It is best to forget about how channels are structured and
// Focus on how they behave: "Signaling"

// A channel allows one goroutine to signal another goroutine about a particular event.

// Thinking of channels as signaling mechanism will allow us to write better code.

/*
Three attributes of signalling in Go
	1. Guarantee Of Delivery
	2. State
	3. With or Without Data
*/

// These three attributes work together to create a design philosophy around signaling.

// ################ 1. Guarantee Of Delivery ####################

// Do I need a guarantee that the signal sent by a particular goroutine has been RECEIVED ?
/*
	go func() {
		p := <-ch // Receiver
	}()

	ch	<- "paper" // Send
*/
// Does the sending goroutine need a guarantee that "paper" being set over was received before moving on ?

// ******************* Unbuffered vs Buffered Channels **********************

// Unbuffered - Guaranteed Delivery
// Buffered - Delivery Not Guranteed

// ########################## 2. State ############################

// The behavior of channel is directly influenced by its current State.

/*
Three states of channel:
	1. nil
	2. open
	3. closed
*/

// A channel is in "nil" state when its declared to its zero value (nil)
// var ch chan string

// A channel can be placed in a nil state by explicitly setting it to nil
// ch = nil

// A channel is in "open" state when it's made using the built-in functions make.
// ch := make(chan string)

// A channel is in closed state when it's closed using the built-in function close.
// close(ch)

/*
The State determines how "send" and "receive" opertions behave:
	1. nil - send: blocked, receive: blocked
	2. open - send: allowed, receive: allowed
	3. closed - send: panic, receive: allowed
*/
// Note: Blocked means code execution won't proceed furthur than the receive/send statement.

// ######################## 3. With And Without Data ###########################

// You signal "with" data by performing a send on a channel.

// ch <- "paper"

/*
When you signal with data, it's usually because:
	- A goroutine is being asked to start a new task.
	- A goroutine reports back a result.
*/

// You signal "without" data by closing a channel.
// close(ch)

/*
When you signal without data, it's usually because:
	- A goroutine is being told to stop what they are doing.
	- A goroutine reports back they are done with no result.
	- A goroutine reports that it has completed processing and shut down.
*/

/*
One benefit of signaling "without" data:
	- A single goroutine can signal many goroutine at once.
	- Signaling "with" data is always 1 to 1 exchange b/w goroutines.
*/

// ******************** Signaling With Data ************************

/*
When you signal with data, there are three channel configurations options
you can choose depending on the type of gurantee you need.

		1. Unbuffered Channel - Guarantee of Delivery
			- Because the "Receive" of the signal "Happens Before" the "Send" of the signal completes.

		2. Buffered > 1 - No Guarantee of Delivery
			- Because the "Send" of the signal "Happens Before" the "Receive" of the signal completes.

		3. Buffered = 1 - Delayed Guarantee
			- It can guarantee that the previous signal that was sent has been received.
			- Because the "Receive" of the "First Signal", "Happens Before" the "Send" of the "Second Signal" completes.

*/

// Size of the buffer must never be a random number, it must always be calculated for some well defined constraint.

// ******************* Signaling Without Data ************************

// Signalling without data is mainly reserved for "cancellation"
// It allows one goroutine to signal another goroutine to "cancel" what they are doing and move on.

/*
When you signal without data, there are three options:

		1. context.Context - First Choice (Prefered)
		2. Unbuffered channel - Second Choice (Okay)
		3. Buffered channel - Bad Choice (Code will smell BAD)
*/

/*
The built-in function "close()" is used to signal without data.
	- You can still receive signals on a "closed" channel.
	- Any "receive" on a closed will not block and the "receive" operation always returns.
*/

// In most cases you want to use the standard library "context" package to implement signaling without data.

// "context" uses an Unbuffered channel underneath for the signaling and the built-in function "close" to signal without data.

/*
NOTE: if you choose to use your "own" channel for cancellation:
	Your channel should be of type "chan struct{}".
	It is zero-space, idiomatic way to indicate a channel is used only for signalling.
*/

// ###################### Scenarios #############################
