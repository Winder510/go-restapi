package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.Sleep(time.Second)
		results <- job * job
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// start workers (fan-out)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	// gửi jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// đóng results khi worker xong
	go func() {
		wg.Wait()
		close(results)
	}()

	// nhận results (fan-in)
	for res := range results {
		fmt.Println(res)
	}
}
