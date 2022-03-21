package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("new goroutine (1)")
		wg.Done()
	}()
	go func() {
		fmt.Println("new goroutine (2)")
		wg.Done()
	}()

	fmt.Println("main goroutine")

	wg.Wait()
}
