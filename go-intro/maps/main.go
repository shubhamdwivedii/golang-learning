package main

import "fmt"

func main() {
	// Maps in Go are similar to Dictionaries in Python or Objects in JavaScript.

	colours := map[string]string{ // [key-type]value-type
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
		"blue":  "#2bf885",
	}
	// All keys in map have to be of same type, All values also need to be same type.
	fmt.Println(colours)

	var shapes map[string]string
	fmt.Println(shapes) // will be initialized to zero value > map[]

	// Cannot add values later.
	// shapes["circle"] = "piRsquare" // Invalid

	// "make" is a keyword in go.
	sounds := make(map[string]string) // creates an empty map.
	fmt.Println(sounds)               // same as above > map[]

	// Adding value to map created by "make"
	sounds["loud"] = "2kHz"
	fmt.Println(sounds)

	// fmt.Println(sounds.loud) // invalid
	fmt.Println(sounds["loud"])

	// To delete a key/value pair
	delete(sounds, "loud")
	fmt.Println(sounds)

	printMap(colours) // Iterating over maps

	// Maps are referece types (therfore passed by reference)
	updateMap(colours)
	fmt.Println(colours)
}

func printMap(c map[string]string) {
	// Iterating through a Map
	for color, hex := range c { // similar to Slices (instead of index we have keys)
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// will receive a reference
func updateMap(c map[string]string) {
	for color, hex := range c {
		c[color] = hex + "updated"
	}
}

/* Map vs Struct
1. All keys and all values in Map are same type. In Struct values can be different types.
2. No need to know all keys at compile time for Map. For Struct all fields must be known at compile time.
3. Map are reference type. Struct are value type.
4. Map keys are indexed, can be itereated over. Structs cannot be iterated over as keys are not indexed.
*/
