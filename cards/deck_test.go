package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	deck := newDeck()

	if len(deck) != 20 {
		t.Errorf("Expected length 20 but found %v", len(deck))
	}
	if deck[0] != "Heart of Ace" {
		t.Errorf("Expected Heart of Ace, found: %v", deck[0])
	}

	if deck[len(deck)-1] != "Clove of Four" {
		t.Errorf("Expected Clove of Four, found: %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Failed to load or Save deck to file %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
