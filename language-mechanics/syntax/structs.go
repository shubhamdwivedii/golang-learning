package main 
import ("fmt")

/******* example struct ********/

// example represents a type with different fields. 

type example struct {
	flag bool 
	counter int16 
	pi float32
}

/***** Declare and initialize *****/

func main() {
	var e1 example // declaring a variable of exmaple (and set to its zero value)

	fmt.Printf("%+v\n", e1) // {flag:false counter:0 pi:0}

	literal()
	samename()
}


/* How much memory do we allocate for example ? 
1 byte for bool (flag), 2 bytes for int16 (counter) and 4 bytes for float32 (pi), total = 7 bytes.
However actual size is 8 bytes. This is due to padding and alignment.  

The padding byte is sitting between the bool and the int16. The reason for this is alignment. 
The idea of alignment: It is more efficient for this piece of hardware to read memory on its alignment boundary. 

Depending on size of a particular value, Go determines the alignment we need. Therfore: 

Rule 1: Every 2 bytes value must follow a 2 bytes boundary. Since bool is only 1 byte and starts at address 0, 
then the next int16 must start on address 2. The byte at the address that get skipped over becomes a 1-byte padding. 
Similarly, if it is a 4-byte value then we will have a 3-byte padding value. 

The largest field represents the padding for the entire struct. We need to mimimize the amount of padding as much as possible. Therefore: 

Rule 2: Always lay out the field from highest to smallest. This will push any padding down to the bottom. 
In this case, the entire struct has to follow a 8 bytes value because int64 is 8 bytes. */ 

type example2 struct {
	counter int64 
	pi float32 
	flag bool
}

func literal() {
	// Declaring a variable of type example2 and initializing using a "struct literal"
	e2 := example2 {
		flag: true, 
		counter: 10,
		pi: 3.141592, // trailing comma is required in go
	}

	fmt.Println("Flag", e2.flag) // Println() automaically adds a "\n" 
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}


/****** Name type vs anonymous type *******/

/* If we two identical structs with different names, we can't assing one to another. 
for eg. example2 and example1 are identical structs: */

type example1 struct {
	counter int64 
	pi float32 
	flag bool
}

func samename() {
	e3 := example1 {
		flag: true, 
		counter: 10, 
		pi: 3.141592,
	}

	e4 := example2 {
		flag: true, 
		counter: 10, 
		pi: 3.141592,
	}

	var e5 example1

	// e5 = e4 // This is not allowed (as e5 and e4 have different types)
	
	e5 = e3 // This IS allowed (as e5 and e3 have save types)

	// we can explicitly perform a conversion
	e5 = example1(e4) // This is allowed

	fmt.Printf("%+v\n", e5)
}
