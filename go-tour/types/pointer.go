package main

import "fmt"

func main() {
	i, j := 42, 2701

	// The type *T is a pointer to a T value.
	// Pointers have zero value of nil

	p := &i // Pointer to i

	// & operator generates a pointer to its operand.

	fmt.Println(*p) // reads i through pointer p

	// * operator denotes the pointer's underlying value.
	// ie: value stored at the address pointer points to.

	*p = 21 // assigns new value to i throught pointer p.

	fmt.Println(i)

	p = &j // pointers can be reassigned new addresses.

	*p = *p / 37 // divides j through pointer p.

	// This is known as dereferencing. (ie: using *p to access variable)

	fmt.Println(j)
}
