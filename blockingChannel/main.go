package main

import (
	"fmt"
	"net/http"
	"sync"
)

func worker(in <-chan string, wg *sync.WaitGroup) {
	for {
		url := <-in
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("url is %s with statusCode %d\n", url, res.StatusCode)
		}
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup

	work := make(chan string, 1024)

	numberOfWorker := 1000

	for i := 0; i < numberOfWorker; i++ {
		go worker(work, &wg)
	}

	urls := [6]string{
		"https://www.google.com/",
		"https://twitter.com",
		"https://us.yahoo.com/",
		"https://www.youtube.com/",
		"https://www.digikala.com/",
		"https://farinmedia.ir/",
	}

	for i := 0; i < 100; i++ {
		for _, v := range urls {
			wg.Add(1)
			work <- v
		}
	}

	wg.Wait()
}
