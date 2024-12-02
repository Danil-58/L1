package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var result int64
	numbers := []int64{2, 4, 6, 8, 10}

	for _, v := range numbers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, v int64) {
			defer wg.Done()
			atomic.AddInt64(&result, v*v)
		}(&wg, v)
	}

	wg.Wait()
	fmt.Println("Result ", result)
}