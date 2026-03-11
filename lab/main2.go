package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	workers := [3]string{"A", "B", "C"}
	for _,worker := range workers {
		wg.Add(1)
		go func() {
			for j := 0; j < 3; j++ {
				println("worker", worker, "is working at step ", j)
				time.Sleep(10 * time.Millisecond)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}