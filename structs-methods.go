package main 

import "fmt" 

// Go's structs are typed collection of fields. Useful for grouping data together to form records.

type person struct { 
	name string 
	age int
}

// newPerson constructs a new person struct with the given name. 
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42 
	return &p 
	// You can safely return a pointer to local variable as local variable will survive the scope of the function.
}


type rect struct {
	width, height int 
}

// Go supports "methods" defined on struct types. 

// (r *rect) or (r rect) is the context of the method (ie: reciever type) not the parameters type
func (r *rect) area() int { // here receiver type is *rect (pointer to a rect struct)
	return r.width * r.height
}
func (r rect) perim() int { // methods can be defined with either pointer or value receiver types.
	return 2*r.width + 2*r.height
} 

func main() {
	fmt.Println(person{"Shubham", 20}) // {Shubham 20}
	// person{"Shubham", 2} is syntax to create a new person.

	fmt.Println(person{name: "Bharat", age: 30}) // {Bharat 30}
	// fields can be named when initializing a struc.

	fmt.Println(person{name: "Deepak"}) // {Deepak 0}
	// omitted fields will be zero-valued

	fmt.Println(&person{name: "Shubh", age: 25}) // &{Shubh 25}
	// an & prefix yields a pointer to the struct.

	fmt.Println(newPerson("Dubey")) // &{Dubey 42}
	// it's good practise to encapsulate new struct creation in constructor functions.

	s := person{name: "Shubham Dwivedi", age: 26}
	fmt.Println(s.name) // Access fields this way. 

	sp := &s 
	fmt.Println(sp.age) // you can use dots with struct pointers. the pointers are automatically dereferenced.

	sp.age = 51 // Structs are mutable (values can be changed later this way.)
	fmt.Println(sp.age)


	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r // Go automatically handles conversion between values and pointers for method calls
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
	
	// Use a pointer receiver type to avoid copying on method calls or to allow method to mutate the receiving struct. 
	// Practise this. ????
}