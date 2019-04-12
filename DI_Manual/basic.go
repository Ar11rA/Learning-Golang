package main

import "fmt"

// Message ...
type Message string

// NewMessage ...
func NewMessage(msg string) Message {
	return Message(msg)
}

// Greeter ...
type Greeter struct {
	Message Message // <- adding a Message field
}

// NewGreeter ...
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent ...
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Event ...
type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

// Greet ...
func (g Greeter) Greet() Message {
	return g.Message
}

// Start ...
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	message := NewMessage("Hola")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}
