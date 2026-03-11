// Lab 1 — Goroutine cơ bản
// Bài toán

// In số từ 1 → 5 bằng goroutine.

// Yêu cầu

// tạo goroutine

// main phải đợi goroutine chạy xong

package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(){
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond) 
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printNumbers()
	}()
	wg.Wait()
}