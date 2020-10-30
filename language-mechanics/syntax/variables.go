package main 
import ("fmt") 

/**********    Built-in types    **********/
	
// What is the amount of memory that we allocate? (eg. 32-bit, 64-bit)
// What does that memory represent? (eg. int, uint, bool, ..)


// Types can be specifc. Example: uint8 contains a base 10 number using one byte of memory.
// int32 contains a base 10 number using 4 bytes of memory. 

// When we declare type like just int, or uint, it gets mapped based on the architecture we are building on.
// ie: on 64-bit OS int will map to int64, on a 32-bit OS it becomes int32. 


/********** Zero value concept ***********/

// Every single value we create must be initialized.
// If we don't specify it, it will be set to the zero value. Zero values for some types are: 
// Boolean - false, Integer - 0, Floating Point - 0, Complex - 0i, String - "", Pointer - nil

/********* Declare and initialize ********/ 

// var is the only guarantee to initalize a zero value for a type. 

func main() {
	var a int 
	var b string
	var c float64 
	var d bool 

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)
	// %T prints type of the variable, %v prints value of the variable

	/* NOTE: Strings are a series of uint8 types.
	A String is a two-word data structure: 
	1. the first word represents a pointer to a backing array.
	2. the second word represents its length. 
	If a string is zero valued then the first word is nil and second word is 0. */
	
	
	// The short variable declaration operator := can define and initialize at the same time. 
	
	aa := 10 
	bb := "hello" // 1st word points to an array of characters, 2nd word is 5 bytes (length)
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n", dd, dd)

	conversion()
}

/*********** Conversion vs casting **********/

// Go doesn't have casting, instead it has conversion. 

func conversion() {
	// Instead of telling a compoiler to pretend to have some more bytes, we have to allocate more memory lke so. 
	
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}