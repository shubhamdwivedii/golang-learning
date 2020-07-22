// Go-Lang Basics - Shubham Dwivedi. 

// to build: "go build cheatsheet.go"
// to run: "./cheatsheet"

package main // tell which package this go file is part of 
import (
	"fmt"
	"math"
	"time"
) 

const S string = "constant_string" // for Constants 

func main() {
	// Hello World ###############
	fmt.Printf("hello, world\n") 


	// Values ####################
	fmt.Println("go" + "lang") // strings 

	fmt.Println("1+1+2 =", 1+1+2) // integers
	fmt.Println("7.0/3.0 =", 7.0/3.0) // floats
 
	fmt.Println(true && false) // false 
	fmt.Println(true || false) // true 
	fmt.Println(!true) // false


	// Variables ##############

	// In Go, variables are explicitly declared and used by the compiler to eg. check type-correctness of function calls. 
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1,2 // var declares 1 or more variables.  
	fmt.Println(b,c)

	var d = true 
	fmt.Println(d) // Go will INFER the type of initialized variables. 

	var e int // varables declared without initialization are ZERO-VALUED. eg. the zero-value for an int is 0. 
	fmt.Println(e)

	f := "apple" // := is shorthand syntax for declaring and initializing a variable. 
	var g string = "apple" // should be same as above
	fmt.Println(f, g)


	first, second := 42, "second" // this is also valid initialization
	fmt.Println("first", first)
	fmt.Println("second", second)


	// Constants ################
	fmt.Println(S) // constant string declared at top. 
	const n = 500000000 // "const" keyword declares sontsant values, can be used anywhere a "var" can be used. 
	const d2 = 3e20/n // cosnt expressions perform arithmetic with arbitrary precision.   
	fmt.Println(d2)
	
	// A numeric constant has no type until it's given one, such as by an explicit conversion.
	fmt.Println(int64(d2))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call.
	fmt.Println(math.Sin(n)) // here math.Sin() expects a float64
	

	// For-loop #################

	// "for" is Go's only looping construct.

	i := 1 
	for i <= 3 {
		fmt.Println("for-",i)
		i = i + 1
	}

	for j := 7; j <=9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop with break")
		break 
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // directly skips to next iteration
		}
		fmt.Println("Odd? ",n)
	}


	// If/Else ####################

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
 
	// if in Go also supports this for-like syntax with variable declaration
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// There is no "ternary if operator" in go ( 4 < 2 ? "this" : "that"), use full if statement even for most basic conditions. 


	// Switch ########################

	k := 2 
	fmt.Print("Write ", k, " as ") // Print does skips to next line 
	
	switch k { 
	case 1: 
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3: 
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:  // multiple expressions in same case are allowed in Go.
		fmt.Println("It's the weekend :-)")
	default:  // default case is optional as usual.
		fmt.Println("It's a weekday :-(")
	}

	// switch without expression is alternative to express if/else logic. 
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default: 
		fmt.Println("It's after noon")
	}


	// a "type switch" compares types instead of values. 
	whatType := func(i interface{}) { // whatType is a function that takes and interface shaped {} as argument (basically any type)
		switch t := i.(type) { // t is initialized to the type of i (argument given)
		case bool: 
			fmt.Println("I'm a bool")
		case int: 
			fmt.Println("I'm an int")
		default: 
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatType(true)
	whatType(12)
	whatType("hello")


	// Arrays #######################

	// In Go, Arrays have specific length.
	

	var arr [5]int 
	fmt.Println("emp:", arr) // will print [0,0,0,0,0] as int's zero-value is 0. 
	// Type of elements and Length are both part of Array's type. 

	arr[4] = 100 
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	// use len() to get Length of an Array. len() is built-in function in Go.
	fmt.Println("len:", len(arr))

	brr := [5]int{1,2,3,4,5} // initializing and Array with declaration
	fmt.Println("ini:", brr)

	var twoD [2][3]int  // a 2-D array (matrix)
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ { // eg. of a nested loop. bc C ki yaad aagai :')
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// Slices are used more often than Arrays in Go. 


	// Slices ##############################

	// In Go, Slices are a key data type, giving more powerful interface to sequences than arrays.

	slc := make([]string, 3) // unlike arrays, slices are typed only by the elements they contain (not the number of elements)
	// make() is a built-in funcion in Go, used here to create a non-zero length slice.
	fmt.Println("empty:", slc) // empty: [  ] as zero-value of strings is ""
	fmt.Println("length:", len(slc)) // 3

	slc[0] = "a"
	slc[1] = "b"
	slc[2] = "c"
	fmt.Println("updated:", slc) // updated: [a b c]
	fmt.Println("lastElement:", slc[2])

	// slice also supports append 
	slc = append(slc, "d") // append returns a NEW slice rather than updating the one its given.
	slc = append(slc, "e", "f") // multiple values can be appended at once. 
	fmt.Println("appeded:", slc) //appeded: [a b c d e f]

	// use copy() to copy a slice into new one. 
	cpy := make([]string, len(slc))
	copy(cpy, slc) // copy(target, source)
	fmt.Println("copied:", cpy)

	sgmnt1 := slc[2:5] // ":" is the "slice" operator [from:to]. this will return a slice [slc[2], slc[3], slc[4]] (just like python)
	sgmnt2 := slc[:3] // this will slice from 0 upto 3 (excluding 3)
	sgmnt3 := slc[2:] // this will slice from 2 (including 2) upto end (ie len(slc)).

	fmt.Println("Segments:", sgmnt1, sgmnt2, sgmnt3)

	initSlc := []string{"g", "h", "i"} // initializing a slice (not make() not needed here)
	fmt.Println("initialized:", initSlc)


	twoDslc := make([][]int, 3) // a 2D slice (matrix)
	for i := 0; i < 3; i++ {
		innerLen := i + 1 
		twoDslc[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDslc[i][j] = i + j 
		}
	}
	fmt.Println("2d Slice:", twoDslc) // [[0] [1 2] [2 3 4]] length of inner slices can vary unlike 2d arrays.


	// Maps ##########################

	// Maps in Go are similar to hashes or dicts (dictonary in python or object in javaScript)

	m := make(map[string]int) // creating an empty map (using built-in make) make(map[key-type]value-type)
	
	m["k1"] = 7  // setting key/value pairs 
	m["k2"] = 13

	fmt.Println("map:", m) // map: map[k1:7 k2:13]
 
	v1 := m["k1"] // accessing values 
	fmt.Println("v1:", v1) // 7
	fmt.Println("len:", len(m)) // 2 - number of key/value pairs 

	delete(m, "k2") // to remove a key/value pair use built-in delete() 
	fmt.Println("map:", m)

	_, present := m["k2"] // optional second return value indicates if a key was present or not in a map. 
	fmt.Println("present?:", present)

	newMap := map[string]int{"foo": 1, "bar": 2} // declaring and initializing 
	fmt.Println("map:", newMap)



	// Range #############################

	// use "range" to iterate over elements in a variety of data structures. 


	nums := []int{2,3,4} // declaring a slice 
	sum := 0 
	for _, num := range nums { // kind of like forEach in javaScript or for x in list: in python
		sum += num
	}

	fmt.Println("Sum of Slice:", sum)

	for i, num := range nums { // range on arrays and slices provides both the index and values for elements
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, val := range kvs {
		fmt.Printf("%s -> %s\n", key, val)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range can be used to iterate over string, but it iterates over Unicode code points.
	for i, c := range "go" {  // index and character in string
		fmt.Println(i, c)
	} /* Prints 
	    	0 103
			1 111  ??? */
			
	
	// NEXT - functions.go

}