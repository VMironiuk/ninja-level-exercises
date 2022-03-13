package main

import "fmt"

const (
	a int     = 42
	b float64 = 3.14
	c string  = "James Bond"
)

const (
	x = 36
	y = 36.6
	z = "John Doe"
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(x, y, z)
}
