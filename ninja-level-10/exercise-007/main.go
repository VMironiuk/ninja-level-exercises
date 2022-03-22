package main

import "fmt"

func main() {
	const goroutinesCount = 10

	for i := 0; i < goroutinesCount; i++ {
		c := make(chan int)

		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
			close(c)
		}()

		for v := range c {
			fmt.Println(v)
		}
	}

	fmt.Println("About to exit...")
}
