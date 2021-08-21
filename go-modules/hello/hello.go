package hello

// Good practice to name package after subfolder

// This will be a submodule of shubham.com/go-modules

import "fmt"

// Only the functions and variables starting with Capital letters are exported.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
