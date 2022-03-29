package main

import "fmt"

func main() {
	a := incrementer(1)
	b := incrementer(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementer(i int) func() int {
	return func() int {
		i++
		return i
	}
}
