package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genData(ch)
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
}

func genData(ch chan int) {
	for i := range rand.Intn(20) {
		ch <- (i + 1) * 10
	}
	close(ch)
}
