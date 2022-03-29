package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 0
}

func bar() (int, string) {
	return 42, "Meaning of live"
}
