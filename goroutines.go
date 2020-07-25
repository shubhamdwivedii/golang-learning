package main 

import (
	"fmt" 
	"time"
)


// A goroutine is a lightweight thread of execution.

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct") // function f here is called the usual way (synchrounously)

	// to invoke f() in a goroutine, use go f(s)
	go f("goroutine") // this new goroutine will execute concurrently with the calling one.

	// a goroutine can also be started for an anonymous function.
	go func(msg string) {
		fmt.Println(msg)
	}("going")


	// Waiting for goroutines to finish. use a "WaitGroup" for more robust apporach.
	time.Sleep(time.Second)
	fmt.Println("done")
}