package main

import (
	"fmt"
)

func main() {
	chanSolvedWithAnonymousFunction()
	chanSolvedWithBuffer()
}

func chanSolvedWithAnonymousFunction() {
	c := make(chan string)

	go func() {
		c <- "Channel solved with anonymous function"
	}()

	fmt.Println(<-c)
}

func chanSolvedWithBuffer() {
	c := make(chan string, 1)

	c <- "Channel solved with buffer"

	fmt.Println(<-c)
}
