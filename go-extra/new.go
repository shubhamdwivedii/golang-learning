package main

import "fmt"

// Just like make(), new() is used to allocate memory

// Memory returned by new() is zeroed.
// new() only returns pointers to initialized memory.
// new() works for all data types (except channel, map)
// new() dynamically allocates space for variable of that type
// new() initializes variable it to zero value of its type and returns a pointer to it.

func main() {
	result := new(int)
	fmt.Println(result) // 0xc000012078

	// This is same as below
	var temp int
	var result2 *int
	result2 = &temp
	fmt.Println(result2) // 0xc0000120b0

	// works for structs too.
	shubham := new(person)
	fmt.Println(shubham) // &{"" 0}
}

type person struct {
	name string
	age  int
}
