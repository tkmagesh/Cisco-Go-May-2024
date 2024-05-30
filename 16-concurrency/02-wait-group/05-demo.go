package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "# of goroutines to start")
	flag.Parse()
	wg := &sync.WaitGroup{}

	fmt.Printf("Starting %d goroutines.. Hit ENTER to start!\n", count)
	fmt.Scanln()
	for id := range count {
		wg.Add(1) // increment the counter by 1
		go fn(id, wg)
	}
	wg.Wait() // block the execution of this function until the counter becomes 0 (default is 0)
	fmt.Println("Done!!")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
