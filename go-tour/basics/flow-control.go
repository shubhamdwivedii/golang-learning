package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	sum := 0

	// A basic for loop in Go
	for i := 0; i < 10; i++ { // Init; Condition; Post
		sum += i
	}
	fmt.Println(sum)

	// Init and Post are optional
	for sum < 100 { // While loop basically
		sum += sum
	}

	for { // Infinite loop
		if sum >= 200 { // If statement in go (Parenthesis are options if (sum > 200) is also valid)
			break // break and continues are available in Go's for
		}
		sum += 50
	}
	fmt.Println(sum)

	fmt.Println(
		pow(3, 3, 10),
		pow(3, 3, 20),
	) // Both "pow" return their results before Println begins.

	switchCase()
	defered()
}

// If syntax with Short Statement #########

func pow(x, n, lim float64) float64 {
	// Like "for", "if" can start with a short statement before executing condition:
	if v := math.Pow(x, n); v < lim { // Init; Condition
		return v // v will only be in scope until the end of "if"
	} else { // v is also available in else block.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// Can't use v here.
	return lim
}

// Swith Case in Go

func switchCase() {
	fmt.Print("Go runs on ")

	// Like "for" and "if" "switch" can also have Init statement before Condition

	switch os := runtime.GOOS; os { // Init; Condition
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("Windows.")
	case "freebsd", "openbsd": // case can have multiple match strings
		fmt.Println("BSD based.")
	default:
		fmt.Println("Unknown.")
	}

	// NOTE: In Go's switch, only the first selected case is run (not the cases that follow),
	// Thus No Need for "break" as in other languages.

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0: // Go's cases can have Conditions too.
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today - 1 + 1: // will never be reaches as same as first case
		fmt.Println("Never Reachable.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	// Switch with no condition is same as "switch true"
	switch { // Can be used in place of long if-else-ifs chains
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer ##########

func defered() {
	// "defer" defers the execution of function until the surrounding function (in same scope) returns.
	defer fmt.Println("world")
	// deffered call's argument are evaluated immediately, but execution only starts after surrounding functions returns
	fmt.Println("hello")

	// Defered function calls are pushed onto a STACK.
	// Calls are executed in Last-In-First-Out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // last pushed (ie: 9) will be printed first.
	}
	fmt.Println("will print before counting starts")
	// "world" would get printed last as it was first to go in STACK.
}

// NEXT >> Pointers-Structs-Slices-Maps
