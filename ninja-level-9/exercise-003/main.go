package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs count:", runtime.NumCPU())

	var wg sync.WaitGroup
	const gsCount = 100
	counter := 0

	wg.Add(gsCount)

	for i := 0; i < gsCount; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
