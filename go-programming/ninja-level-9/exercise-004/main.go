package main

import (
	"fmt"
	"sync"
)

func main() {
	const gsCount = 100
	counter := 0

	var wg sync.WaitGroup
	wg.Add(gsCount)

	var m sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
