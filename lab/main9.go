package main

import "fmt"

// stage 1: generate numbers
func generate(n int) <-chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

// stage 2: square numbers
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

// stage 3: filter even numbers
func filterEven(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
		close(out)
	}()

	return out
}

func main() {
	numbers := generate(10)
	squared := square(numbers)
	evens := filterEven(squared)

	for v := range evens {
		fmt.Println(v)
	}
}
