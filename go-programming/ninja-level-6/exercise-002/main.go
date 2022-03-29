package main

import "fmt"

func main() {
	fmt.Println(foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...))
	fmt.Println(bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func foo(n ...int) (total int) {
	for _, v := range n {
		total += v
	}
	return total
}

func bar(s []int) (total int) {
	for _, v := range s {
		total += v
	}
	return total
}
