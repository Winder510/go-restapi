package main

import "fmt"

// func main() {
// 	resultChan := make(chan int)
// 	numbers := []int{1, 2, 3, 4, 5}

// 	var wg sync.WaitGroup

// 	for _, num := range numbers {
// 		wg.Add(1)

// 		go func(n int) {
// 			defer wg.Done()
// 			resultChan <- n * n
// 		}(num)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(resultChan)
// 	}()

// 	for res := range resultChan {
// 		fmt.Println(res)
// 	}
// }

func main() {
	resultChan := make(chan int)
	numbers := []int{1, 2, 3, 4, 5}
	go func() {
		for _, num := range numbers {
			resultChan <- num * num
		}
		close(resultChan)
	}()

	for res := range resultChan {
		fmt.Println(res)
	}
}
