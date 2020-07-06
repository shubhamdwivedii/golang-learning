// Go-Lang Basics - Shubham Dwivedi. 

// to build: "go build cheatsheet.go"
// to run: "./cheatsheet"

package main // tell which package this go file is part of 
import "fmt" 

func main() {
	// Hello World ###############
	fmt.Printf("hello, world\n") 

	// Values ####################
	fmt.Println("go" + "lang") // strings 

	fmt.Println("1+1+2 =", 1+1+2) // integers
	fmt.Println("7.0/3.0 =", 7.0/3.0) // floats
 
	fmt.Println(true && false) // false 
	fmt.Println(true || false) // true 
	fmt.Println(!true) // false

	// Variables ##############

	// In Go, variables are explicitly declared and used by the compiler to eg. check type-correctness of function calls. 
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1,2 // var declares 1 or more variables.  
	fmt.Println(b,c)

	var d = true 
	fmt.Println(d) // Go will INFER the type of initialized variables. 

	var e int // varables declared without initialization are ZERO-VALUED. eg. the zero-value for an int is 0. 
	fmt.Println(e)

	f := "apple" // := is shorthand syntax for declaring and initializing a variable. 
	var g string = "apple" // should be same as above
	fmt.Println(f, g)


	// Constants ################
	
	
}