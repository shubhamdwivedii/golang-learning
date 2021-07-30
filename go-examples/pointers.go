package main 

import "fmt" 

// Go supports pointers, allowing you to pass references to values and records within your program.


// zeroval has an int parameter, so arguments will be passed to it by value. 
func zeroval(ival int) {
	ival = 0 
}

// zeroptr has *int parameter, meaning it takes an int pointer. ie arguments are passed by reference.
func zeroptr(iptr *int) {
	*iptr = 0 
}


func main() {
	i := 1 
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // &i means memory address of variable i (ie: a pointer to i)
	fmt.Println("zeroptr:", i) 

	fmt.Println("pointer:", &i) // pointers can be printed too.
}