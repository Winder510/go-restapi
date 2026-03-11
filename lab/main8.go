package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range jobs {
		fmt.Println("Downloading", file)

		time.Sleep(1 * time.Second)

		fmt.Println("Done", file)
	}
}

func main() {
	files := []string{"file1", "file2", "file3", "file4"}

	jobs := make(chan string)

	var wg sync.WaitGroup

	// start 3 workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// gửi jobs
	for _, file := range files {
		jobs <- file
	}

	close(jobs)

	wg.Wait()
}
