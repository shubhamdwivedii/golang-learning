package main

// Run Command: "go mod init shubham.com/go-modules"

// This will create a go.mod file to track your code's dependencies.
// As you add dependencies, go.mod will list the versions you code depends on.

// First check out /hello/hello.go

import (
	"fmt"

	"shubham.com/go-modules/hello"
	// submodules can be imported entire package name (checkout go.mod file)

	// Named imports are also supported
	h "shubham.com/go-modules/how"

	// Import as . to use exported elements directly.
	. "shubham.com/go-modules/bye"
	// You can also import as _ for side effects of package (Google it.)
)

func main() {
	message := hello.Hello("Shubham")
	fmt.Println(message)
	h.How()
	Bye("Shubham")
}

// Check out the hello/
