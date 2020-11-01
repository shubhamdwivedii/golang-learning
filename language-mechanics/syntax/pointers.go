package main 
import ("fmt")


/****** Everything is about pass by value ******/

/* Pointers share values across the program boundaries. There are several types of program boundaries. 
   The most common one is between function calls. (We also have boundaries b/w Goroutines) */

/* When this program starts up, the runtime creates a Goroutine. 
   Every Goroutine is a separate path of execution that contains instructions that needed to be executed by the machine. 
   We can think of Goroutines as lightweight threads. This program as only 1 Goroutine: the main Goroutine. */


/* Every Goroutine is given a block of memory, called the stack.
   The stack memory in Go starts out at 2K (it is very small) which can change over time.
   Every time a function is called, a piece of stack is used to help that function run.
   The growing direction of the stack is downward. */

/* Every Function is given a stack frame, memory execution of a function. 
   The size of every stack frame is known at compile time. No value can be placed on a stack unless the compiler KNOWS its size ahead of time. 
   If we don't know the size of something at compile time, it has to be on the heap. */

/* Zero value enables us to initialize every stack frame that we take. Stacks are self cleaning. 
   We clean our stack on the way down ????. Every time we make a function, zero value initialization cleans the stack frame.
   We leave that memory on the way up because we don't know if we would need that again. ???? */ 

/****** Pass by value ******/


func main() {
	count := 10 
	
	// To get the address of the value, we use & 
	fmt.Println("count:\tValue Of [", count, "]\tAddr Of [", &count, "]")

	// Pass the "value of" count. 
	increment1(count)

	fmt.Println("count:\tValue Of [", count, "]\tAdd Of [", &count, "]")
	// Nothing has changed

	// Pass the "address of" count. (this is still considered pass by value, not by reference because the address itself is a value)
	increment2(&count)

	fmt.Println("count:\tValue Of [", count, "]\tAdd Of [", &count, "]")
	// Value has changed (incremented by 1)
}

func increment1(inc int) {
	inc++  // Incrementing the "value of" inc.
	fmt.Println("inc1:\tValue Of [", inc, "]\tAddr Of [", &inc, "]")
}

/* increment2 declares count as a pointer variable whose value is always an address and points to values of type int. 
   Every type that is declared, whether you declare or it is predeclared, you get for free a pointer. */

func increment2(inc *int) { // The * here is not an operator. It is part of the type name. 
	// Increment the "value of" count that the "pointer points to". 
	*inc++ // The * operator tells us the value of the address which pointer points to.
	fmt.Println("inc2:\tValue Of [", inc, "]\tAddr Of [", &inc, "]\tValue Points To [", *inc, "]")
}
