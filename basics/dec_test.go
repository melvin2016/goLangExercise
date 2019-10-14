package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected Deck to be 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first item of deck to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first item of deck to be Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestNewDeckAndSaveAndReadfromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.writeToFile("_decktesting")
	loadedDeck := readFromFile("_decktesting")
	if len(loadedDeck) != 20 {
		t.Errorf("Expected Deck to be 20, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
