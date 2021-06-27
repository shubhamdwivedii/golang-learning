package main

import "fmt"

type contactInfo struct { // added later.
	email   string
	zipCode int
}

// Struct: collection of different properties that are related together.
type person struct {
	firstName string // order may be important
	lastName  string
	// contact   contactInfo // can include other structs
	contactInfo // also valid.
}

func main() {
	// shubham := person{"Shubham", "Dwivedi"} // one way of creating instance
	// this way isn't good as order is important

	babloo := person{firstName: "Babloo", lastName: "Rajput"} // better way

	var deepak person // when you don't assign values to fields in declaration >
	// Go will assing "zero" values to fields (depending on type of field)
	// eg. string "", int 0, float 0, bool false

	// fmt.Println(shubham)
	fmt.Println(babloo)
	fmt.Println(deepak) // Note: field values won't be nil, null or undefined like other languages.

	fmt.Printf("%+v", babloo) // %+v will print all field name and values.

	babloo.firstName = "Rishi" // values can be updated like this.
	fmt.Println(babloo)

	shubham := person{ // declaring with nested structs
		firstName: "Shubham",
		lastName:  "Dwivedi",
		contactInfo: contactInfo{
			email:   "shubham@mail.com",
			zipCode: 110022,
		},
	}

	shubham.print()

	shubham.updateName("Shubh") // Will have no effect (passed by value)
	shubham.print()

	shubhamPointer := &shubham              // & means address of shubham (pointer stores the address)
	shubhamPointer.updateNamePntr("shubhh") // will have effect (passed by reference)
	// NOTE: shubhamPointer is a pointer of type shubham (pointers also have type)
	shubham.print()

	var shubhPointer *person = &shubham // long syntax more clear
	shubhPointer.updateNamePntr("shubham")

	// NOTE: Go allows shortcut of calling updateNamePntr directly for shubham (and provice pointer to shubham automatically)
	shubham.updateNamePntr("shubhDirect") // This may seem like a type mismatch
	// updateNamePntr will automatically receive &shubham (Go does this implicitly)
	shubham.print()

	// Go also have Arrays but are rarely used due to fixed size.
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice) // will print [Bye There How Are You]
	// Go is still passing by value for slice, the value is updated because :=
	// Slice is just a data structure (based on Array) that contains
	// 1. Length (Current),
	// 2. Capacity (Max),
	// 3. Pointer to an Array

	// When Go passes by value, it creates a copy of the Slice,
	// The Copy still POINTS to the SAME array.

	// These type of data structures that points to another data structure in memory
	// Are called "Reference Types".

	/*  Reference Types in GO:
	slices, maps, channels, pointers, functions
	*/

	/*	Value Types in GO:
		int, float, string, bool, structs
	*/

}

// Receiver function for person
func (p person) print() {
	fmt.Printf("%+v", p)
}

// NOTE: By default Go passes by value

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName // will have no effect on actual p
}

// &XX means address of XX. *YY means value at address YY. (YY = &XX is called a pointer (basically memory address))
// *TYPE means a pointer of type TYPE (eg *person or *string)

func (pointerToPerson *person) updateNamePntr(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
