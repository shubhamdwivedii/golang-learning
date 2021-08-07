package main

import (
	"fmt"
	"strings"
)

func main() {

	// Arrays ########
	// The type [n]T is an array of n values of type T.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// Array's cannot be resized. An array's length is part of its type.
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices #######
	// The type []T (withour n) is a slice with elements of type T.

	// A slice is formed by specifying two indices: arr[low : high]
	var s []int = primes[1:4] // Including low upto high, excluding high.
	fmt.Println(s)            // [3 5 7]

	// Slices are like references to arrays.
	// A slice does not store any data, it just describes a section of an array.

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	as := names[0:2]
	bs := names[1:3]

	// Changing elements of a slice modifies the underlying array it refers to.
	// Any other slice sharing the same array will see those changes.

	bs[0] = "XXX"
	fmt.Println(as, bs)
	fmt.Println(names)

	// A slice literal is like an array literl without the length
	q := []int{2, 3, 5, 7, 11, 13}
	// This will create an array [6]int{2, 3, 5, 7, 11, 13} and
	// then builds a slice that references it.
	fmt.Println(q)

	sts := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(sts)

	// When slicing low and high bonds may be ommitted to use their defaults

	sts = sts[1:5] // from 1 (including) upto 5 (excluding)
	fmt.Println(sts)

	sts = sts[:3] // from 0 (default) upto 3 (excluding)
	fmt.Println(sts)

	sts = sts[1:]    // from 1 upto len(sts) (default)
	fmt.Println(sts) // last element is still there.

	sts = sts[:]     // from 0 upto len(sts)
	fmt.Println(sts) // same as before.

	r := []int{2, 3, 5, 7, 11, 13}

	// Slice has both a length and a capacity.

	length := len(r) // Length is the number of elements in the slice (current)
	fmt.Println("length:", length)

	r = r[1:5]
	capacity := cap(r) // capacity of slice is number of elements in the underlying array (counting from first element in the slice)
	length = len(r)
	fmt.Println(capacity, length) // 5 4
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.

	// Nil slices ######

	// Zero value of a slice is "nil"
	// A "nil" slice has length and capacity 0 and no underlying array.
	var ss []int
	fmt.Println(ss, len(ss), cap(ss)) // [] 0 0
	if ss == nil {
		fmt.Println("nil!")
	}

	// "make" keyword ######

	// make function allocates a zeroed zrray and returns a slice that refers to that array.
	ms := make([]int, 5)                    // len(ms) is 5
	fmt.Println("ms", ms, len(ms), cap(ms)) // [0 0 0 0 0] 5 5
	// Each element is initialized to its zero value

	// Third argument to specify capacity:
	mb := make([]int, 0, 5)
	fmt.Println("mb", mb, len(mb), cap(mb)) // [] 0 5

	// Slices can contain any type, even other Slices.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
		// Note: "strings" is a package. "string" is a type.
	}

	// Appending to a slice ######

	// Go provides a built-in "append" function.

	var sl []int
	fmt.Println(sl)

	sl = append(sl, 0) // append even works on nil slices.
	// first arg is slice, rest are elements to append.
	sl = append(sl, 1, 2, 3, 4) // The slice GROWS as needed.
	fmt.Println(sl)

	// Note: append return a new slice, original is unchanged.
	// The backing array of sl will be same as long as it can fit all given values.
	// Otherwise a newly allocated array will be assigned to the returned slice.

	// Range ########

	// "range" is used in for loop to iterate over a slice or map.
	for index, value := range sl {
		fmt.Printf("2**%d = %d\n", index, value*2)
	}

	for _, v := range sl { // index can be ommited with "_"
		fmt.Println(v)
	}

	for i := range sl { // value can also be ommited entirely (first return value is always the index)
		fmt.Println(i)
	}
}
