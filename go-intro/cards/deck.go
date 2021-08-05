package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/****** Custom Types in Go *******/

// Creating a new type 'deck' which is a 'slice of strings'
type deck []string

/****** Receiver functions ********/

func (d deck) print() { // a receive function that expects a deck type argument
	for i, card := range d { // d here is used like 'this' or 'self' from other languages
		fmt.Println(i, card)
	}
} // any variable of type 'deck' can call this function on itself. Example d.print() where d is a deck

// NOTE: receiver is before function name and return type is after.

func newDeck() deck { // return type is deck
	cards := deck{}
	suits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

/****** Range syntax (Slices) and Multiple return values *******/

func deal(d deck, handSize int) (deck, deck) { // multiple return types
	// slc[startIndexIncluding:uptoNotIncluding] like python
	return d[:handSize], d[handSize:] // hand and deck(remaining)
}

/****** Type conversion *******/

// deck > []string > string > []byte

func (d deck) toString() string {
	convertedDeck := []string(d) // converting deck to a []string
	// join takes a []string and a separator string.
	return strings.Join(convertedDeck, ",") // return strings.Join([]string(d), ",")
}

/****** Writing to file *******/

func (d deck) saveToFile(filename string) error { // returns an error sometimes
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// WireFile takes filename, []byte and permission (perm os.FileMode) and returns error
}

/***** Reading from file ******/

func readFromFile(filename string) deck {
	byteslice, err := ioutil.ReadFile(filename)

	/**** Basic Error handling ****/
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) // any value other that 0 indicates something when wrong.
	}

	joinedString := string(byteslice) // converting []byte to string

	cardstrings := strings.Split(joinedString, ",") // Split is like javascript split.

	return deck(cardstrings) // converting []string to deck.
}

/***** delete file ********/
func removeFile(filename string) error {
	return os.Remove(filename)
}

/****** Random numbers *****/

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // source acts as seed for psuedo  randomization.
	r := rand.New(source)                           // r is randomizer
	for i := range d {                              // we only need index here
		newPosition := r.Intn(len(d) - 1)           // will generate a random number between 0 and the argument
		d[i], d[newPosition] = d[newPosition], d[i] // you can do this signle line swap in Go.
	}
}
