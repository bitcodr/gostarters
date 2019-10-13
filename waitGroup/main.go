package main

import (
	"fmt"
	"sync"
)

func increase(x int, wg *sync.WaitGroup) {
	if x%2 == 0 {
		fmt.Printf("%d is even\n", x)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increase(i, &wg)
	}

	wg.Wait()
}
