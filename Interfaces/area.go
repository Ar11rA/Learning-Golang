package main

import (
	"fmt"
	"math"
)

func area() {
	circle := Circle{
		radius: 10,
	}
	square := Square{
		length: 5,
	}
	rectangle := Rectangle{
		length:  10,
		breadth: 12,
	}
	getArea(circle)
	getArea(square)
	getArea(rectangle)
}

// Shape ....
type Shape interface {
	area() float64
}

// Circle ....
type Circle struct {
	radius float64
}

// Square ....
type Square struct {
	length float64
}

// Rectangle ....
type Rectangle struct {
	breadth float64
	length  float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.breadth * r.length
}

func (s Square) area() float64 {
	return s.length * s.length
}

func getArea(s Shape) {
	fmt.Println(s.area())
}
