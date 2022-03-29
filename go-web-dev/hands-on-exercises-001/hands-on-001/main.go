package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Println("The area is:", s.area())
}

func main() {
	s := square{side: 3}
	c := circle{radius: 2}
	info(s)
	info(c)
}
