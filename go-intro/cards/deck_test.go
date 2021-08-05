package main

import "testing"

func TestNewDeck(t *testing.T) { // Tests starts with UPPERCASE
	// t is the test handler
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d)) // use t.Error() for not formatting
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card of King of Diamonds but got %v", d[len(d)-1])
	}
}

// to test run > go test
// go env -w GO111MODULE=auto
