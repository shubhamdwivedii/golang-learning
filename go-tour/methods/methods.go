package main

import (
	"fmt"
	"math"
)

// Go does not have Classes. However we can define methods on classes.

type Vertex struct {
	X, Y float64
}

// A "method" is a function with a special "receiver" argument.

// The "receiver" apprears in its own argument list (b/w the "func" and method name)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// A method is just a function with a receiver argument.

func main() {
	v := Vertex{3, 4}
	abs := v.Abs()   // methods can be called with "." if defined.
	fmt.Println(abs) // 5

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer Receiver
	v.Scale(10) // a pointer to v (&v) will automatically be passed to Scale here.
	// No need to (&v).Scale(10)
	// Go will interpret v.Scale(5) as (&v).Scale(5) automatically

	// Same thing happens in reverse too.
	p := &v
	// Since "Abs" has a value receiver type.
	res := p.Abs() // No need to call as (*p).Abs()
	// Go will automatically interpret (&v).Abs() as *(&v).Abs() (or v.Abs() or *p.Abs())

	fmt.Println(res)

}

type MyFloat float64

// Methods can be declared on non-struct types too.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// NOTE: You can only declare a method with a receiver -
// whose type is defined in the same package as the method.

// You cannot declare a method with a receiver whose type is defined in another package
// This includes the built-in types such as "int" (eg. you can't declare method for int type)

// Pointer Receivers ############

// You can delcare methods with pointer receivers
// These methods have literal syntax "*T" for some typ "T" (T cannot be a pointer itself)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f // These methods can modify the value to which the recevier points.
	v.Y = v.Y * f
}

// Since methods often need to modify their receiver,
// Pointer receivers are more common than value receivers.

// With a value receiver, the "Scale" method will operate on a copy of original "Vertex" value.

/* Choosing a value or pointer receiver.

Two reasons to choose a pointer receiver:

1. Method can modify the value its receiver points to.
2. To avoid copying the value on each method call. (more memory efficient for large structs)

Even if "Abs()" don't modify the value of Vertex, making it a pointer receiver method would save memory. */

// General Rule: All methods on a given type should have
// either value or pointer receivers, but not a mixture of both.
