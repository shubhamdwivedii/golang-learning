package main

import (
	"fmt"
	"math"
)

// Functions are values too. They can be passed around just like other values.
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function values may be used as arguments and return values.
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
		// Value of "sum" will be different for both pos and neg
	}

}

// Function CLOSURES

// Go functions my be closures.
// A "closure" is a function value that references variables from outside its body.

func adder() func(int) int { // adder returns a closure.
	sum := 0
	return func(x int) int {
		sum += x // closures my access and assign to the referenced variables.
		// In a sense the function is "bound" to the variable.
		return sum
		// Each closure is bound to its OWN "sum" variable.
	}
}
