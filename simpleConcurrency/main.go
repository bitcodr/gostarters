package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("hello world")
	time.Sleep(1 * time.Second)
}
