//+build wireinject

package main

import "github.com/google/wire"

// InitializeEvent ...
func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}
