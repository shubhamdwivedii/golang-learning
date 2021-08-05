package main // Every Go program is made up of packages.
// Programs always start running in package "main"

import ( // Grouped imports are called "factored" import statement (more clean)
	"fmt"
	"math"
	"math/cmplx"
	"math/rand" // By convention, the package name is same as last element of import path (eg. "rand" here)
	// "math/rand" package comprises of files that begin with statement 'package rand'
)

//import "math" // can also write multiple import statements

func main() {
	fmt.Println("My favourite number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// In Go, a name is exported if it begins with capital letter.
	// fmt.Println(math.pi) // will give error
	// When importing a package, you can only refer to its exported names.
	fmt.Println(math.Pi) // unexported names (starting with small letter) are not accessible from outside the package.

	fmt.Println(add(42, 13))
	variables()
	basicTypes()
	constants()
}

// In Go type comes afer the variable name
func add(x int, y int) int {
	return x + y
}

func swap(x, y int) (int, int) { // can omit type if argument is same type as next argument.
	return y, x // can return multiple types
}

func split(sum int) (x, y int) { // return values may be named
	x = sum * 4 / 9
	y = sum - x
	return // a return statement without arguments return the named return values (x & y here)
	// this return type are known as "Naked return" (avoid them in longer functions)
}

// "var" can be used to declare multiple variables
var c, python, javascript bool

// var declaration can include initializers, one per variable
var i, j int = 1, 2

// A var statement can be at package or function level.
func variables() {
	var i int
	fmt.Println(i, c, python, javascript) // 0, false, false, false

	var ruby, pearl, java = true, false, "no!" // variables of multiple types can be initialized together like this.
	fmt.Println(i, j, ruby, pearl, java)

	// use := for short initialization (declaration + assignment)
	k := 3 // type is decided implicitly based on the value assigned
	scala, dart, lua := true, false, "nope!"
	// := is only available inside a function.
	fmt.Println(k, scala, dart, lua)
}

// Basic Types ########

/* Basic Types in Go are:
bool
string
int int8 int16 int32 int64
uint uint8 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for uint32 (represents a Unicode code point)
float32 float64
complex64 complex128 */

func basicTypes() {
	var (
		toBe      bool       = false
		twoCube   int        = 2 << 3  // 2 times 2, 3 times (basically 2 to the power 3)
		cubeRoot9 int32      = 16 >> 3 // 16 diveded by 2, 3 times
		maxInt    uint64     = 1<<64 - 1
		z         complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", twoCube, twoCube)
	fmt.Printf("Type: %T Value: %v\n", cubeRoot9, cubeRoot9)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero Values ########

	// Variables declared without an explicit initial value are given their "zero" value
	var (
		i int     // zero value is 0
		f float64 // zero value is 0.0
		b bool    // zero value is false
		s string  // zero value is ""
	)

	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type Conversions ######

	var x, y int = 3, 4
	// use "T(v)" to convert "v" to type "T"
	var ff float64 = math.Sqrt(float64(x*x + y*y)) // math.Sqrt() only takes a float64
	var zz uint = uint(f)
	fmt.Println(x, y, ff, zz)

	// for numeric constants Type is inferred based on precision point.
	ii := 42           //int
	fff := 3.142       // float64
	gg := 0.867 + 0.5i // complex128

	fmt.Printf("ii is of type %T\n", ii)
	fmt.Println(fff, gg)
}

// Constants #####

// In Go constants can be character, string, boolen or numeric values
// Constants cannot be declared using :=

const Pi = 3.14 // can be declared outside functions
// use "const" instead of "var"

func constants() {
	const World = "Japan"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Numeric Constants ######

	// Numeric constants are high-precision values
	// An untyped constant takes the type needed by its context

	const (
		Big   = 1 << 100  // Creates a huge number by shifting a 1 bit left 100 places (In other words, a binary number that is 1 followed by 100 zeroes)
		Small = Big >> 99 // Shifting it right again 99 places, so we end up with 1<<1, or 2.
	)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// needInt(Big) // will give error: cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to needInt (overflows)

}

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

// NEXT >>> Flow Control
