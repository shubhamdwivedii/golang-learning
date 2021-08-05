package main // First line of each file should declare the package it belongs to
// package name decides if you are making an Executable (named "main") or Reusable package (can have any name)

import "fmt" // A standard library package for formatted I/O

func main() {
	fmt.Println("Hello World!!")
}

// execute by > go run hello-world.go
// to build executatble > go build hello-world.go

// Start a Go Project: go mod init github.com/USERNAME/my-go-project

// Test comment