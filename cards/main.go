package main

import "fmt"

func main() {

	/****** Variable Declaration ******/
	var card string = "Ace of Spades" // declaring and assigning value to a variable
	fmt.Println(card)

	// type string can also be inferred from value assigned.
	card2 := "Ace of Spades 2nd" // alternate way of declaring and assigning a value
	fmt.Println(card2)
	card2 = "Five of Diamonds" // don't use := for reassingment
	fmt.Println(card2)

	/****** Functions and return types ******/
	card3 := newCard()
	fmt.Println(card3)

	/****** Arrays Slices and For loops  *******/

	// Array = Fixed length
	// Slice = Can grow or shrink

	// declaring a Slice of type string
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades") // appending to end of Slice
	// append returns a new slice (does not modify existing). We reassign new slice to cards here.

	for i, card := range cards { // for loop in Go (i is index (optional?))
		fmt.Println(i, card)
	} // range is used to iterate over a Slice. := is used to reassign i and card each time.

	/****** Go is not a OOP language ******/
	fmt.Println("******************")

	/****** Custom Types in Go *******/
	// See deck.go now >>>>>>>>>>>
	cards2 := deck{"Ace of Diamonds", newCard()} // deck is custom type
	cards2 = append(cards2, "Six of Spades")
	cards2.print() // A receiver function
	// To run > go run main.go deck.go

	fmt.Println("********************")

	/****** Slice ranges *******/
	cards3 := newDeck()
	fmt.Println(cards3)
	fmt.Println(len(cards3)) // len(slice) returns the length of slice
	hand, remainingDeck := deal(cards3, 7)

	hand.print()
	fmt.Println(remainingDeck) // not calling remaingingDeck.print() to avoid clutter in terminal

	/***** Type conversion *****/
	fmt.Println([]byte("Hello There")) // []byte is the type we want to convert to.
	// [72 101 108 108 111 32 84 104 101 114 101]

	fmt.Println(hand.toString()) // toString() is a receiver function in deck.go

	/***** Writing to file ******/
	hand.saveToFile("my_cards")

	/***** Reading from file ******/
	fmt.Println("===========")
	handNew := readFromFile("my_cards")
	handNew.print()

	/***** Remove file *****/
	removeFile("my_cards") // see deck.go

	/***** Random numbers *****/
	fmt.Println("**********")
	handNew.shuffle()
	handNew.print()
	fmt.Println("**********")
	handNew.shuffle()
	handNew.print()

}
func newCard() string { // string is the return type of function
	return "Four of Clubs"
}
