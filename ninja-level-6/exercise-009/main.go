package main

import "fmt"

func main() {
	fmt.Println(doSomeMath(sum, 2, 5))
	fmt.Println(doSomeMath(div, 8, 3))
}

func doSomeMath(f func(int, int) int, a, b int) int {
	return f(a, b)
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a - b
}
