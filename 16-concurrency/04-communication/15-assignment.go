package main

import (
	"fmt"
	"sync"
	"time"
)

// Perform the operations (add, subtract...) concurrently
// Printing the results SHOULD be only in the main()

// consumer
func main() {
	start := time.Now()
	addCh := func() chan int {
		ch := make(chan int)
		go func() {
			result := add(100, 200)
			ch <- result
		}()
		return ch
	}()
	subtractCh := func() chan int {
		ch := make(chan int)
		go func() {
			result := subtract(100, 200)
			ch <- result
		}()
		return ch
	}()
	multiplyCh := func() chan int {
		ch := make(chan int)
		go func() {
			result := multiply(100, 200)
			ch <- result
		}()
		return ch
	}()
	divideCh := func() chan int {
		ch := make(chan int)
		go func() {
			result := divide(100, 200)
			ch <- result
		}()
		return ch
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Add Result :", <-addCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Subtract Result :", <-subtractCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Multiply Result :", <-multiplyCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Divide Result :", <-divideCh)
	}()

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Time taken :", elapsed)
}

func add(x, y int) int {
	time.Sleep(time.Duration(5) * time.Second)
	return x + y
}

func subtract(x, y int) int {
	time.Sleep(time.Duration(5) * time.Second)
	return x - y
}

func multiply(x, y int) int {
	time.Sleep(time.Duration(5) * time.Second)
	return x * y
}

func divide(x, y int) int {
	time.Sleep(time.Duration(5) * time.Second)
	return x / y
}
