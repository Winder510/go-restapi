package main

import (
	"fmt"
	"time"
)

func main() {
	// 5 request / second -> 1 request mỗi 200ms
	//Ticker sẽ phát tín hiệu mỗi 200ms
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for i := 1; i <= 10; i++ {
		<-ticker.C // chờ tick

		fmt.Printf("Processing request %d at %v\n", i, time.Now().Format("15:04:05.000"))
	}
}
