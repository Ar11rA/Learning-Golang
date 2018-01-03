package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %d ", len(d))
	}
}

func TestGetSuits(t *testing.T) {
	d := getSuits()
	if len(d) != 4 {
		t.Errorf("Expected length of 4 but got %d ", len(d))
	}
}

func TestGetValues(t *testing.T) {
	d := getValues()
	if len(d) != 13 {
		t.Errorf("Expected length of 4 but got %d ", len(d))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	h1, h2 := deal(d, 20)
	if len(h1) != 20 {
		t.Errorf("Expected length of 20 but got %d ", len(d))
	}
	if len(h2) != 32 {
		t.Errorf("Expected length of 32 but got %d ", len(d))
	}
}
