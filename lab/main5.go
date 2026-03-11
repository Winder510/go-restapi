package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 20)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 20; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for v := range jobs {
				fmt.Printf("Worker %d processed job %d\n", i, v)
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}
