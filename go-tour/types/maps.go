package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// A map maps keys to values.

var m map[string]Vertex

// The zero value of map is "nil"
// A "nil" map has no keys, nor can keys be added.

func main() {
	m = make(map[string]Vertex)
	// "make" returns a map of the given type, initialized and ready for use.

	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	// Map literals are like struct literals, but the keys are required.
	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": { // Type name can be ommited here optionally.
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	// Mutating Maps #######
	mp := make(map[string]int)

	// To insert or update an element
	mp["Answer"] = 42
	fmt.Println("The value:", mp["Answer"])

	mp["Answer"] = 48
	fmt.Println("The value:", mp["Answer"])

	// To delete an element
	delete(mp, "Answer")
	fmt.Println("The value:", mp["Answer"])

	// To test if key is present use two-value assignment:
	v, ok := mp["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// If v, ok have already been decalred they can be reassigned
	v, ok = mp["Question"]

}
