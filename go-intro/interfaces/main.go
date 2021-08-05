package main

import "fmt"

type bot interface { // any type with receiver function getGreeting() will be member of interface bot.
	getGreeting() string
	// interface can also have other interfaces (see http)
}

type englishBot struct{} // structs can be declared with zero properties
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// function overloading is not supported in GO.
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// will give error that printGreeting is already declared.
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) { // takes an interface as argument ?
	fmt.Println(b.getGreeting())
} // This way we don't need to declare separate printGreeting for both englishBot and spanishBot
/*** POLYMORPHISM ^^ ***/

// a receiver function for type englishBot
func (englishBot) getGreeting() string {
	// assuming getGreeting have very differenct custom logic for englishBoth and spanishBot
	return "Hello Friend!"
}

func (spanishBot) getGreeting() string {
	// since we do not use sb in here we can omit it from (sb spanishBot) receiver type.
	return "Hola Amigo!"
}

// ########## Concrete vs Interface Types ###########

// A Concrete Type is something we can create a value out of directly, and then access it or change it or create copies of it.
// map, struct, int, string, englishBot, spanishBot are all Concrete Types

// An Interface Type is something we cannot create a value out of directly.
// bot is an Interface Type.

// ##################################################

// Interfaces in Go are not "generic" types like in other languages. Go does not have generic types.

// Go will "implicitly" associate types (such as englishBot or spanishBot) to an Interface type (bot).
// We don't need to manually declare association.

// ###### INTERFACES CONTINUED IN http #########
