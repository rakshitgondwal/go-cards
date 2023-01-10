package main

import "testing"

func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 9 {
		t.Errorf("Wrong length of the deck. Expected 9 but got %v", len(d))
	}
}