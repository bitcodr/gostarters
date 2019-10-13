package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeCake(cakeType string, number int, in chan<- string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Microsecond)
		in <- cakeType
	}
	close(in)
}

func main() {
	ch1 := make(chan string, 4)
	ch2 := make(chan string, 3)

	go makeCake("cream", 4, ch1)
	go makeCake("black", 3, ch2)

	moreCream := true
	moreBlack := true
	var cake string

	for moreCream || moreBlack {
		select {
		case cake, moreCream = <-ch1:
			if moreCream {
				fmt.Printf("cake %s\n", cake)
			}
		case cake, moreBlack = <-ch2:
			if moreBlack {
				fmt.Printf("cake %s\n", cake)
			}
		}
	}

}
