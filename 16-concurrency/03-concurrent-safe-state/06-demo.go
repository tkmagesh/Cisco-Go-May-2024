/*
To detect data race
go run --race 06-demo.go
*/
package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	wg := &sync.WaitGroup{}
	for range 200 {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	// data race
	count++
}
