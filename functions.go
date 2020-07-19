package main 

import "fmt"

// Functions in Go ######################

func plus(a int, b int) int { // both parameters are of type int, return type is also int. 
	return a + b
}

// Go requires explicit returns. ie: it won't automatically return the value of the last expression. 

func plusPlus(a, b, c int) int { // decalre type this way for multiple consecutivce parameters of same type. 
	return a + b + c 
}

// Multiple Return Values #####################

// Go has support for multiple return values (just like lua)

// Eg. to return both result and error values 

func vals() (int, int) {  // signature (int, int) shows that functin return two ints. 
	return 3, 7
}


// Variadic Functions ########################

// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


// Closures ################################

// Go supports anonymous function (like lambda functions in python or arrow functions in javascript), which can form closures. 
// Anonymous functions are useful when you want to define a function inline without having to name it. 

// intSeq() returns another function, which we define anonymously in its body. 
func intSeq() func() int { // "func() int" is the signature (return type) for intSeq()
	i := 0 
	return func() int {
		i++  // The returned function "closes over" the variable i to form a closure.
		return i 
	}
}


// Recursion #############################

// Classic factorial example.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}


func main() {
	res := plus(1, 3)
	fmt.Println("1+2 =", res)
	
	res = plusPlus(1, 2, 3) 
	fmt.Println("1+2+3 =", res)	

	a, b := vals() 
	fmt.Println(a, b)

	_, c := vals() // use blank identifier "_" if you are only interested in a subset of return values. 
	fmt.Println(c) 

	sum(1, 3) // calling variadic function with any number of arguments
	sum(1, 3, 4)
	nums := []int{1, 2, 3, 4}
	sum(nums...) // if you have multiple arguments in a slice (or array ?), apply them to a variaic function using func(slice...)  
	// just like using spread operator in javascript (but not exaclty).	

	nextInt := intSeq() // nextInt is assigned the returned value of intSeq which is an anonymous function. 
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3
	fmt.Println(nextInt()) // 4
	
	newInts := intSeq() // each time intSeq() is called the returned anonymous function captures its own version of "i", which it updates on each call.
	fmt.Println(newInts()) // 1
	fmt.Println(newInts()) // 2


	fmt.Println("Factorial of 7:", fact(7))

}

