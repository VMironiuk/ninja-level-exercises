package main

import "fmt"

func main() {
	var b bool

	b = 42 == 42
	fmt.Printf("42 == 42: %v\n", b)

	b = 42 <= 42
	fmt.Printf("42 <= 42: %v\n", b)

	b = 37 >= 42
	fmt.Printf("37 >= 42: %v\n", b)

	b = 42 != 42
	fmt.Printf("42 != 42: %v\n", b)

	b = 37 < 42
	fmt.Printf("37 < 42: %v\n", b)

	b = 37 > 42
	fmt.Printf("37 > 42: %v\n", b)
}
