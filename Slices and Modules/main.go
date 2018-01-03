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
	//helpers
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a, arr := pop(arr, 4)
	fmt.Println("Popped   : ", a, arr)
	a, arr = shift(arr)
	fmt.Println("Shifted  : ", a, arr)
	arr = unshift(arr, 10)
	fmt.Println("Unshifted: ", arr)
	arr = reverse(arr)
	fmt.Println("Reversed : ", arr)
	arr = maps(arr, func(r int) int {
		return r * 2
	})
	fmt.Println("Mapped   : ", arr)
	a = reduce(arr, func(prev int, now int) int {
		return prev + now
	})
	fmt.Println("Reduced  : ", a)
}
