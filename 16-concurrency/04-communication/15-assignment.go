package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Perform the operations (add, subtract...) concurrently
// Printing the results SHOULD be only in the main()

// consumer
func main() {
	start := time.Now()
	fmt.Println(add(100, 200))
	fmt.Println(subtract(100, 200))
	fmt.Println(multiply(100, 200))
	fmt.Println(divide(100, 200))
	elapsed := time.Since(start)
	fmt.Println("Time taken :", elapsed)
}

func add(x, y int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return x + y
}

func subtract(x, y int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return x - y
}

func multiply(x, y int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return x * y
}

func divide(x, y int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return x / y
}
