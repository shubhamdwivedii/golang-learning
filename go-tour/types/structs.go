package main

import "fmt"

// A struct is a collection of fields.

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2} // Vertex{1, 2} created a new Vertex
	fmt.Println(v)
	v.X = 4 // Struct fields are accessed using a dot.
	fmt.Println(v.X)

	p := &v   // a pointer to struct v
	p.X = 1e9 // Struct fields can be accesses via a struct pointer.
	// No need for (*p).X
	fmt.Println(v) // {1000000000, 2}

	var (
		v1 = Vertex{1, 2}       // X:1 Y:2 // Order is important here
		v2 = Vertex{X: 1}       // Y:0 implicitly // Order not important here
		v3 = Vertex{}           // X:0 Y:0 implicitly
		v4 = Vertex{Y: 5, X: 6} // Order not important when specifying field names
		pt = &Vertex{1, 2}      // has type *Vertext (thus is a pointer)
	)

	fmt.Println(v1, pt, v2, v3, v4)
}
