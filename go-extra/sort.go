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

	//###### Custom Sorting Functions ######

	// Implement sort.Interface { Len(), Swap(), Less() }

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits) // [kiwi peach banana]
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
