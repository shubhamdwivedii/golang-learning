package main

import "fmt"

type bot interface { // any type with receiver function getGreeting() will be member of interface bot.
	getGreeting() string
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
