package main

import (
	"fmt" 
	"math" 
)

// In Go, Interfaces are named collection of method signatures. 

// A basic geometry interface 
type geometry interface {
	area() float64 
	perim() float64
}

type rect struct {
	width, height float64 
}

type circle struct {
	radius float64 
}

// To implement an interface in Go, we just need to implement all the methods in the interface.

// Implementing geometry on rect
func (r rect) area() float64 {
	return r.width * r.height 
}
func (r rect) perim() float64 {
	return 2 * r.width + 2 * r.height 
}

// Implementing geometry on circle 
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius 
}

// If a variable has an interface type, then we can call methods that are in the named interface.

func measure(g geometry) { // this measure function takes advantage of this to work on any geometry.
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// circle and rect both implement geometry. So we can use instances of them as arguments in measure() function.
	measure(r)
	measure(c)
}
