package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const gs = 100
	var counter int64

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
