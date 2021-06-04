package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 30, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Club" {
		t.Errorf("Expected last card of Ace of Club, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // remove file

	d := newDeck()               // making new deck
	d.saveToFile("_decktesting") // save file to deck

	loadedDeck := newDeckFromFile("_decktesting") //we call the file in this case "loaded deck"

	if len(loadedDeck) != 56 {
		t.Errorf("Expected 56 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting") // remove file
}
