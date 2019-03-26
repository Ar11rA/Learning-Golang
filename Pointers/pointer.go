package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) PrintFullNameValue() {
	fmt.Printf("PrintFullNameValue: address of p is %p\n", &p)
	p.FirstName = "def"
}

func SetFullName(p *Person) {
	fmt.Printf("SetFullName: p is %p\n", p)
	fmt.Printf("SetFullName: p is %s\n", p.FirstName)
	p.FirstName = "abc"
}

func main() {
	p := Person{
		"John",
		"Doe",
	}
	fmt.Printf("address of p: %p\n", &p)
	SetFullName(&p)
	p.PrintFullNameValue()
	fmt.Printf("%s", p.FirstName)
}
