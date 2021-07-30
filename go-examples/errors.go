package main 

import (
	"errors" 
	"fmt" 
)

// In Go, by convention errors are usually returned as the last return value and have type "error" (a built-in interface)
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42") // errors.new() constructs a basic "error" value with the given error message
	}
	return arg + 3, nil // a nil value indicates there was no error.
}

// It's possible to use custom types as errors by implementing the Error() method on them.
type argError struct { // represents argument error.
	arg int 
	prob string
}
func (e *argError) Error() string { // implementing Error() method on argError
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}


func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} 
		// using &argError syntax to build a new struct.
	}
	return arg + 3, nil
}


func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil { // an inline error check (e != nil) on the if line is a common practise Go code ) 
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 workded:", r)
		}
	}

	_, e := f2(42) 
	if ae, ok := e.(*argError); ok { //  getting the error as an instance of the custom error type via type assertion.
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}