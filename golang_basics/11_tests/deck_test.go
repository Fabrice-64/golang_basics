package main

import (
	"os"
	"testing"
)

func TestNewDecl(t *testing.T) {
	d := newDeck()
	if len(d) != 40 {
		t.Errorf("Expected Deck of 40, but got %v", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts, but got %v", d[0])
	}
	if d[len(d)-1] != "Jack of Spades" {
		t.Errorf("Expected Jack of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := readFile("_decktesting")
	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40 cards, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
