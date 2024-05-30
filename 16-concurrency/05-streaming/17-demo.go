package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genData(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		break
	}

}

func genData(ch chan int) {
	for i := range rand.Intn(20) {
		ch <- (i + 1) * 10
	}
	close(ch)
}
