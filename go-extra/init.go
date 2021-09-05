package main

import "fmt"

// In Go, the "init()" can be used withing a 'package' block.
// regardless of how many times that package is imported, "init()" is only called once.

// Since init() is only called once, it can be used to set up database connection,
// register various services, or perform any task you only want to do once.

var name string

func init() {
	fmt.Println("This will get called on main package initialization")
	name = "Shubham"
}

func main() {
	fmt.Println("Init Example Program")
	fmt.Println("My name is", name)
}

// Init is called implicitly. We don't need to explicitly call it anywhere in our program.
