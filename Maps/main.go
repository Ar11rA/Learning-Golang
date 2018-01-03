package main

import (
	"fmt"
)

type hash map[string]int

func main() {
	// colors := map[string]int
	colors := hash{
		"red":    0,
		"green":  1,
		"blue":   2,
		"yellow": 3,
	}
	colors.print()
	delete(colors, "yellow")
	colors.print()
}

func (h hash) print() {
	for color, number := range h {
		fmt.Println(color, number)
	}
}
