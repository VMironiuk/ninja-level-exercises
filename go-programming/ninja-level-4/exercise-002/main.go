package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", s)
}
