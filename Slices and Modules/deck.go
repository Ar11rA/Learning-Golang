package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func newDeck() deck {
	suits := getSuits()
	values := getValues()
	cards := deck{}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func getSuits() []string {
	return []string{"Spades", "Hearts", "Clubs", "Diamonds"}
}

func getValues() []string {
	return []string{"Ace", "King", "Queen", "Jack", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
}

func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func readDeck(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("Error", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[newPosition], d[i] = d[i], d[newPosition]
	}
}
