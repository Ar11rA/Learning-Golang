package main

import "fmt"

func main() {
	// initialize new deck
	cards := newDeck()
	// deal a hand
	firstHand, secondHand := deal(cards, 18)
	// print above decks
	cards.print()
	firstHand.print()
	secondHand.print()
	// convert deck type to string
	fmt.Println(cards.toString())
	// write to file
	firstHand.saveToFile("first_hands.txt")
	// read from file
	readDeck("first_hands.txt").print()
	// shuffling
	firstHand.shuffle()
	firstHand.print()
	firstHand.shuffle()
	firstHand.print()
}
