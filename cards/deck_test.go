package main

import (
    "testing"
    "os"
)

func TestNewDeck(t *testing.T) {
    d := newDeck()

    if len(d) !=  16 {
        t.Errorf("Expected deck length of 16, but got %v", len(d))
    }


    if d[0] != "Ace of Spades" {
        t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
    }

    if d[len(d) - 1] != "Four of Club" {
        t.Errorf("Expected first card of Four of Club, but got %v", d[len(d) - 1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")

    d := newDeck()
    d.saveToFile("_decktesting")

    loadedDeck := newDeckFromFile("_decktesting")

    if  len(loadedDeck) != 16 {
        t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
    }

    os.Remove("_decktesting")
}
