package main

import "fmt"

// 3 dots operator has 4 different use in GO:

// 1. Variadic Function Parameters ###########

func Sum(nums ...int) int { // type of nums here is []int
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// 2. Arguments to variadic functions ########

func main() {
	primes := []int{2, 3, 4, 5}
	fmt.Println(Sum(primes...)) // can pass a slice this way.

	// 3. Array literals ######################

	stooges := [...]string{"Moe", "Larry", "Curly"} // len(stooges) = 3
	// In array literal ... notation specifies a length equal to the number of elements in the literal.
	fmt.Println(len(stooges))
}

// 4. Go Command

//> go test ./...

// Will test all packages in current directory and its subdirectories.
