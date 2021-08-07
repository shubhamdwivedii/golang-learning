package main

import (
	"fmt"
	"math"
)

// An "interface type" is defined as a set of method signatures.

type Abser interface {
	Abs() float64
}

// A value of interface type can hold any value that implements those methods.

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// a = v // Will Give Error
	// Vertex does NOT implements Abser (*Vertex does however)

	fmt.Println(a.Abs()) // 5

	// Interfaces are implicitly implemented
	var i I = T{"hello"}
	i.M()

	// An interface value holds a value of a specific underlying concrete type.
	describe(i) // ({hello}, main.T)
	// Calling a method on an interface value executes the method of the same name on its underlying type.
	i.M() // hello

	i = F(math.Pi) // i can be assinged an F value (as F implements I interface type of which i is an instance of)
	describe(i)    //(3.141592653589793, main.F)
	i.M()          // 3.141592653589793

	// Unlike some other languages where a "null pointer exception" would be triggered,
	var e *E // zero value of *E is nil
	i = e
	describe(i) // (<nil>, *main.E)
	// In Go it is common to write methods that gracefully handle being called with a "nil" receiver
	i.M() // <nil>
	// Note that an interface value (i) that holds a "nil" concrete value is itself "non-nil"

	// A "nil" interface value holds neither value nor concrete type.
	var ii I
	describe(ii) // (<nil>, <nil>)
	// Calling a method on "nil" interface will give run-time error. (because there is no type inside the interface to indicate which "concrete" method to call)
	// ii.M() // Will give error.

	// The Empty Interface #######

	// The interface type that specifies zero methods is known as the "empty interface":
	var j interface{}
	describe2(j) // (<nil>, <nil>)
	// describe(j) // Will give error.
	j = 42
	describe2(j) // (42, int)
	j = "hello"
	describe2(j) // (hello, string)
}

// Interfaces are implemented implicitly #######

type I interface {
	M()
}

// A Type implements an interface by implementing its methods.
// There is no explicit declaration of intent. ie: No "implements" keyword needed in GO.

type T struct {
	S string
}

// This method means type T implements the interface I
func (t T) M() {
	fmt.Println(t.S)
}

// Interface implementation could appear in any package without prearrangement. (definition of an interface is decoupled from its implementations)

type F float64

func (f F) M() {
	fmt.Println(f)
}

// Will take either T or F as they both implement I.
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i) // %T will print type
}

// If the concrete value inside the interface itself is nil,
// the method will be called with a nil receiver.

type E struct {
	S string
}

func (e *E) M() {
	if e == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(e.S)
}

// Empty Interface

// An empty interface may hold values of any type.
func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Empty interfaces are used by code that handles values of unknown type.
// Example: "fmt.Print" takes any number of arguments of type "interface{}"
