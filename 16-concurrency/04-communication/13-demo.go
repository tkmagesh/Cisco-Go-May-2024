package main

import (
	"fmt"
	"time"
)

// Use anon functions to spin up goroutines on the fly

func main() {
	ch := make(chan int)
	go func() {
		result := add(100, 200)
		ch <- result
	}()
	result := <-ch
	fmt.Println(result)
}

/* 3rd party library */
func add(x, y int) int {
	time.Sleep(3 * time.Second)
	result := x + y
	return result
}
