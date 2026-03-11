package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jobChan := make(chan int)

	var wg sync.WaitGroup

	go func() {
		for _, v := range jobs {
			jobChan <- v
		}
		close(jobChan)
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range jobChan {
				fmt.Println(v * v)
			}
		}()
	}

	wg.Wait()
}
