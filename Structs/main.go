package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	email     string
	contact   contactInfo
}

type contactInfo struct {
	street  string
	city    string
	state   string
	country string
	code    int
	phone   string
}

func (p person) print() {
	fmt.Printf("user: %+v", p)
	fmt.Println("")
}

func (p *person) updateEmail(email string) {
	p.email = email
}

func main() {
	user := person{
		firstName: "alex",
		lastName:  "summers",
		email:     "alex.summers@mail.com",
		contact: contactInfo{
			street:  "brooklyn",
			city:    "new york city",
			state:   "new york",
			country: "usa",
			code:    1,
			phone:   "555- 4444",
		},
	}
	user.print()
	user.updateEmail("alex@summers.com")
	user.print()
}
