package main

import (
	"fmt"
	"sync"
)

func increse(number *int, wg *sync.WaitGroup) {
	*number++
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	var x int = 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go increse(&x, &wg)
	}

	wg.Wait()

	fmt.Println(x)
}
