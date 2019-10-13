package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello World 1"
	}()

	fmt.Println(<-ch)
}
