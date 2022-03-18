package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	l float64
	w float64
}

func (s square) area() float64 {
	return s.l * s.w
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{
		l: 20,
		w: 30,
	}
	c := circle{
		r: 20,
	}
	info(s)
	info(c)
}
