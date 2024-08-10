package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected 16, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades, got %v", d[0])
	}

	if d[len(d)-1] != "Three of Clubs" {
		t.Errorf("Expected Three of Clubs, got %v", d[len(d)-1])
	}
}