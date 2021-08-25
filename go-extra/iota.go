package main

import "fmt"

// The "iota" keyword represents successive integer constants 0,1,2...

// It resets to 0 whenever the word "const" appears in code.
func main() {
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Println(c0, c1, c2) // 0 1 2

	// short syntax
	const (
		b0 = iota
		b1
		b2
		b3
	)
	fmt.Println(b0, b1, b2, b3) // 0 1 2 3

	// Start from one
	const (
		a1 = iota + 1
		a2
		a3
	)
	fmt.Println(a1, a2, a3) // 1 2 3

	// Skip a value
	const (
		d1 = iota + 1
		_
		d3
		d4
	)
	fmt.Println(d1, d3, d4) // 1 3 4

	// Using Enum types (see below)
	var d Direction = North
	fmt.Print(d) // Print will call d.String()
	switch d {
	case North:
		fmt.Println(" goes up.") // North goes up.
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}

// Enum type with strings [best practice]
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}
