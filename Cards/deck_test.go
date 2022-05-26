package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	dck := newDeck()

	if len(dck) != 48 {
		t.Errorf("Expected deck length of 48, but got %v", len(dck))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	dck := newDeck()

	dck.saveToFile("_decktesting")

	loadedDck := newDeckFromFile("_decktesting")

	if len(loadedDck) != 48 {
		t.Errorf("Expected deck length of 48, but got %v", len(loadedDck))
	}

	os.Remove("_decktesting")
}
