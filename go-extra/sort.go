package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)            // in-built function to sort strings (increasing order)
	fmt.Println("Strings:", strs) // Strings: [a b c]

	ints := []int{7, 2, 4}
	sort.Ints(ints)            // in-built function to sort ints (increasing order)
	fmt.Println("Ints:", ints) // Ints: [2 4 7]

	s := sort.IntsAreSorted(ints) // in-built function to check if sorted.
	fmt.Println("Sorted:", s)     // Sorted: true

	// ###### Sort A Slice #############

	colors := []string{"blue", "yellow", "red", "green"}

	// This is similar to just implementing the Less() function
	sort.Slice(colors, func(i, j int) bool {
		return len(colors[i]) < len(colors[j])
	})
	fmt.Println(colors) // [red blue green yellow]

	//###### Custom Sorting Functions ######

	// Implement sort.Interface { Len(), Swap(), Less() }

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits) // [kiwi peach banana]
}

type byLength []string

// Len returns number of elements in the collection.
func (s byLength) Len() int {
	return len(s)
}

// How to swap the elements
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// How to decide if one item is less than other
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
