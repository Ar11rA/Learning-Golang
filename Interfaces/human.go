package main

import (
	"fmt"
)

type singer interface {
	sing()
}

type writer interface {
	write()
}

type human struct {
	name    string
	address string
}

type onlySinger struct {
	name    string
	address string
}

func (h *human) sing() {
	fmt.Printf("%s is singing\n", h.name)
}

func (h *human) write() {
	fmt.Printf("%s is writing\n", h.name)
}

func (o *onlySinger) sing() {
	fmt.Printf("%s is singing\n", o.name)
}

func print(s singer) {
	s.sing()
}

func main() {
	p := &human{
		name:    "p",
		address: "a1",
	}
	print(p)
	p.write()
	s := &onlySinger{
		name:    "s",
		address: "a2",
	}
	print(s)
}
