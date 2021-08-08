package main

import (
	"fmt"
	"strconv"
	"time"
)

// Go program express error state with "error" values

// The "error" type is a build-in interface similar to "fmt.Stringer"

/*
	type error interface {
		Error() string
	}
*/

// As with fmt.Stringer, the "fmt" package
// looks for the "error" interface when printing values.

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func main() {
	// Functions often return an "error" value.
	i, err := strconv.Atoi("42")
	if err != nil { // A "nil" error denotes success.
		fmt.Printf("Error: Could not convert number: %v\n", err)
	} else {
		fmt.Println("Converted integer:", i)
	}

	// Impelementing "error" interface
	if err := run(); err != nil {
		fmt.Println(err)
	}

}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
