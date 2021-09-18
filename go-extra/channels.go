package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

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

NOTE: Unbuffed Channel has buffer size zero.
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

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@ Scenarios @@@@@@@@@@@@@@@@@@@@@@@@@@@@@

// 1. Signal With Data - Guaranteed - Unbuffered Channels

// Two Scenarios: "Wait-For-Task" and "Wait-For-Result"

// a. Wait-For-Task

func waitForTask() { // Employee needs to wait for paper to perform task.
	ch := make(chan string)

	go func() {
		p := <-ch // "recieve" is blocked until "send" starts

		fmt.Println("Employee working on Paper", p)
		// Employee performs work here
		// Employee is done and free to go.
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper1" // "send" will complete when "receive" is done.

	// Only reached when employee has "received" his "paper"
	fmt.Println("Manager continues his own work")
}

// NOTE: After both channle operations (send and receive) the "Scheduler" choose to execute any statement it wants.
// The next line of code that is executed either by "manager" or "employee" is nondeterministic.
// This means next print statement can be either of the above ones.

// b. Wait-For-Result

func waitForResult() { // Manager waits for Employee to submit his paper.
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		ch <- "paper2" // "sends" paper

		fmt.Println("Employee has submitted his paper. Now he's free to go.")
	}()

	p := <-ch // "receive" paper
	fmt.Println("Manager has received paper:", p)

}

// ***************** The Cost of Guarantee - Latency **********************

// The cost of guarantee is unknown "latency".

// in Wait-For-Task - Employee has no idea how long to wait for paper.
// in Wait-For-Result - Manager has no idea how long till employee submit paper.

//======================================================================

// 2. Signal With Data - No Guarantee - Buffered Channels (Buffer > 1)

// Two Scenarios: "Fan-Out" and "Drop"

/*
To Decide how much buffer you need, ask these questions:

	1. Do I have a well defined amount of work to be completed ?
		- How much work is there ?

	2. If my employee can't keep up, can I discard any new work ?
		- How much outstanding work puts me at capacity ?

	3. What level of risk am I willing to accept if my program terminates unexpectedly ?
		- Anything waiting in the buffer will be lost.

Only use buffered channel if you can answer all these questions.
*/

// a. Fan-Out

func fanOut() { // You have n number of employees who work concurrently, thus n size buffer to receive all reports.
	emps := 20

	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper" // If two employees reach here at same time -
			// - they must take turns sending the report in the channel

			// Employee need not wait for paper to be "received", "sent" paper is buffered.
			fmt.Println("Employee submitted paper. Now can go home.")
		}()
	}

	for emps > 0 {
		p := <-ch // Will be blocked here at each iteration, until a paper is "sent"
		fmt.Println(p)
		emps--
	} // Manager will have to wait for all 20 employees to submit paper.

	fmt.Println("All Employees submitted paper. Manager now go home.")
}

// b. Drop

func selectDrop() { // Manager throws work away when an employee is at capacity
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch { // will break when close(ch) is executed.
			fmt.Println("employee : received :", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : sent paper is received")
		default: // reached if ch is blocked.
			fmt.Println("manager : drop")
		}
	}

	close(ch) // this will signal employee that they are done and free to go home.
}

// *************** Benefit of No-Guarantee - Low Latency **********************

// in Fan-Out: there is buffer space for each employee to send paper, and go home (without waiting for manager to receive paper.)
// in Drop: buffer is measured for capacity, if capacity reached work is dropped to things can keep moving.

//======================================================================

// 3. Signal With Data - Delayed Gurantee - Buffered Channel (Buffer = 1)

// When you need to know if previous signal was "received" before "sending" a new signal.

// One Scenario: "Wait For Tasks"

func waitForTasks() { // Employee is given multiple tasks, he must finish(sumbit) a task before starting a new one.
	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			// While employee works on a paper, Manager can send 1 more work in the channel which will be received on next iteration.
			fmt.Println("employee: working :", p)
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper" // if buffer is empty, you can send work (it is guaranteed that employee received last work we sent.)
		// if blocked (can't send work), we know that employee hasn't started on (received) last work we sent.
	}
	close(ch) // Work done employee and manager go home.

	// Latency is reduced if employee can work as fast as manager can send.
}

//======================================================================

// 4. Signal Without Data - Context

// Contexts leverage an Unbuffered channel that is closed to perform a signal without data.

// One Scenario: With-Timeout

func withTimeout() { // Manager is on deadline now, if employee does not finish in time, we can't wait.
	duration := 50 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	// context.WithTimeout() created a goroutine that will close the Unbuffered channel associated with context once duration is met.
	defer cancel() // We should still call cancel() regardless of how things turn out.
	// cancel() will clean up things that have been created for the Context.
	// Its okay to call cancel more than once.

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		ch <- "paper"
	}()

	select { // select here receives on two channels:
	case p := <-ch: // Waiting for employee to submit work.
		fmt.Println("Employee submitted in time. Yeah!", p)

	case <-ctx.Done(): // Waiting for timer to run out (context will be closed)
		fmt.Println("Employee didn't finish in time, moving on...")
	}

	// Which ever case is met first will get executed and we'll reach here.
	fmt.Println("One of the cases was executed.")

	/*
		An important aspect of this algorithm is the use of the Buffered channel of 1.
		If the employee doesn’t finish in time, you are moving on without giving the employee any notice.
		From the employee perspective, they will always send you the report and they are blind if you are there or not to receive it.
		If you use an Unbuffered channel, the employee will block forever trying to send you the report if you move on.
		This would create a goroutine leak. So a Buffered channel of 1 is being used to prevent this from happening.
	*/
}

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ Summary @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

/*
Unbuffered channels:
	Receive happens before the Send.
	Benefit: 100% guarantee the signal has been received.
	Cost: Unknown latency on when the signal will be received.

Buffered channels:
	Send happens before the Receive.
	Benefit: Reduce blocking latency between signaling.
	Cost: No guarantee when the signal has been received.
		- The larger the buffer, the less guarantee.
		- Buffer of 1 can give you one delayed send of guarantee.

Closing channels:
	Close happens before the Receive (like Buffered).
	Signaling without data.
	Perfect for signaling cancellations and deadlines.

nil channels:
	Send and Receive block.
	Turn off signaling
	Perfect for rate limiting or short term stoppages.
*/

// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@ Design philosophy @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

/*
If any given Send on a channel CAN cause the sending goroutine to block:
		Not allowed to use a Buffered channel larger than 1.
			- Buffers larger than 1 must have reason/measurements.
		Must know what happens when the sending goroutine blocks.

If any given Send on a channel WON’T cause the sending goroutine to block:
	You have the exact number of buffers for each send.
		- Fan Out pattern

	You have the buffer measured for max capacity.
		-Drop pattern

Less is more with buffers.
	Don’t think about performance when thinking about buffers.

	Buffers can help to reduce blocking latency between signaling.
		- Reducing blocking latency towards zero does not necessarily mean better throughput.
		- If a buffer of one is giving you good enough throughput then keep it.
		- Question buffers that are larger than one and measure for size.
		- Find the smallest buffer possible that provides good enough throughput.
*/

// ##################################################################################################
// ##################################################################################################
