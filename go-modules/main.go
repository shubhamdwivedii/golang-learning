package main

// Run Command: "go mod init shubham.com/go-modules"

// This will create a go.mod file to track your code's dependencies.
// As you add dependencies, go.mod will list the versions you code depends on.

// First check out /hello/hello.go

import (
	"fmt"

	"shubham.com/go-modules/hello"
	// submodules can be imported entire package name (checkout go.mod file)
)

func main() {
	message := hello.Hello("Shubham")
	fmt.Println(message)
}

// Check out the hello/
